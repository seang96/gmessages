package libgm

import (
	"go.mau.fi/mautrix-gmessages/libgm/events"
	"go.mau.fi/mautrix-gmessages/libgm/pblite"

	"go.mau.fi/mautrix-gmessages/libgm/binary"
)

func (c *Client) handleUpdatesEvent(res *pblite.Response) {
	switch res.Data.Action {
	case binary.ActionType_GET_UPDATES:
		data, ok := res.Data.Decrypted.(*binary.UpdateEvents)
		if !ok {
			c.Logger.Error().Type("data_type", res.Data.Decrypted).Msg("Unexpected data type in GET_UPDATES event")
			return
		}

		switch evt := data.Event.(type) {
		case *binary.UpdateEvents_UserAlertEvent:
			c.rpc.logContent(res)
			c.handleUserAlertEvent(res, evt.UserAlertEvent)

		case *binary.UpdateEvents_SettingsEvent:
			c.rpc.logContent(res)
			c.triggerEvent(evt.SettingsEvent)

		case *binary.UpdateEvents_ConversationEvent:
			if c.rpc.deduplicateUpdate(res) {
				return
			}
			c.triggerEvent(evt.ConversationEvent.GetData())

		case *binary.UpdateEvents_MessageEvent:
			if c.rpc.deduplicateUpdate(res) {
				return
			}
			c.triggerEvent(evt.MessageEvent.GetData())

		case *binary.UpdateEvents_TypingEvent:
			c.rpc.logContent(res)
			c.triggerEvent(evt.TypingEvent.GetData())

		default:
			c.Logger.Trace().Any("evt", evt).Msg("Got unknown event type")
		}
	default:
		c.Logger.Trace().Any("response", res).Msg("Got unexpected response")
	}
}

func (c *Client) handleClientReady(newSessionId string) {
	conversations, convErr := c.ListConversations(25, binary.ListConversationsPayload_INBOX)
	if convErr != nil {
		panic(convErr)
	}
	notifyErr := c.NotifyDittoActivity()
	if notifyErr != nil {
		panic(notifyErr)
	}
	readyEvt := events.NewClientReady(newSessionId, conversations)
	c.triggerEvent(readyEvt)
}

func (c *Client) handleUserAlertEvent(res *pblite.Response, data *binary.UserAlertEvent) {
	alertType := data.AlertType
	switch alertType {
	case binary.AlertType_BROWSER_ACTIVE:
		newSessionID := res.Data.RequestID
		c.Logger.Debug().Any("session_id", newSessionID).Msg("Got browser active notification")
		if newSessionID != c.sessionHandler.sessionID {
			evt := events.NewBrowserActive(newSessionID)
			c.triggerEvent(evt)
		} else {
			go c.handleClientReady(newSessionID)
		}
	default:
		c.triggerEvent(data)
	}
}
