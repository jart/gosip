
//line addr_parse.rl:1
// -*-go-*-

package sip

import (
	"errors"
	"fmt"
	"strings"
)


//line addr_parse.rl:12

//line addr_parse.go:17
const addr_start int = 1
const addr_first_final int = 68
const addr_error int = 0

const addr_en_params int = 69
const addr_en_name_addr int = 35
const addr_en_bare_addr int = 65
const addr_en_addr int = 1


//line addr_parse.rl:13

// ParseAddr turns a SIP address into a data structure.
func ParseAddr(s string) (addr *Addr, err error) {
	if s == "" {
		return nil, errors.New("Empty SIP message")
	}
	return ParseAddrBytes([]byte(s), nil)
}

// ParseAddr turns a SIP address slice into a data structure.
func ParseAddrBytes(data []byte, next *Addr) (addr *Addr, err error) {
	if data == nil {
		return nil, nil
	}
	addr = new(Addr)
	result := addr
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	mark := 0
	var name string

	
//line addr_parse.go:55
	{
	cs = addr_start
	}

//line addr_parse.go:60
	{
	var _widec int16
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 68:
		goto st_case_68
	case 35:
		goto st_case_35
	case 0:
		goto st_case_0
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 74:
		goto st_case_74
	case 75:
		goto st_case_75
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 76:
		goto st_case_76
	case 44:
		goto st_case_44
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 57:
		goto st_case_57
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 61:
		goto st_case_61
	case 62:
		goto st_case_62
	case 63:
		goto st_case_63
	case 64:
		goto st_case_64
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 67:
		goto st_case_67
	case 77:
		goto st_case_77
	case 78:
		goto st_case_78
	case 69:
		goto st_case_69
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 70:
		goto st_case_70
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 71:
		goto st_case_71
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 72:
		goto st_case_72
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 12:
		goto st_case_12
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 73:
		goto st_case_73
	case 15:
		goto st_case_15
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		switch data[p] {
		case 59:
			goto tr1
		case 60:
			goto tr2
		}
		goto tr0
tr109:
//line addr_parse.rl:78

			p = mark
			p--

			{goto st35 }
		
	goto st68
tr0:
//line addr_parse.rl:39

			mark = p
		
	goto st68
tr1:
//line addr_parse.rl:39

			mark = p
		
//line addr_parse.rl:84

			p = mark
			p--

			{goto st65 }
		
	goto st68
tr2:
//line addr_parse.rl:39

			mark = p
		
//line addr_parse.rl:78

			p = mark
			p--

			{goto st35 }
		
	goto st68
tr108:
//line addr_parse.rl:84

			p = mark
			p--

			{goto st65 }
		
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line addr_parse.go:294
		switch data[p] {
		case 59:
			goto tr108
		case 60:
			goto tr109
		}
		goto st68
tr50:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line addr_parse.go:317
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr50
		case 32:
			goto tr50
		case 33:
			goto tr51
		case 34:
			goto st50
		case 37:
			goto tr51
		case 39:
			goto tr51
		case 60:
			goto tr53
		case 126:
			goto tr51
		case 525:
			goto tr54
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr51
				}
			case _widec >= 42:
				goto tr51
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr51
				}
			case _widec >= 65:
				goto tr51
			}
		default:
			goto tr51
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr51:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st36
tr56:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line addr_parse.go:395
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr55
		case 32:
			goto tr55
		case 33:
			goto tr56
		case 37:
			goto tr56
		case 39:
			goto tr56
		case 126:
			goto tr56
		case 525:
			goto tr57
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr56
				}
			case _widec >= 42:
				goto tr56
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr56
				}
			case _widec >= 65:
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
tr55:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st37
tr58:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line addr_parse.go:465
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr58
		case 32:
			goto tr58
		case 33:
			goto tr56
		case 37:
			goto tr56
		case 39:
			goto tr56
		case 60:
			goto tr59
		case 126:
			goto tr56
		case 525:
			goto tr60
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr56
				}
			case _widec >= 42:
				goto tr56
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr56
				}
			case _widec >= 65:
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
tr53:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st38
tr59:
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st38
tr94:
//line addr_parse.rl:60

			addr.Display = string(buf[0:amt])
		
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line addr_parse.go:541
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr61
			}
		case data[p] >= 65:
			goto tr61
		}
		goto st0
