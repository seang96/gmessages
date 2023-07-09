// mautrix-gmessages - A Matrix-Google Messages puppeting bridge.
// Copyright (C) 2023 Tulir Asokan
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <https://www.gnu.org/licenses/>.

package main

import (
	"maunium.net/go/mautrix/bridge/status"
)

const (
	GMListenError      status.BridgeStateErrorCode = "gm-listen-error"
	GMFatalError       status.BridgeStateErrorCode = "gm-listen-fatal-error"
	GMNotConnected     status.BridgeStateErrorCode = "gm-not-connected"
	GMConnecting       status.BridgeStateErrorCode = "gm-connecting"
	GMConnectionFailed status.BridgeStateErrorCode = "gm-connection-failed"
)

func init() {
	status.BridgeStateHumanErrors.Update(status.BridgeStateErrorMap{})
}

func (user *User) GetRemoteID() string {
	if user == nil {
		return ""
	}
	return user.Phone
}

func (user *User) GetRemoteName() string {
	return ""
}

/*func (prov *ProvisioningAPI) BridgeStatePing(w http.ResponseWriter, r *http.Request) {
	if !prov.bridge.AS.CheckServerToken(w, r) {
		return
	}
	userID := r.URL.Query().Get("user_id")
	user := prov.bridge.GetUserByMXID(id.UserID(userID))
	var global status.BridgeState
	global.StateEvent = status.StateRunning
	var remote status.BridgeState
	if user.IsConnected() {
		if user.Client.IsLoggedIn() {
			remote.StateEvent = status.StateConnected
		} else if user.Session != nil {
			remote.StateEvent = status.StateConnecting
			remote.Error = WAConnecting
		} // else: unconfigured
	} else if user.Session != nil {
		remote.StateEvent = status.StateBadCredentials
		remote.Error = WANotConnected
	} // else: unconfigured
	global = global.Fill(nil)
	resp := status.GlobalBridgeState{
		BridgeState:  global,
		RemoteStates: map[string]status.BridgeState{},
	}
	if len(remote.StateEvent) > 0 {
		remote = remote.Fill(user)
		resp.RemoteStates[remote.RemoteID] = remote
	}
	user.log.Debugfln("Responding bridge state in bridge status endpoint: %+v", resp)
	jsonResponse(w, http.StatusOK, &resp)
	if len(resp.RemoteStates) > 0 {
		user.BridgeState.SetPrev(remote)
	}
}*/