
//line uri_parse.rl:1
// -*-go-*-

package sip

import (
	"bytes"
	"errors"
	"fmt"
)


//line uri_parse.rl:12

//line uri_parse.go:17
const uri_start int = 1
const uri_first_final int = 45
const uri_error int = 0

const uri_en_uriSansUser int = 27
const uri_en_uriWithUser int = 1


//line uri_parse.rl:13

// ParseURI turns a a SIP URI into a data structure.
func ParseURI(s string) (uri *URI, err error) {
	if s == "" {
		return nil, errors.New("Empty URI")
	}
	return ParseURIBytes([]byte(s))
}

// ParseURI turns a a SIP URI byte slice into a data structure.
func ParseURIBytes(data []byte) (uri *URI, err error) {
	if data == nil {
		return nil, nil
	}
	uri = new(URI)
	cs := 0
	p := 0
	pe := len(data)
	eof := len(data)
	buf := make([]byte, len(data))
	amt := 0
	var b1, b2 string
	var hex byte

	
//line uri_parse.rl:149


	
//line uri_parse.go:56
	{
	cs = uri_start
	}

//line uri_parse.rl:152
	if bytes.IndexByte(data, '@') == -1 {
		cs = uri_en_uriSansUser;
	} else {
		cs = uri_en_uriWithUser;
	}
	
//line uri_parse.go:68
	{
	if p == pe {
		goto _test_eof
	}
	switch cs {
	case 1:
		goto st_case_1
	case 0:
		goto st_case_0
	case 2:
		goto st_case_2
	case 3:
		goto st_case_3
	case 4:
		goto st_case_4
	case 5:
		goto st_case_5
	case 6:
		goto st_case_6
	case 7:
		goto st_case_7
	case 8:
		goto st_case_8
	case 9:
		goto st_case_9
	case 10:
		goto st_case_10
	case 11:
		goto st_case_11
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 12:
		goto st_case_12
	case 47:
		goto st_case_47
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 48:
		goto st_case_48
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 51:
		goto st_case_51
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 52:
		goto st_case_52
	case 25:
		goto st_case_25
	case 26:
		goto st_case_26
	case 53:
		goto st_case_53
	case 27:
		goto st_case_27
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 54:
		goto st_case_54
	case 55:
		goto st_case_55
	case 30:
		goto st_case_30
	case 56:
		goto st_case_56
	case 31:
		goto st_case_31
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 57:
		goto st_case_57
	case 34:
		goto st_case_34
	case 35:
		goto st_case_35
	case 36:
		goto st_case_36
	case 58:
		goto st_case_58
	case 59:
		goto st_case_59
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 39:
		goto st_case_39
	case 60:
		goto st_case_60
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 42:
		goto st_case_42
	case 61:
		goto st_case_61
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 62:
		goto st_case_62
	}
	goto st_out
	st1:
		if p++; p == pe {
			goto _test_eof1
		}
	st_case_1:
		if data[p] == 32 {
			goto st1
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st1
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr2
			}
		default:
			goto tr2
		}
		goto st0
st_case_0:
	st0:
		cs = 0
		goto _out
tr2:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st2
tr3:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line uri_parse.go:258
		switch data[p] {
		case 43:
			goto tr3
		case 58:
			goto tr4
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr3
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr3
				}
			case data[p] >= 65:
				goto tr3
			}
		default:
			goto tr3
		}
		goto st0
tr4:
//line uri_parse.rl:57

			uri.Scheme = string(buf[0:amt])
		
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line uri_parse.go:294
		switch data[p] {
		case 33:
			goto tr5
		case 37:
			goto tr6
		case 59:
			goto tr5
		case 61:
			goto tr5
		case 63:
			goto tr5
		case 95:
			goto tr5
		case 126:
			goto tr5
		}
		switch {
		case data[p] < 65:
			if 36 <= data[p] && data[p] <= 57 {
				goto tr5
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr5
			}
		default:
			goto tr5
		}
		goto st0