tr61:
//line addr_parse.rl:39

			mark = p
		
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line addr_parse.go:562
		switch data[p] {
		case 43:
			goto st39
		case 58:
			goto st40
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st39
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st39
				}
			case data[p] >= 65:
				goto st39
			}
		default:
			goto st39
		}
		goto st0
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
		switch data[p] {
		case 33:
			goto st41
		case 61:
			goto st41
		case 93:
			goto st41
		case 95:
			goto st41
		case 126:
			goto st41
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto st41
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto st0
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
		switch data[p] {
		case 33:
			goto st41
		case 62:
			goto tr65
		case 93:
			goto st41
		case 95:
			goto st41
		case 126:
			goto st41
		}
		switch {
		case data[p] < 61:
			if 36 <= data[p] && data[p] <= 59 {
				goto st41
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st41
			}
		default:
			goto st41
		}
		goto st0
tr65:
//line addr_parse.rl:64

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line addr_parse.go:659
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st74
		case 32:
			goto st74
		case 269:
			goto tr121
		case 525:
			goto st42
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr121
			}
		default:
			goto tr121
		}
		goto st0
tr121:
//line addr_parse.rl:90

			p--

			{goto st69 }
		
	goto st75
	st75:
		if p++; p == pe {
			goto _test_eof75
		}
	st_case_75:
//line addr_parse.go:699
		goto st0
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
		if data[p] == 10 {
			goto st43
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		switch data[p] {
		case 9:
			goto st76
		case 32:
			goto st76
		}
		goto st0
	st76:
		if p++; p == pe {
			goto _test_eof76
		}
	st_case_76:
		switch data[p] {
		case 9:
			goto st76
		case 32:
			goto st76
		}
		goto tr121
tr57:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st44
tr60:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line addr_parse.go:757
		if data[p] == 10 {
			goto tr68
		}
		goto st0
tr68:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line addr_parse.go:774
		switch data[p] {
		case 9:
			goto tr69
		case 32:
			goto tr69
		}
		goto st0
tr69:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st46
tr70:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line addr_parse.go:805
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr70
		case 32:
			goto tr70
		case 33:
			goto tr56
		case 37:
			goto tr56
		case 39:
			goto tr56
		case 60:
			goto tr59
		case 126:
			goto tr56
		case 525:
			goto tr71
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr56
				}
			case _widec >= 42:
				goto tr56
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr56
				}
			case _widec >= 65:
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
tr102:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st47
tr71:
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st47
tr95:
//line addr_parse.rl:60

			addr.Display = string(buf[0:amt])
		
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line addr_parse.go:881
		if data[p] == 10 {
			goto st48
		}
		goto st0
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
		switch data[p] {
		case 9:
			goto st49
		case 32:
			goto st49
		}
		goto st0
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
		switch data[p] {
		case 9:
			goto st49
		case 32:
			goto st49
		case 60:
			goto st38
		}
		goto st0
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr75
		case 34:
			goto tr76
		case 92:
			goto tr77
		case 525:
			goto tr83
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr78
				}
			case _widec >= 32:
				goto tr75
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr80
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr82
				}
			default:
				goto tr81
			}
		default:
			goto tr79
		}
		goto st0
tr75:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st51
tr84:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line addr_parse.go:984
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr84
		case 34:
			goto st52
		case 92:
			goto st54
		case 525:
			goto tr92
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr87
				}
			case _widec >= 32:
				goto tr84
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr89
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr91
				}
			default:
				goto tr90
			}
		default:
			goto tr88
		}
		goto st0
tr76:
//line addr_parse.rl:43

			amt = 0
		
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line addr_parse.go:1040
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr93
		case 32:
			goto tr93
		case 60:
			goto tr94
		case 525:
			goto tr95
		}
		goto st0
tr93:
//line addr_parse.rl:60

			addr.Display = string(buf[0:amt])
		
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line addr_parse.go:1070
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st53
		case 32:
			goto st53
		case 60:
			goto st38
		case 525:
			goto st47
		}
		goto st0
tr77:
//line addr_parse.rl:43

			amt = 0
		
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line addr_parse.go:1100
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr84
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr84
			}
		default:
			goto tr84
		}
		goto st0
tr78:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st55
tr87:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line addr_parse.go:1137
		if 128 <= data[p] && data[p] <= 191 {
			goto tr84
		}
		goto st0
tr79:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st56
tr88:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line addr_parse.go:1165
		if 128 <= data[p] && data[p] <= 191 {
			goto tr87
		}
		goto st0
tr80:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st57
tr89:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line addr_parse.go:1193
		if 128 <= data[p] && data[p] <= 191 {
			goto tr88
		}
		goto st0
