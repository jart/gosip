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

// RFC2833 RTP DTMF Telephone Events

package rtp

import (
	"errors"
	"fmt"
)

// Turns telephone event into ASCII character.
func DtmfToChar(event uint8) (byte, error) {
	switch event {
	case 0:
		return '0', nil
	case 1:
		return '1', nil
	case 2:
		return '2', nil
	case 3:
		return '3', nil
	case 4:
		return '4', nil
	case 5:
		return '5', nil
	case 6:
		return '6', nil
	case 7:
		return '7', nil
	case 8:
		return '8', nil
	case 9:
		return '9', nil
	case 10:
		return '*', nil
	case 11:
		return '#', nil
	case 12:
		return 'A', nil
	case 13:
		return 'B', nil
	case 14:
		return 'C', nil
	case 15:
		return 'D', nil
	case 16:
		return '!', nil
	default:
		return '\x00', errors.New(fmt.Sprintf("bad tel event: %v", event))
	}
}

// Turns ascii character into corresponding telephone event.
func CharToDtmf(ch byte) (event uint8, err error) {
	switch ch {
	case '0':
		event = 0
	case '1':
		event = 1
	case '2':
		event = 2
	case '3':
		event = 3
	case '4':
		event = 4
	case '5':
		event = 5
	case '6':
		event = 6
	case '7':
		event = 7
	case '8':
		event = 8
	case '9':
		event = 9
	case '*':
		event = 10
	case '#':
		event = 11
	case 'a', 'A':
		event = 12
	case 'b', 'B':
		event = 13
	case 'c', 'C':
		event = 14
	case 'd', 'D':
		event = 15
	case '!':
		event = 16
	default:
		err = errors.New(fmt.Sprint("bad dtmf char:", string(ch)))
	}
	return
}
