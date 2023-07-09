package routes

import "go.mau.fi/mautrix-gmessages/libgm/binary"

var LIST_CONVERSATIONS_WITH_UPDATES = Route{
	Action:         binary.ActionType_LIST_CONVERSATIONS,
	MessageType:    binary.MessageType_BUGLE_ANNOTATION,
	BugleRoute:     binary.BugleRoute_DataEvent,
	ResponseStruct: &binary.Conversations{},
	UseSessionID:   false,
	UseTTL:         true,
}

var LIST_CONVERSATIONS = Route{
	Action:         binary.ActionType_LIST_CONVERSATIONS,
	MessageType:    binary.MessageType_BUGLE_MESSAGE,
	BugleRoute:     binary.BugleRoute_DataEvent,
	ResponseStruct: &binary.Conversations{},
	UseSessionID:   false,
	UseTTL:         true,
}

var GET_CONVERSATION_TYPE = Route{
	Action:         binary.ActionType_GET_CONVERSATION_TYPE,
	MessageType:    binary.MessageType_BUGLE_MESSAGE,
	BugleRoute:     binary.BugleRoute_DataEvent,
	ResponseStruct: &binary.GetConversationTypeResponse{},
	UseSessionID:   false,
	UseTTL:         true,
}