tr81:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st58
tr90:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line addr_parse.go:1221
		if 128 <= data[p] && data[p] <= 191 {
			goto tr89
		}
		goto st0
tr82:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st59
tr91:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line addr_parse.go:1249
		if 128 <= data[p] && data[p] <= 191 {
			goto tr90
		}
		goto st0
tr83:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st60
tr92:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line addr_parse.go:1277
		if data[p] == 10 {
			goto tr98
		}
		goto st0
tr98:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line addr_parse.go:1294
		switch data[p] {
		case 9:
			goto tr84
		case 32:
			goto tr84
		}
		goto st0
tr54:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line addr_parse.go:1317
		if data[p] == 10 {
			goto st63
		}
		goto st0
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
		switch data[p] {
		case 9:
			goto st64
		case 32:
			goto st64
		}
		goto st0
tr101:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:56

			addr.Display = strings.TrimRight(string(buf[0:amt]), " \t\r\n")
		
	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line addr_parse.go:1349
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr101
		case 32:
			goto tr101
		case 33:
			goto tr51
		case 34:
			goto st50
		case 37:
			goto tr51
		case 39:
			goto tr51
		case 60:
			goto tr53
		case 126:
			goto tr51
		case 525:
			goto tr102
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr51
				}
			case _widec >= 42:
				goto tr51
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr51
				}
			case _widec >= 65:
				goto tr51
			}
		default:
			goto tr51
		}
		goto st0
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
		switch {
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr103
			}
		case data[p] >= 65:
			goto tr103
		}
		goto st0
tr103:
//line addr_parse.rl:39

			mark = p
		
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line addr_parse.go:1425
		switch data[p] {
		case 43:
			goto st66
		case 58:
			goto st67
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto st66
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto st66
				}
			case data[p] >= 65:
				goto st66
			}
		default:
			goto st66
		}
		goto st0
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
		switch data[p] {
		case 33:
			goto st77
		case 61:
			goto st77
		case 93:
			goto st77
		case 95:
			goto st77
		case 126:
			goto st77
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 58 {
				goto st77
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st77
			}
		default:
			goto st77
		}
		goto st0
	st77:
		if p++; p == pe {
			goto _test_eof77
		}
	st_case_77:
		switch data[p] {
		case 33:
			goto st77
		case 61:
			goto st77
		case 93:
			goto st77
		case 95:
			goto st77
		case 126:
			goto st77
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 58 {
				goto st77
			}
		case data[p] > 91:
			if 97 <= data[p] && data[p] <= 122 {
				goto st77
			}
		default:
			goto st77
		}
		goto tr124
tr124:
//line addr_parse.rl:64

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
//line addr_parse.rl:90

			p--

			{goto st69 }
		
	goto st78
	st78:
		if p++; p == pe {
			goto _test_eof78
		}
	st_case_78:
//line addr_parse.go:1528
		goto st0
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st2
		case 32:
			goto st2
		case 44:
			goto st3
		case 59:
			goto st7
		case 525:
			goto st10
		}
		goto st0
tr116:
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line addr_parse.go:1569
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st2
		case 32:
			goto st2
		case 44:
			goto st3
		case 59:
			goto st7
		case 525:
			goto st10
		}
		goto st0
tr112:
//line addr_parse.rl:52

			name = string(data[mark:p])
		
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st3
tr118:
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line addr_parse.go:1617
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st3
		case 32:
			goto st3
		case 269:
			goto tr8
		case 525:
			goto st4
		}
		switch {
		case _widec > 12:
			if 14 <= _widec {
				goto tr8
			}
		default:
			goto tr8
		}
		goto st0
tr8:
//line addr_parse.rl:95

			addr.Next = new(Addr)
			addr = addr.Next
			p--

			{goto st1 }
		
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line addr_parse.go:1659
		goto st0
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
		if data[p] == 10 {
			goto st5
		}
		goto st0
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
		switch data[p] {
		case 9:
			goto st6
		case 32:
			goto st6
		}
		goto st0
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
		switch data[p] {
		case 9:
			goto st6
		case 32:
			goto st6
		}
		goto tr8
tr113:
//line addr_parse.rl:52

			name = string(data[mark:p])
		
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st7
tr119:
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line addr_parse.go:1721
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st7
		case 32:
			goto st7
		case 33:
			goto tr12
		case 37:
			goto tr12
		case 39:
			goto tr12
		case 126:
			goto tr12
		case 525:
			goto st32
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr12
				}
			case _widec >= 42:
				goto tr12
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto tr12
				}
			case _widec >= 65:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