tr5:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st4
tr7:
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st4
tr12:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line uri_parse.go:355
		switch data[p] {
		case 33:
			goto tr7
		case 37:
			goto st5
		case 58:
			goto tr9
		case 61:
			goto tr7
		case 64:
			goto tr10
		case 95:
			goto tr7
		case 126:
			goto tr7
		}
		switch {
		case data[p] < 63:
			if 36 <= data[p] && data[p] <= 59 {
				goto tr7
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr7
			}
		default:
			goto tr7
		}
		goto st0
tr6:
//line uri_parse.rl:38

			amt = 0
		
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line uri_parse.go:396
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr11
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr11
			}
		default:
			goto tr11
		}
		goto st0
tr11:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st6
	st6:
		if p++; p == pe {
			goto _test_eof6
		}
	st_case_6:
//line uri_parse.go:421
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr12
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr12
			}
		default:
			goto tr12
		}
		goto st0
tr9:
//line uri_parse.rl:61

			uri.User = string(buf[0:amt])
		
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line uri_parse.go:446
		switch data[p] {
		case 33:
			goto tr13
		case 37:
			goto tr14
		case 61:
			goto tr13
		case 95:
			goto tr13
		case 126:
			goto tr13
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto tr13
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr13
				}
			case data[p] >= 65:
				goto tr13
			}
		default:
			goto tr13
		}
		goto st0
tr13:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st8
tr15:
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st8
tr19:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line uri_parse.go:508
		switch data[p] {
		case 33:
			goto tr15
		case 37:
			goto st9
		case 61:
			goto tr15
		case 64:
			goto tr17
		case 95:
			goto tr15
		case 126:
			goto tr15
		}
		switch {
		case data[p] < 48:
			if 36 <= data[p] && data[p] <= 46 {
				goto tr15
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr15
				}
			case data[p] >= 65:
				goto tr15
			}
		default:
			goto tr15
		}
		goto st0
tr14:
//line uri_parse.rl:38

			amt = 0
		
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line uri_parse.go:552
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr18
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr18
			}
		default:
			goto tr18
		}
		goto st0
tr18:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line uri_parse.go:577
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr19
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr19
			}
		default:
			goto tr19
		}
		goto st0
tr10:
//line uri_parse.rl:61

			uri.User = string(buf[0:amt])
		
	goto st11
tr17:
//line uri_parse.rl:65

			uri.Pass = string(buf[0:amt])
		
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line uri_parse.go:608
		switch data[p] {
		case 43:
			goto tr20
		case 91:
			goto st25
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr20
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr20
				}
			case data[p] >= 65:
				goto tr20
			}
		default:
			goto tr20
		}
		goto st0
tr20:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st45
tr69:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line uri_parse.go:664
		switch data[p] {
		case 32:
			goto tr68
		case 43:
			goto tr69
		case 58:
			goto tr70
		case 59:
			goto tr71
		case 63:
			goto tr72
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr69
				}
			case data[p] >= 9:
				goto tr68
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr69
				}
			case data[p] >= 65:
				goto tr69
			}
		default:
			goto tr69
		}
		goto st0
tr68:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st46
tr76:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st46
tr82:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line uri_parse.go:739
		switch data[p] {
		case 32:
			goto st46
		case 59:
			goto st12
		case 63:
			goto st18
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st46
		}
		goto st0
tr71:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st12
tr79:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st12
tr85:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line uri_parse.go:791
		switch data[p] {
		case 33:
			goto tr22
		case 37:
			goto tr23
		case 93:
			goto tr22
		case 95:
			goto tr22
		case 126:
			goto tr22
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto tr22
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr22
				}
			case data[p] >= 65:
				goto tr22
			}
		default:
			goto tr22
		}
		goto st0
tr77:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st47
tr25:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st47
tr22:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line uri_parse.go:866
		switch data[p] {
		case 32:
			goto tr76
		case 33:
			goto tr77
		case 37:
			goto st13
		case 59:
			goto tr79
		case 61:
			goto tr80
		case 63:
			goto tr81
		case 93:
			goto tr77
		case 95:
			goto tr77
		case 126:
			goto tr77
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr77
				}
			case data[p] >= 9:
				goto tr76
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr77
				}
			case data[p] >= 65:
				goto tr77
			}
		default:
			goto tr77
		}
		goto st0
