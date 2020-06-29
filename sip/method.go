// Copyright 2020 Justine Alexandra Roberts Tunney
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// SIP Protocol Method Definitions

package sip

const (
	MethodInvite    = "INVITE"    // Indicates a client is being invited to participate in a call session.
	MethodAck       = "ACK"       // Confirms that the client has received a final response to an INVITE request.
	MethodBye       = "BYE"       // Terminates a call and can be sent by either the caller or the callee.
	MethodCancel    = "CANCEL"    // Cancels any pending request.
	MethodOptions   = "OPTIONS"   // Queries the capabilities of servers.
	MethodRegister  = "REGISTER"  // Registers the address listed in the To header field with a SIP server.
	MethodPrack     = "PRACK"     // Provisional acknowledgement.
	MethodSubscribe = "SUBSCRIBE" // Subscribes for an Event of Notification from the Notifier.
	MethodNotify    = "NOTIFY"    // Notify the subscriber of a new Event.
	MethodPublish   = "PUBLISH"   // Publishes an event to the Server.
	MethodInfo      = "INFO"      // Sends mid-session information that does not modify the session state.
	MethodRefer     = "REFER"     // Asks recipient to issue SIP request (call transfer.)
	MethodMessage   = "MESSAGE"   // Transports instant messages using SIP.
	MethodUpdate    = "UPDATE"    // Modifies the state of a session without changing the state of the dialog.
)