tr12:
//line addr_parse.rl:39

			mark = p
		
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line addr_parse.go:1779
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr110
		case 32:
			goto tr110
		case 33:
			goto st71
		case 37:
			goto st71
		case 39:
			goto st71
		case 44:
			goto tr112
		case 59:
			goto tr113
		case 61:
			goto tr114
		case 126:
			goto st71
		case 525:
			goto tr115
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto st71
			}
		case _widec > 57:
			switch {
			case _widec > 90:
				if 95 <= _widec && _widec <= 122 {
					goto st71
				}
			case _widec >= 65:
				goto st71
			}
		default:
			goto st71
		}
		goto st0
tr110:
//line addr_parse.rl:52

			name = string(data[mark:p])
		
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line addr_parse.go:1845
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st8
		case 32:
			goto st8
		case 44:
			goto st3
		case 59:
			goto st7
		case 61:
			goto st9
		case 525:
			goto st29
		}
		goto st0
tr114:
//line addr_parse.rl:52

			name = string(data[mark:p])
		
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line addr_parse.go:1879
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st9
		case 32:
			goto st9
		case 33:
			goto tr17
		case 34:
			goto st13
		case 37:
			goto tr17
		case 39:
			goto tr17
		case 93:
			goto tr17
		case 126:
			goto tr17
		case 525:
			goto st23
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr17
				}
			case _widec >= 42:
				goto tr17
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr17
				}
			case _widec >= 65:
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
tr17:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st72
tr117:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line addr_parse.go:1953
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr116
		case 32:
			goto tr116
		case 33:
			goto tr117
		case 37:
			goto tr117
		case 39:
			goto tr117
		case 44:
			goto tr118
		case 59:
			goto tr119
		case 93:
			goto tr117
		case 126:
			goto tr117
		case 525:
			goto tr120
		}
		switch {
		case _widec < 48:
			if 42 <= _widec && _widec <= 46 {
				goto tr117
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr117
				}
			case _widec >= 65:
				goto tr117
			}
		default:
			goto tr117
		}
		goto st0
tr120:
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line addr_parse.go:2015
		if data[p] == 10 {
			goto st11
		}
		goto st0
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
		switch data[p] {
		case 9:
			goto st12
		case 32:
			goto st12
		}
		goto st0
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
		switch data[p] {
		case 9:
			goto st12
		case 32:
			goto st12
		case 44:
			goto st3
		case 59:
			goto st7
		}
		goto st0
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr22
		case 34:
			goto tr23
		case 92:
			goto tr24
		case 525:
			goto tr30
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr25
				}
			case _widec >= 32:
				goto tr22
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr27
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr29
				}
			default:
				goto tr28
			}
		default:
			goto tr26
		}
		goto st0
tr22:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st14
tr31:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line addr_parse.go:2120
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr31
		case 34:
			goto st73
		case 92:
			goto st15
		case 525:
			goto tr39
		}
		switch {
		case _widec < 224:
			switch {
			case _widec > 126:
				if 192 <= _widec && _widec <= 223 {
					goto tr34
				}
			case _widec >= 32:
				goto tr31
			}
		case _widec > 239:
			switch {
			case _widec < 248:
				if 240 <= _widec && _widec <= 247 {
					goto tr36
				}
			case _widec > 251:
				if 252 <= _widec && _widec <= 253 {
					goto tr38
				}
			default:
				goto tr37
			}
		default:
			goto tr35
		}
		goto st0
tr23:
//line addr_parse.rl:43

			amt = 0
		
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line addr_parse.go:2176
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto tr116
		case 32:
			goto tr116
		case 44:
			goto tr118
		case 59:
			goto tr119
		case 525:
			goto tr120
		}
		goto st0
tr24:
//line addr_parse.rl:43

			amt = 0
		
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line addr_parse.go:2208
		switch {
		case data[p] < 11:
			if data[p] <= 9 {
				goto tr31
			}
		case data[p] > 12:
			if 14 <= data[p] && data[p] <= 127 {
				goto tr31
			}
		default:
			goto tr31
		}
		goto st0
tr25:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st16
tr34:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line addr_parse.go:2245
		if 128 <= data[p] && data[p] <= 191 {
			goto tr31
		}
		goto st0
tr26:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st17
tr35:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line addr_parse.go:2273
		if 128 <= data[p] && data[p] <= 191 {
			goto tr34
		}
		goto st0