tr23:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line uri_parse.go:926
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr24
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr24
			}
		default:
			goto tr24
		}
		goto st0
tr24:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line uri_parse.go:951
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr25
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr25
			}
		default:
			goto tr25
		}
		goto st0
tr80:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line uri_parse.go:977
		switch data[p] {
		case 33:
			goto tr26
		case 37:
			goto tr27
		case 93:
			goto tr26
		case 95:
			goto tr26
		case 126:
			goto tr26
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto tr26
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr26
				}
			case data[p] >= 65:
				goto tr26
			}
		default:
			goto tr26
		}
		goto st0
tr26:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st48
tr83:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st48
tr29:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line uri_parse.go:1047
		switch data[p] {
		case 32:
			goto tr82
		case 33:
			goto tr83
		case 37:
			goto st16
		case 59:
			goto tr85
		case 63:
			goto tr86
		case 93:
			goto tr83
		case 95:
			goto tr83
		case 126:
			goto tr83
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr83
				}
			case data[p] >= 9:
				goto tr82
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr83
				}
			case data[p] >= 65:
				goto tr83
			}
		default:
			goto tr83
		}
		goto st0
tr27:
//line uri_parse.rl:38

			amt = 0
		
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line uri_parse.go:1100
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr28
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr28
			}
		default:
			goto tr28
		}
		goto st0
tr28:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line uri_parse.go:1125
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr29
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr29
			}
		default:
			goto tr29
		}
		goto st0
tr72:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st18
tr81:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st18
tr86:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st18
tr90:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st18
tr96:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line uri_parse.go:1206
		switch data[p] {
		case 33:
			goto tr30
		case 36:
			goto tr30
		case 37:
			goto tr31
		case 63:
			goto tr30
		case 93:
			goto tr30
		case 95:
			goto tr30
		case 126:
			goto tr30
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto tr30
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr30
				}
			case data[p] >= 65:
				goto tr30
			}
		default:
			goto tr30
		}
		goto st0
tr88:
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st49
tr33:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st49
tr30:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line uri_parse.go:1277
		switch data[p] {
		case 32:
			goto tr87
		case 33:
			goto tr88
		case 37:
			goto st19
		case 38:
			goto tr90
		case 61:
			goto tr91
		case 63:
			goto tr88
		case 93:
			goto tr88
		case 95:
			goto tr88
		case 126:
			goto tr88
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr88
				}
			case data[p] >= 9:
				goto tr87
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr88
				}
			case data[p] >= 65:
				goto tr88
			}
		default:
			goto tr88
		}
		goto st0
tr87:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st50
tr93:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line uri_parse.go:1354
		if data[p] == 32 {
			goto st50
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st50
		}
		goto st0
tr31:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line uri_parse.go:1378
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr32
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr32
			}
		default:
			goto tr32
		}
		goto st0
tr32:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line uri_parse.go:1403
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr33
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr33
			}
		default:
			goto tr33
		}
		goto st0
tr91:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line uri_parse.go:1429
		switch data[p] {
		case 33:
			goto tr34
		case 36:
			goto tr34
		case 37:
			goto tr35
		case 63:
			goto tr34
		case 93:
			goto tr34
		case 95:
			goto tr34
		case 126:
			goto tr34
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto tr34
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr34
				}
			case data[p] >= 65:
				goto tr34
			}
		default:
			goto tr34
		}
		goto st0
tr34:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st51
tr94:
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st51
tr37:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st51
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
//line uri_parse.go:1495
		switch data[p] {
		case 32:
			goto tr93
		case 33:
			goto tr94
		case 37:
			goto st22
		case 38:
			goto tr96
		case 63:
			goto tr94
		case 93:
			goto tr94
		case 95:
			goto tr94
		case 126:
			goto tr94
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr94
				}
			case data[p] >= 9:
				goto tr93
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr94
				}
			case data[p] >= 65:
				goto tr94
			}
		default:
			goto tr94
		}
		goto st0
