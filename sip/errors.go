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

package sip

import (
	"fmt"
)

// ResponseError encapsulates an unhandled >=400 SIP error response.
type ResponseError struct {
	Msg *Msg
}

func (err *ResponseError) Error() string {
	return fmt.Sprintf("%d: %s", err.Msg.Status, err.Msg.Phrase)
}