tr27:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st18
tr36:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line addr_parse.go:2301
		if 128 <= data[p] && data[p] <= 191 {
			goto tr35
		}
		goto st0
tr28:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st19
tr37:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line addr_parse.go:2329
		if 128 <= data[p] && data[p] <= 191 {
			goto tr36
		}
		goto st0
tr29:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st20
tr38:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line addr_parse.go:2357
		if 128 <= data[p] && data[p] <= 191 {
			goto tr37
		}
		goto st0
tr30:
//line addr_parse.rl:43

			amt = 0
		
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st21
tr39:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line addr_parse.go:2385
		if data[p] == 10 {
			goto tr40
		}
		goto st0
tr40:
//line addr_parse.rl:47

			buf[amt] = data[p]
			amt++
		
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line addr_parse.go:2402
		switch data[p] {
		case 9:
			goto tr31
		case 32:
			goto tr31
		}
		goto st0
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
		if data[p] == 10 {
			goto st24
		}
		goto st0
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
		switch data[p] {
		case 9:
			goto st25
		case 32:
			goto st25
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		_widec = int16(data[p])
		if 13 <= data[p] && data[p] <= 13 {
			_widec = 256 + (int16(data[p]) - 0)
			if  p + 2 < pe && (data[p+2] == ' ' || data[p+2] == '\t')  {
				_widec += 256
			}
		}
		switch _widec {
		case 9:
			goto st25
		case 32:
			goto st25
		case 33:
			goto tr17
		case 34:
			goto st13
		case 37:
			goto tr17
		case 39:
			goto tr17
		case 93:
			goto tr17
		case 126:
			goto tr17
		case 525:
			goto st26
		}
		switch {
		case _widec < 48:
			switch {
			case _widec > 43:
				if 45 <= _widec && _widec <= 46 {
					goto tr17
				}
			case _widec >= 42:
				goto tr17
			}
		case _widec > 58:
			switch {
			case _widec > 91:
				if 95 <= _widec && _widec <= 122 {
					goto tr17
				}
			case _widec >= 65:
				goto tr17
			}
		default:
			goto tr17
		}
		goto st0
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
		if data[p] == 10 {
			goto st27
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		switch data[p] {
		case 9:
			goto st28
		case 32:
			goto st28
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 9:
			goto st28
		case 32:
			goto st28
		case 34:
			goto st13
		}
		goto st0
tr115:
//line addr_parse.rl:52

			name = string(data[mark:p])
		
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line addr_parse.go:2539
		if data[p] == 10 {
			goto st30
		}
		goto st0
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
		switch data[p] {
		case 9:
			goto st31
		case 32:
			goto st31
		}
		goto st0
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
		switch data[p] {
		case 9:
			goto st31
		case 32:
			goto st31
		case 44:
			goto st3
		case 59:
			goto st7
		case 61:
			goto st9
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 10 {
			goto st33
		}
		goto st0
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
		switch data[p] {
		case 9:
			goto st34
		case 32:
			goto st34
		}
		goto st0
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
		switch data[p] {
		case 9:
			goto st34
		case 32:
			goto st34
		case 33:
			goto tr12
		case 37:
			goto tr12
		case 39:
			goto tr12
		case 126:
			goto tr12
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 43:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr12
				}
			case data[p] >= 42:
				goto tr12
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 95 <= data[p] && data[p] <= 122 {
					goto tr12
				}
			case data[p] >= 65:
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof
	_test_eof75: cs = 75; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof76: cs = 76; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof77: cs = 77; goto _test_eof
	_test_eof78: cs = 78; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 77:
//line addr_parse.rl:64

			addr.Uri, err = ParseURIBytes(data[mark:p])
			if err != nil { return nil, err }
		
		case 72, 73:
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
		case 1, 68:
//line addr_parse.rl:78

			p = mark
			p--

			{goto st35 }
		
		case 71:
//line addr_parse.rl:52

			name = string(data[mark:p])
		
//line addr_parse.rl:69

			if addr.Params == nil {
				addr.Params = Params{}
			}
			addr.Params[name] = string(buf[0:amt])
		
//line addr_parse.go:2754
		}
	}

	_out: {}
	}

//line addr_parse.rl:198


	if cs < addr_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete SIP address: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in SIP address at offset %d: %s", p, string(data)))
		}
	}

	if next != nil {
		next.Last().Next = result
		return next, nil
	}
	return result, nil
}