tr35:
//line uri_parse.rl:38

			amt = 0
		
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line uri_parse.go:1548
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr36
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr36
			}
		default:
			goto tr36
		}
		goto st0
tr36:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line uri_parse.go:1573
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr37
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr37
			}
		default:
			goto tr37
		}
		goto st0
tr70:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line uri_parse.go:1598
		if 48 <= data[p] && data[p] <= 57 {
			goto tr38
		}
		goto st0
tr38:
//line uri_parse.rl:73

			uri.Port = uri.Port * 10 + uint16(data[p] - 0x30)
		
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line uri_parse.go:1614
		switch data[p] {
		case 32:
			goto st46
		case 59:
			goto st12
		case 63:
			goto st18
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr38
			}
		case data[p] >= 9:
			goto st46
		}
		goto st0
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
		if data[p] == 46 {
			goto tr39
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto tr39
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr39
			}
		default:
			goto tr39
		}
		goto st0
tr39:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st26
tr40:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line uri_parse.go:1684
		switch data[p] {
		case 46:
			goto tr40
		case 93:
			goto tr41
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto tr40
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr40
			}
		default:
			goto tr40
		}
		goto st0
tr41:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line uri_parse.go:1715
		switch data[p] {
		case 32:
			goto st46
		case 58:
			goto st24
		case 59:
			goto st12
		case 63:
			goto st18
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st46
		}
		goto st0
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
		if data[p] == 32 {
			goto st27
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st27
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr43
			}
		default:
			goto tr43
		}
		goto st0
tr43:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st28
tr44:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st28
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
//line uri_parse.go:1782
		switch data[p] {
		case 43:
			goto tr44
		case 58:
			goto tr45
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr44
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr44
				}
			case data[p] >= 65:
				goto tr44
			}
		default:
			goto tr44
		}
		goto st0
tr45:
//line uri_parse.rl:57

			uri.Scheme = string(buf[0:amt])
		
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line uri_parse.go:1818
		switch data[p] {
		case 43:
			goto tr46
		case 91:
			goto st43
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr46
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr46
				}
			case data[p] >= 65:
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr46:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st54
tr99:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line uri_parse.go:1874
		switch data[p] {
		case 32:
			goto tr98
		case 43:
			goto tr99
		case 58:
			goto tr100
		case 59:
			goto tr101
		case 63:
			goto tr102
		}
		switch {
		case data[p] < 48:
			switch {
			case data[p] > 13:
				if 45 <= data[p] && data[p] <= 46 {
					goto tr99
				}
			case data[p] >= 9:
				goto tr98
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr99
				}
			case data[p] >= 65:
				goto tr99
			}
		default:
			goto tr99
		}
		goto st0
tr98:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st55
tr106:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st55
tr112:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line uri_parse.go:1949
		switch data[p] {
		case 32:
			goto st55
		case 59:
			goto st30
		case 63:
			goto st36
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st55
		}
		goto st0
tr101:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st30
tr109:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st30
tr115:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line uri_parse.go:2001
		switch data[p] {
		case 33:
			goto tr48
		case 37:
			goto tr49
		case 93:
			goto tr48
		case 95:
			goto tr48
		case 126:
			goto tr48
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto tr48
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr48
				}
			case data[p] >= 65:
				goto tr48
			}
		default:
			goto tr48
		}
		goto st0
tr107:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st56
tr51:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st56
tr48:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line uri_parse.go:2076
		switch data[p] {
		case 32:
			goto tr106
		case 33:
			goto tr107
		case 37:
			goto st31
		case 59:
			goto tr109
		case 61:
			goto tr110
		case 63:
			goto tr111
		case 93:
			goto tr107
		case 95:
			goto tr107
		case 126:
			goto tr107
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr107
				}
			case data[p] >= 9:
				goto tr106
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr107
				}
			case data[p] >= 65:
				goto tr107
			}
		default:
			goto tr107
		}
		goto st0
tr49:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line uri_parse.go:2136
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr50
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr50
			}
		default:
			goto tr50
		}
		goto st0
tr50:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st32
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
//line uri_parse.go:2161
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr51
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr51
			}
		default:
			goto tr51
		}
		goto st0
tr110:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line uri_parse.go:2187
		switch data[p] {
		case 33:
			goto tr52
		case 37:
			goto tr53
		case 93:
			goto tr52
		case 95:
			goto tr52
		case 126:
			goto tr52
		}
		switch {
		case data[p] < 45:
			if 36 <= data[p] && data[p] <= 43 {
				goto tr52
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr52
				}
			case data[p] >= 65:
				goto tr52
			}
		default:
			goto tr52
		}
		goto st0
tr52:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st57
tr113:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st57
tr55:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line uri_parse.go:2257
		switch data[p] {
		case 32:
			goto tr112
		case 33:
			goto tr113
		case 37:
			goto st34
		case 59:
			goto tr115
		case 63:
			goto tr116
		case 93:
			goto tr113
		case 95:
			goto tr113
		case 126:
			goto tr113
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr113
				}
			case data[p] >= 9:
				goto tr112
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr113
				}
			case data[p] >= 65:
				goto tr113
			}
		default:
			goto tr113
		}
		goto st0
tr53:
//line uri_parse.rl:38

			amt = 0
		
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line uri_parse.go:2310
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr54
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr54
			}
		default:
			goto tr54
		}
		goto st0
tr54:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line uri_parse.go:2335
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr55
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr55
			}
		default:
			goto tr55
		}
		goto st0
tr102:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st36
tr111:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st36
tr116:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
	goto st36
tr120:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st36
tr126:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line uri_parse.go:2416
		switch data[p] {
		case 33:
			goto tr56
		case 36:
			goto tr56
		case 37:
			goto tr57
		case 63:
			goto tr56
		case 93:
			goto tr56
		case 95:
			goto tr56
		case 126:
			goto tr56
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto tr56
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr56
				}
			case data[p] >= 65:
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
tr118:
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st58
tr59:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st58
tr56:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line uri_parse.go:2487
		switch data[p] {
		case 32:
			goto tr117
		case 33:
			goto tr118
		case 37:
			goto st37
		case 38:
			goto tr120
		case 61:
			goto tr121
		case 63:
			goto tr118
		case 93:
			goto tr118
		case 95:
			goto tr118
		case 126:
			goto tr118
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr118
				}
			case data[p] >= 9:
				goto tr117
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr118
				}
			case data[p] >= 65:
				goto tr118
			}
		default:
			goto tr118
		}
		goto st0
tr117:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st59
tr123:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line uri_parse.go:2564
		if data[p] == 32 {
			goto st59
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st59
		}
		goto st0
tr57:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line uri_parse.go:2588
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr58
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr58
			}
		default:
			goto tr58
		}
		goto st0
tr58:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line uri_parse.go:2613
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr59
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr59
			}
		default:
			goto tr59
		}
		goto st0
tr121:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line uri_parse.go:2639
		switch data[p] {
		case 33:
			goto tr60
		case 36:
			goto tr60
		case 37:
			goto tr61
		case 63:
			goto tr60
		case 93:
			goto tr60
		case 95:
			goto tr60
		case 126:
			goto tr60
		}
		switch {
		case data[p] < 45:
			if 39 <= data[p] && data[p] <= 43 {
				goto tr60
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr60
				}
			case data[p] >= 65:
				goto tr60
			}
		default:
			goto tr60
		}
		goto st0
tr60:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st60
tr124:
//line uri_parse.rl:42

			buf[amt] = data[p]
			amt++
		
	goto st60
tr63:
//line uri_parse.rl:51

			hex += unhex(data[p])
			buf[amt] = hex
			amt++
		
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line uri_parse.go:2705
		switch data[p] {
		case 32:
			goto tr123
		case 33:
			goto tr124
		case 37:
			goto st40
		case 38:
			goto tr126
		case 63:
			goto tr124
		case 93:
			goto tr124
		case 95:
			goto tr124
		case 126:
			goto tr124
		}
		switch {
		case data[p] < 45:
			switch {
			case data[p] > 13:
				if 36 <= data[p] && data[p] <= 43 {
					goto tr124
				}
			case data[p] >= 9:
				goto tr123
			}
		case data[p] > 58:
			switch {
			case data[p] > 91:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr124
				}
			case data[p] >= 65:
				goto tr124
			}
		default:
			goto tr124
		}
		goto st0
tr61:
//line uri_parse.rl:38

			amt = 0
		
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line uri_parse.go:2758
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr62
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr62
			}
		default:
			goto tr62
		}
		goto st0
tr62:
//line uri_parse.rl:47

			hex = unhex(data[p]) * 16
		
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line uri_parse.go:2783
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr63
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr63
			}
		default:
			goto tr63
		}
		goto st0
tr100:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line uri_parse.go:2808
		if 48 <= data[p] && data[p] <= 57 {
			goto tr64
		}
		goto st0
tr64:
//line uri_parse.rl:73

			uri.Port = uri.Port * 10 + uint16(data[p] - 0x30)
		
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line uri_parse.go:2824
		switch data[p] {
		case 32:
			goto st55
		case 59:
			goto st30
		case 63:
			goto st36
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr64
			}
		case data[p] >= 9:
			goto st55
		}
		goto st0
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
		if data[p] == 46 {
			goto tr65
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto tr65
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr65
			}
		default:
			goto tr65
		}
		goto st0
tr65:
//line uri_parse.rl:38

			amt = 0
		
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st44
tr66:
//line uri_parse.rl:87

			if 'A' <= data[p] && data[p] <= 'Z' {
				buf[amt] = data[p] + 0x20
			} else {
				buf[amt] = data[p]
			}
			amt++
		
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line uri_parse.go:2894
		switch data[p] {
		case 46:
			goto tr66
		case 93:
			goto tr67
		}
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 58 {
				goto tr66
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr66
			}
		default:
			goto tr66
		}
		goto st0
tr67:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line uri_parse.go:2925
		switch data[p] {
		case 32:
			goto st55
		case 58:
			goto st42
		case 59:
			goto st30
		case 63:
			goto st36
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st55
		}
		goto st0
	st_out:
	_test_eof1: cs = 1; goto _test_eof
	_test_eof2: cs = 2; goto _test_eof
	_test_eof3: cs = 3; goto _test_eof
	_test_eof4: cs = 4; goto _test_eof
	_test_eof5: cs = 5; goto _test_eof
	_test_eof6: cs = 6; goto _test_eof
	_test_eof7: cs = 7; goto _test_eof
	_test_eof8: cs = 8; goto _test_eof
	_test_eof9: cs = 9; goto _test_eof
	_test_eof10: cs = 10; goto _test_eof
	_test_eof11: cs = 11; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof55: cs = 55; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 45, 54:
//line uri_parse.rl:69

			uri.Host = string(buf[0:amt])
		
		case 47, 56:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
		case 49, 58:
//line uri_parse.rl:77

			b1 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
		case 48, 57:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:96

			if uri.Params == nil {
				uri.Params = Params{}
			}
			uri.Params[b1] = b2
		
		case 51, 60:
//line uri_parse.rl:82

			b2 = string(buf[0:amt])
			amt = 0
		
//line uri_parse.rl:103

			if uri.Headers == nil {
				uri.Headers = URIHeaders{}
			}
			uri.Headers[b1] = b2
		
//line uri_parse.go:3064
		}
	}

	_out: {}
	}

//line uri_parse.rl:158

	if cs < uri_first_final {
		if p == pe {
			return nil, errors.New(fmt.Sprintf("Incomplete URI: %s", data))
		} else {
			return nil, errors.New(fmt.Sprintf("Error in URI at pos %d: %s", p, data))
		}
	}
	return uri, nil
}
