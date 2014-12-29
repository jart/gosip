
//line uri_parse.rl:1
package sip

import (
  "bytes"
  "errors"
  "fmt"
  "strconv"
)


//line uri_parse.rl:11

//line uri_parse.go:16
const uri_start int = 1
const uri_first_final int = 55
const uri_error int = 0

const uri_en_uriSansUser int = 32
const uri_en_uriWithUser int = 1


//line uri_parse.rl:12

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

  
//line uri_parse.rl:119


  
//line uri_parse.go:55
	{
	cs = uri_start
	}

//line uri_parse.rl:122
  if bytes.IndexByte(data, '@') == -1 {
    cs = uri_en_uriSansUser;
  } else {
    cs = uri_en_uriWithUser;
  }
  
//line uri_parse.go:67
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
	case 55:
		goto st_case_55
	case 56:
		goto st_case_56
	case 12:
		goto st_case_12
	case 57:
		goto st_case_57
	case 13:
		goto st_case_13
	case 14:
		goto st_case_14
	case 15:
		goto st_case_15
	case 58:
		goto st_case_58
	case 16:
		goto st_case_16
	case 17:
		goto st_case_17
	case 18:
		goto st_case_18
	case 59:
		goto st_case_59
	case 60:
		goto st_case_60
	case 19:
		goto st_case_19
	case 20:
		goto st_case_20
	case 21:
		goto st_case_21
	case 61:
		goto st_case_61
	case 22:
		goto st_case_22
	case 23:
		goto st_case_23
	case 24:
		goto st_case_24
	case 25:
		goto st_case_25
	case 62:
		goto st_case_62
	case 26:
		goto st_case_26
	case 27:
		goto st_case_27
	case 63:
		goto st_case_63
	case 28:
		goto st_case_28
	case 29:
		goto st_case_29
	case 30:
		goto st_case_30
	case 31:
		goto st_case_31
	case 64:
		goto st_case_64
	case 32:
		goto st_case_32
	case 33:
		goto st_case_33
	case 34:
		goto st_case_34
	case 65:
		goto st_case_65
	case 66:
		goto st_case_66
	case 35:
		goto st_case_35
	case 67:
		goto st_case_67
	case 36:
		goto st_case_36
	case 37:
		goto st_case_37
	case 38:
		goto st_case_38
	case 68:
		goto st_case_68
	case 39:
		goto st_case_39
	case 40:
		goto st_case_40
	case 41:
		goto st_case_41
	case 69:
		goto st_case_69
	case 70:
		goto st_case_70
	case 42:
		goto st_case_42
	case 43:
		goto st_case_43
	case 44:
		goto st_case_44
	case 71:
		goto st_case_71
	case 45:
		goto st_case_45
	case 46:
		goto st_case_46
	case 47:
		goto st_case_47
	case 48:
		goto st_case_48
	case 72:
		goto st_case_72
	case 49:
		goto st_case_49
	case 50:
		goto st_case_50
	case 73:
		goto st_case_73
	case 51:
		goto st_case_51
	case 52:
		goto st_case_52
	case 53:
		goto st_case_53
	case 54:
		goto st_case_54
	case 74:
		goto st_case_74
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
				goto tr3
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
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st2
tr3:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st2
tr4:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st2
tr6:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st2
	st2:
		if p++; p == pe {
			goto _test_eof2
		}
	st_case_2:
//line uri_parse.go:275
		switch data[p] {
		case 43:
			goto tr4
		case 58:
			goto tr5
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr4
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr4
				}
			case data[p] >= 65:
				goto tr6
			}
		default:
			goto tr4
		}
		goto st0
tr5:
//line uri_parse.rl:46
 uri.Scheme = string(buf[0:amt])  
	goto st3
	st3:
		if p++; p == pe {
			goto _test_eof3
		}
	st_case_3:
//line uri_parse.go:309
		switch data[p] {
		case 37:
			goto tr8
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 64 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr7
tr7:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st4
tr9:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st4
tr14:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st4
	st4:
		if p++; p == pe {
			goto _test_eof4
		}
	st_case_4:
//line uri_parse.go:357
		switch data[p] {
		case 37:
			goto st5
		case 38:
			goto st0
		case 47:
			goto st0
		case 58:
			goto tr11
		case 64:
			goto tr12
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 59 <= data[p] && data[p] <= 63 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr9
tr8:
//line uri_parse.rl:37
 amt = 0                          
	goto st5
	st5:
		if p++; p == pe {
			goto _test_eof5
		}
	st_case_5:
//line uri_parse.go:398
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr13
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr13
			}
		default:
			goto tr13
		}
		goto st0
tr13:
//line uri_parse.rl:40
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
				goto tr14
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr14
			}
		default:
			goto tr14
		}
		goto st0
tr11:
//line uri_parse.rl:43
 uri.User = string(buf[0:amt])    
	goto st7
	st7:
		if p++; p == pe {
			goto _test_eof7
		}
	st_case_7:
//line uri_parse.go:444
		switch data[p] {
		case 37:
			goto tr16
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 64 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr15
tr15:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st8
tr17:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st8
tr21:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st8
	st8:
		if p++; p == pe {
			goto _test_eof8
		}
	st_case_8:
//line uri_parse.go:492
		switch data[p] {
		case 37:
			goto st9
		case 38:
			goto st0
		case 47:
			goto st0
		case 64:
			goto tr19
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 63 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr17
tr16:
//line uri_parse.rl:37
 amt = 0                          
	goto st9
	st9:
		if p++; p == pe {
			goto _test_eof9
		}
	st_case_9:
//line uri_parse.go:531
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr20
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr20
			}
		default:
			goto tr20
		}
		goto st0
tr20:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st10
	st10:
		if p++; p == pe {
			goto _test_eof10
		}
	st_case_10:
//line uri_parse.go:554
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr21
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr21
			}
		default:
			goto tr21
		}
		goto st0
tr12:
//line uri_parse.rl:43
 uri.User = string(buf[0:amt])    
	goto st11
tr19:
//line uri_parse.rl:44
 uri.Pass = string(buf[0:amt])    
	goto st11
	st11:
		if p++; p == pe {
			goto _test_eof11
		}
	st_case_11:
//line uri_parse.go:581
		switch data[p] {
		case 37:
			goto tr23
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st28
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr24
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto tr22
tr24:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st55
tr22:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st55
tr99:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st55
tr104:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st55
tr45:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st55
	st55:
		if p++; p == pe {
			goto _test_eof55
		}
	st_case_55:
//line uri_parse.go:644
		switch data[p] {
		case 32:
			goto tr98
		case 37:
			goto st24
		case 38:
			goto st0
		case 47:
			goto st0
		case 58:
			goto tr101
		case 59:
			goto tr102
		case 63:
			goto tr103
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr98
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] < 60:
				if 34 <= data[p] && data[p] <= 35 {
					goto st0
				}
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr104
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr99
tr98:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st56
tr108:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st56
tr115:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st56
tr131:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st56
tr136:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st56
	st56:
		if p++; p == pe {
			goto _test_eof56
		}
	st_case_56:
//line uri_parse.go:747
		switch data[p] {
		case 32:
			goto st56
		case 59:
			goto st12
		case 63:
			goto st18
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st0
tr102:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st12
tr111:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st12
tr118:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st12
tr134:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st12
tr139:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st12
	st12:
		if p++; p == pe {
			goto _test_eof12
		}
	st_case_12:
//line uri_parse.go:813
		switch data[p] {
		case 37:
			goto tr27
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr28
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto tr26
tr109:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st57
tr114:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st57
tr30:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st57
tr26:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st57
tr28:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st57
	st57:
		if p++; p == pe {
			goto _test_eof57
		}
	st_case_57:
//line uri_parse.go:880
		switch data[p] {
		case 32:
			goto tr108
		case 37:
			goto st13
		case 38:
			goto st0
		case 47:
			goto st0
		case 59:
			goto tr111
		case 61:
			goto tr112
		case 63:
			goto tr113
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr108
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] < 58:
				if 34 <= data[p] && data[p] <= 35 {
					goto st0
				}
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr114
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr109
tr27:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
	goto st13
	st13:
		if p++; p == pe {
			goto _test_eof13
		}
	st_case_13:
//line uri_parse.go:941
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
tr29:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st14
	st14:
		if p++; p == pe {
			goto _test_eof14
		}
	st_case_14:
//line uri_parse.go:964
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr30
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr30
			}
		default:
			goto tr30
		}
		goto st0
tr112:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
	goto st15
	st15:
		if p++; p == pe {
			goto _test_eof15
		}
	st_case_15:
//line uri_parse.go:987
		switch data[p] {
		case 37:
			goto tr32
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr33
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto tr31
tr33:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st58
tr31:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st58
tr116:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st58
tr120:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st58
tr35:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st58
	st58:
		if p++; p == pe {
			goto _test_eof58
		}
	st_case_58:
//line uri_parse.go:1050
		switch data[p] {
		case 32:
			goto tr115
		case 37:
			goto st16
		case 38:
			goto st0
		case 47:
			goto st0
		case 59:
			goto tr118
		case 63:
			goto tr119
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr115
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] < 58:
				if 34 <= data[p] && data[p] <= 35 {
					goto st0
				}
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr120
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr116
tr32:
//line uri_parse.rl:37
 amt = 0                          
	goto st16
	st16:
		if p++; p == pe {
			goto _test_eof16
		}
	st_case_16:
//line uri_parse.go:1107
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr34
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr34
			}
		default:
			goto tr34
		}
		goto st0
tr34:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st17
	st17:
		if p++; p == pe {
			goto _test_eof17
		}
	st_case_17:
//line uri_parse.go:1130
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr35
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr35
			}
		default:
			goto tr35
		}
		goto st0
tr103:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st18
tr113:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st18
tr119:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st18
tr124:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st18
tr130:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st18
tr135:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st18
tr140:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st18
	st18:
		if p++; p == pe {
			goto _test_eof18
		}
	st_case_18:
//line uri_parse.go:1219
		switch data[p] {
		case 37:
			goto tr37
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 64 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr36
tr122:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st59
tr39:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st59
tr36:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st59
	st59:
		if p++; p == pe {
			goto _test_eof59
		}
	st_case_59:
//line uri_parse.go:1269
		switch data[p] {
		case 32:
			goto tr121
		case 37:
			goto st19
		case 38:
			goto tr124
		case 47:
			goto st0
		case 61:
			goto tr125
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr121
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] > 35:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] >= 34:
				goto st0
			}
		default:
			goto st0
		}
		goto tr122
tr121:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st60
tr127:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st60
	st60:
		if p++; p == pe {
			goto _test_eof60
		}
	st_case_60:
//line uri_parse.go:1338
		if data[p] == 32 {
			goto st60
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st60
		}
		goto st0
tr37:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
	goto st19
	st19:
		if p++; p == pe {
			goto _test_eof19
		}
	st_case_19:
//line uri_parse.go:1357
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr38
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr38
			}
		default:
			goto tr38
		}
		goto st0
tr38:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st20
	st20:
		if p++; p == pe {
			goto _test_eof20
		}
	st_case_20:
//line uri_parse.go:1380
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
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
tr125:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
	goto st21
	st21:
		if p++; p == pe {
			goto _test_eof21
		}
	st_case_21:
//line uri_parse.go:1403
		switch data[p] {
		case 37:
			goto tr41
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 64 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr40
tr40:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st61
tr128:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st61
tr43:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st61
	st61:
		if p++; p == pe {
			goto _test_eof61
		}
	st_case_61:
//line uri_parse.go:1451
		switch data[p] {
		case 32:
			goto tr127
		case 37:
			goto st22
		case 38:
			goto tr130
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr127
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] > 35:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] >= 34:
				goto st0
			}
		default:
			goto st0
		}
		goto tr128
tr41:
//line uri_parse.rl:37
 amt = 0                          
	goto st22
	st22:
		if p++; p == pe {
			goto _test_eof22
		}
	st_case_22:
//line uri_parse.go:1500
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr42
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr42
			}
		default:
			goto tr42
		}
		goto st0
tr42:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st23
	st23:
		if p++; p == pe {
			goto _test_eof23
		}
	st_case_23:
//line uri_parse.go:1523
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr43
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr43
			}
		default:
			goto tr43
		}
		goto st0
tr23:
//line uri_parse.rl:37
 amt = 0                          
	goto st24
	st24:
		if p++; p == pe {
			goto _test_eof24
		}
	st_case_24:
//line uri_parse.go:1546
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr44
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr44
			}
		default:
			goto tr44
		}
		goto st0
tr44:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st25
	st25:
		if p++; p == pe {
			goto _test_eof25
		}
	st_case_25:
//line uri_parse.go:1569
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr45
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr45
			}
		default:
			goto tr45
		}
		goto st0
tr101:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st62
	st62:
		if p++; p == pe {
			goto _test_eof62
		}
	st_case_62:
//line uri_parse.go:1592
		switch data[p] {
		case 32:
			goto tr131
		case 37:
			goto tr132
		case 59:
			goto tr134
		case 63:
			goto tr135
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr133
			}
		case data[p] >= 9:
			goto tr131
		}
		goto st0
tr132:
//line uri_parse.rl:37
 amt = 0                          
	goto st26
	st26:
		if p++; p == pe {
			goto _test_eof26
		}
	st_case_26:
//line uri_parse.go:1621
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr46
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr46
			}
		default:
			goto tr46
		}
		goto st0
tr46:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st27
	st27:
		if p++; p == pe {
			goto _test_eof27
		}
	st_case_27:
//line uri_parse.go:1644
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr47
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr47
			}
		default:
			goto tr47
		}
		goto st0
tr133:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st63
tr138:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st63
tr47:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st63
	st63:
		if p++; p == pe {
			goto _test_eof63
		}
	st_case_63:
//line uri_parse.go:1678
		switch data[p] {
		case 32:
			goto tr136
		case 37:
			goto st26
		case 59:
			goto tr139
		case 63:
			goto tr140
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr138
			}
		case data[p] >= 9:
			goto tr136
		}
		goto st0
	st28:
		if p++; p == pe {
			goto _test_eof28
		}
	st_case_28:
		switch data[p] {
		case 37:
			goto tr49
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr50
				}
			case data[p] >= 59:
				goto st0
			}
		default:
			goto st0
		}
		goto tr48
tr50:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st29
tr48:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st29
tr51:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st29
tr53:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st29
tr56:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st29
	st29:
		if p++; p == pe {
			goto _test_eof29
		}
	st_case_29:
//line uri_parse.go:1765
		switch data[p] {
		case 37:
			goto st30
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto tr54
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr53
				}
			case data[p] >= 59:
				goto st0
			}
		default:
			goto st0
		}
		goto tr51
tr49:
//line uri_parse.rl:37
 amt = 0                          
	goto st30
	st30:
		if p++; p == pe {
			goto _test_eof30
		}
	st_case_30:
//line uri_parse.go:1807
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
tr55:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st31
	st31:
		if p++; p == pe {
			goto _test_eof31
		}
	st_case_31:
//line uri_parse.go:1830
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr56
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr56
			}
		default:
			goto tr56
		}
		goto st0
tr54:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st64
	st64:
		if p++; p == pe {
			goto _test_eof64
		}
	st_case_64:
//line uri_parse.go:1853
		switch data[p] {
		case 32:
			goto st56
		case 58:
			goto st62
		case 59:
			goto st12
		case 63:
			goto st18
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st56
		}
		goto st0
	st32:
		if p++; p == pe {
			goto _test_eof32
		}
	st_case_32:
		if data[p] == 32 {
			goto st32
		}
		switch {
		case data[p] < 65:
			if 9 <= data[p] && data[p] <= 13 {
				goto st32
			}
		case data[p] > 90:
			if 97 <= data[p] && data[p] <= 122 {
				goto tr59
			}
		default:
			goto tr58
		}
		goto st0
tr58:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st33
tr59:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st33
tr60:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st33
tr62:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st33
	st33:
		if p++; p == pe {
			goto _test_eof33
		}
	st_case_33:
//line uri_parse.go:1914
		switch data[p] {
		case 43:
			goto tr60
		case 58:
			goto tr61
		}
		switch {
		case data[p] < 48:
			if 45 <= data[p] && data[p] <= 46 {
				goto tr60
			}
		case data[p] > 57:
			switch {
			case data[p] > 90:
				if 97 <= data[p] && data[p] <= 122 {
					goto tr60
				}
			case data[p] >= 65:
				goto tr62
			}
		default:
			goto tr60
		}
		goto st0
tr61:
//line uri_parse.rl:46
 uri.Scheme = string(buf[0:amt])  
	goto st34
	st34:
		if p++; p == pe {
			goto _test_eof34
		}
	st_case_34:
//line uri_parse.go:1948
		switch data[p] {
		case 37:
			goto tr64
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st51
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr65
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto tr63
tr65:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st65
tr63:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st65
tr143:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st65
tr148:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st65
tr86:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st65
	st65:
		if p++; p == pe {
			goto _test_eof65
		}
	st_case_65:
//line uri_parse.go:2011
		switch data[p] {
		case 32:
			goto tr142
		case 37:
			goto st47
		case 38:
			goto st0
		case 47:
			goto st0
		case 58:
			goto tr145
		case 59:
			goto tr146
		case 63:
			goto tr147
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr142
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] < 60:
				if 34 <= data[p] && data[p] <= 35 {
					goto st0
				}
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr148
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr143
tr142:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st66
tr152:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st66
tr159:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st66
tr175:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st66
tr180:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st66
	st66:
		if p++; p == pe {
			goto _test_eof66
		}
	st_case_66:
//line uri_parse.go:2114
		switch data[p] {
		case 32:
			goto st66
		case 59:
			goto st35
		case 63:
			goto st41
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st66
		}
		goto st0
tr146:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st35
tr155:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st35
tr162:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st35
tr178:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st35
tr183:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st35
	st35:
		if p++; p == pe {
			goto _test_eof35
		}
	st_case_35:
//line uri_parse.go:2180
		switch data[p] {
		case 37:
			goto tr68
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr69
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto tr67
tr153:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st67
tr158:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st67
tr71:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st67
tr67:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st67
tr69:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st67
	st67:
		if p++; p == pe {
			goto _test_eof67
		}
	st_case_67:
//line uri_parse.go:2247
		switch data[p] {
		case 32:
			goto tr152
		case 37:
			goto st36
		case 38:
			goto st0
		case 47:
			goto st0
		case 59:
			goto tr155
		case 61:
			goto tr156
		case 63:
			goto tr157
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr152
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] < 58:
				if 34 <= data[p] && data[p] <= 35 {
					goto st0
				}
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr158
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr153
tr68:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
	goto st36
	st36:
		if p++; p == pe {
			goto _test_eof36
		}
	st_case_36:
//line uri_parse.go:2308
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr70
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr70
			}
		default:
			goto tr70
		}
		goto st0
tr70:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st37
	st37:
		if p++; p == pe {
			goto _test_eof37
		}
	st_case_37:
//line uri_parse.go:2331
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr71
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr71
			}
		default:
			goto tr71
		}
		goto st0
tr156:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
	goto st38
	st38:
		if p++; p == pe {
			goto _test_eof38
		}
	st_case_38:
//line uri_parse.go:2354
		switch data[p] {
		case 37:
			goto tr73
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr74
				}
			case data[p] >= 58:
				goto st0
			}
		default:
			goto st0
		}
		goto tr72
tr74:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st68
tr72:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st68
tr160:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st68
tr164:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st68
tr76:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st68
	st68:
		if p++; p == pe {
			goto _test_eof68
		}
	st_case_68:
//line uri_parse.go:2417
		switch data[p] {
		case 32:
			goto tr159
		case 37:
			goto st39
		case 38:
			goto st0
		case 47:
			goto st0
		case 59:
			goto tr162
		case 63:
			goto tr163
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr159
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] < 58:
				if 34 <= data[p] && data[p] <= 35 {
					goto st0
				}
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr164
				}
			default:
				goto st0
			}
		default:
			goto st0
		}
		goto tr160
tr73:
//line uri_parse.rl:37
 amt = 0                          
	goto st39
	st39:
		if p++; p == pe {
			goto _test_eof39
		}
	st_case_39:
//line uri_parse.go:2474
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr75
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr75
			}
		default:
			goto tr75
		}
		goto st0
tr75:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st40
	st40:
		if p++; p == pe {
			goto _test_eof40
		}
	st_case_40:
//line uri_parse.go:2497
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr76
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr76
			}
		default:
			goto tr76
		}
		goto st0
tr147:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st41
tr157:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st41
tr163:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
	goto st41
tr168:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st41
tr174:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st41
tr179:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st41
tr184:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
	goto st41
	st41:
		if p++; p == pe {
			goto _test_eof41
		}
	st_case_41:
//line uri_parse.go:2586
		switch data[p] {
		case 37:
			goto tr78
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 64 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr77
tr166:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st69
tr80:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st69
tr77:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st69
	st69:
		if p++; p == pe {
			goto _test_eof69
		}
	st_case_69:
//line uri_parse.go:2636
		switch data[p] {
		case 32:
			goto tr165
		case 37:
			goto st42
		case 38:
			goto tr168
		case 47:
			goto st0
		case 61:
			goto tr169
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr165
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] > 35:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] >= 34:
				goto st0
			}
		default:
			goto st0
		}
		goto tr166
tr165:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st70
tr171:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
	goto st70
	st70:
		if p++; p == pe {
			goto _test_eof70
		}
	st_case_70:
//line uri_parse.go:2705
		if data[p] == 32 {
			goto st70
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st70
		}
		goto st0
tr78:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
	goto st42
	st42:
		if p++; p == pe {
			goto _test_eof42
		}
	st_case_42:
//line uri_parse.go:2724
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr79
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr79
			}
		default:
			goto tr79
		}
		goto st0
tr79:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st43
	st43:
		if p++; p == pe {
			goto _test_eof43
		}
	st_case_43:
//line uri_parse.go:2747
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr80
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr80
			}
		default:
			goto tr80
		}
		goto st0
tr169:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
	goto st44
	st44:
		if p++; p == pe {
			goto _test_eof44
		}
	st_case_44:
//line uri_parse.go:2770
		switch data[p] {
		case 37:
			goto tr82
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			if 58 <= data[p] && data[p] <= 64 {
				goto st0
			}
		default:
			goto st0
		}
		goto tr81
tr81:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st71
tr172:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st71
tr84:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st71
	st71:
		if p++; p == pe {
			goto _test_eof71
		}
	st_case_71:
//line uri_parse.go:2818
		switch data[p] {
		case 32:
			goto tr171
		case 37:
			goto st45
		case 38:
			goto tr174
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 14:
			switch {
			case data[p] > 8:
				if 9 <= data[p] && data[p] <= 13 {
					goto tr171
				}
			default:
				goto st0
			}
		case data[p] > 31:
			switch {
			case data[p] > 35:
				if 58 <= data[p] && data[p] <= 64 {
					goto st0
				}
			case data[p] >= 34:
				goto st0
			}
		default:
			goto st0
		}
		goto tr172
tr82:
//line uri_parse.rl:37
 amt = 0                          
	goto st45
	st45:
		if p++; p == pe {
			goto _test_eof45
		}
	st_case_45:
//line uri_parse.go:2867
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr83
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr83
			}
		default:
			goto tr83
		}
		goto st0
tr83:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st46
	st46:
		if p++; p == pe {
			goto _test_eof46
		}
	st_case_46:
//line uri_parse.go:2890
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr84
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr84
			}
		default:
			goto tr84
		}
		goto st0
tr64:
//line uri_parse.rl:37
 amt = 0                          
	goto st47
	st47:
		if p++; p == pe {
			goto _test_eof47
		}
	st_case_47:
//line uri_parse.go:2913
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr85
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr85
			}
		default:
			goto tr85
		}
		goto st0
tr85:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st48
	st48:
		if p++; p == pe {
			goto _test_eof48
		}
	st_case_48:
//line uri_parse.go:2936
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr86
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr86
			}
		default:
			goto tr86
		}
		goto st0
tr145:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st72
	st72:
		if p++; p == pe {
			goto _test_eof72
		}
	st_case_72:
//line uri_parse.go:2959
		switch data[p] {
		case 32:
			goto tr175
		case 37:
			goto tr176
		case 59:
			goto tr178
		case 63:
			goto tr179
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr177
			}
		case data[p] >= 9:
			goto tr175
		}
		goto st0
tr176:
//line uri_parse.rl:37
 amt = 0                          
	goto st49
	st49:
		if p++; p == pe {
			goto _test_eof49
		}
	st_case_49:
//line uri_parse.go:2988
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr87
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr87
			}
		default:
			goto tr87
		}
		goto st0
tr87:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st50
	st50:
		if p++; p == pe {
			goto _test_eof50
		}
	st_case_50:
//line uri_parse.go:3011
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr88
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr88
			}
		default:
			goto tr88
		}
		goto st0
tr177:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st73
tr182:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st73
tr88:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st73
	st73:
		if p++; p == pe {
			goto _test_eof73
		}
	st_case_73:
//line uri_parse.go:3045
		switch data[p] {
		case 32:
			goto tr180
		case 37:
			goto st49
		case 59:
			goto tr183
		case 63:
			goto tr184
		}
		switch {
		case data[p] > 13:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr182
			}
		case data[p] >= 9:
			goto tr180
		}
		goto st0
	st51:
		if p++; p == pe {
			goto _test_eof51
		}
	st_case_51:
		switch data[p] {
		case 37:
			goto tr90
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto st0
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr91
				}
			case data[p] >= 59:
				goto st0
			}
		default:
			goto st0
		}
		goto tr89
tr91:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st52
tr89:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st52
tr92:
//line uri_parse.rl:38
 buf[amt] = data[p]; amt++             
	goto st52
tr94:
//line uri_parse.rl:39
 buf[amt] = data[p] + 0x20; amt++      
	goto st52
tr97:
//line uri_parse.rl:41
 hex += unhex(data[p])
                        buf[amt] = hex; amt++            
	goto st52
	st52:
		if p++; p == pe {
			goto _test_eof52
		}
	st_case_52:
//line uri_parse.go:3132
		switch data[p] {
		case 37:
			goto st53
		case 38:
			goto st0
		case 47:
			goto st0
		case 91:
			goto st0
		case 93:
			goto tr95
		case 127:
			goto st0
		}
		switch {
		case data[p] < 34:
			if data[p] <= 32 {
				goto st0
			}
		case data[p] > 35:
			switch {
			case data[p] > 64:
				if 65 <= data[p] && data[p] <= 90 {
					goto tr94
				}
			case data[p] >= 59:
				goto st0
			}
		default:
			goto st0
		}
		goto tr92
tr90:
//line uri_parse.rl:37
 amt = 0                          
	goto st53
	st53:
		if p++; p == pe {
			goto _test_eof53
		}
	st_case_53:
//line uri_parse.go:3174
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr96
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr96
			}
		default:
			goto tr96
		}
		goto st0
tr96:
//line uri_parse.rl:40
 hex = unhex(data[p]) * 16             
	goto st54
	st54:
		if p++; p == pe {
			goto _test_eof54
		}
	st_case_54:
//line uri_parse.go:3197
		switch {
		case data[p] < 65:
			if 48 <= data[p] && data[p] <= 57 {
				goto tr97
			}
		case data[p] > 70:
			if 97 <= data[p] && data[p] <= 102 {
				goto tr97
			}
		default:
			goto tr97
		}
		goto st0
tr95:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
	goto st74
	st74:
		if p++; p == pe {
			goto _test_eof74
		}
	st_case_74:
//line uri_parse.go:3220
		switch data[p] {
		case 32:
			goto st66
		case 58:
			goto st72
		case 59:
			goto st35
		case 63:
			goto st41
		}
		if 9 <= data[p] && data[p] <= 13 {
			goto st66
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
	_test_eof55: cs = 55; goto _test_eof
	_test_eof56: cs = 56; goto _test_eof
	_test_eof12: cs = 12; goto _test_eof
	_test_eof57: cs = 57; goto _test_eof
	_test_eof13: cs = 13; goto _test_eof
	_test_eof14: cs = 14; goto _test_eof
	_test_eof15: cs = 15; goto _test_eof
	_test_eof58: cs = 58; goto _test_eof
	_test_eof16: cs = 16; goto _test_eof
	_test_eof17: cs = 17; goto _test_eof
	_test_eof18: cs = 18; goto _test_eof
	_test_eof59: cs = 59; goto _test_eof
	_test_eof60: cs = 60; goto _test_eof
	_test_eof19: cs = 19; goto _test_eof
	_test_eof20: cs = 20; goto _test_eof
	_test_eof21: cs = 21; goto _test_eof
	_test_eof61: cs = 61; goto _test_eof
	_test_eof22: cs = 22; goto _test_eof
	_test_eof23: cs = 23; goto _test_eof
	_test_eof24: cs = 24; goto _test_eof
	_test_eof25: cs = 25; goto _test_eof
	_test_eof62: cs = 62; goto _test_eof
	_test_eof26: cs = 26; goto _test_eof
	_test_eof27: cs = 27; goto _test_eof
	_test_eof63: cs = 63; goto _test_eof
	_test_eof28: cs = 28; goto _test_eof
	_test_eof29: cs = 29; goto _test_eof
	_test_eof30: cs = 30; goto _test_eof
	_test_eof31: cs = 31; goto _test_eof
	_test_eof64: cs = 64; goto _test_eof
	_test_eof32: cs = 32; goto _test_eof
	_test_eof33: cs = 33; goto _test_eof
	_test_eof34: cs = 34; goto _test_eof
	_test_eof65: cs = 65; goto _test_eof
	_test_eof66: cs = 66; goto _test_eof
	_test_eof35: cs = 35; goto _test_eof
	_test_eof67: cs = 67; goto _test_eof
	_test_eof36: cs = 36; goto _test_eof
	_test_eof37: cs = 37; goto _test_eof
	_test_eof38: cs = 38; goto _test_eof
	_test_eof68: cs = 68; goto _test_eof
	_test_eof39: cs = 39; goto _test_eof
	_test_eof40: cs = 40; goto _test_eof
	_test_eof41: cs = 41; goto _test_eof
	_test_eof69: cs = 69; goto _test_eof
	_test_eof70: cs = 70; goto _test_eof
	_test_eof42: cs = 42; goto _test_eof
	_test_eof43: cs = 43; goto _test_eof
	_test_eof44: cs = 44; goto _test_eof
	_test_eof71: cs = 71; goto _test_eof
	_test_eof45: cs = 45; goto _test_eof
	_test_eof46: cs = 46; goto _test_eof
	_test_eof47: cs = 47; goto _test_eof
	_test_eof48: cs = 48; goto _test_eof
	_test_eof72: cs = 72; goto _test_eof
	_test_eof49: cs = 49; goto _test_eof
	_test_eof50: cs = 50; goto _test_eof
	_test_eof73: cs = 73; goto _test_eof
	_test_eof51: cs = 51; goto _test_eof
	_test_eof52: cs = 52; goto _test_eof
	_test_eof53: cs = 53; goto _test_eof
	_test_eof54: cs = 54; goto _test_eof
	_test_eof74: cs = 74; goto _test_eof

	_test_eof: {}
	if p == eof {
		switch cs {
		case 55, 65:
//line uri_parse.rl:45
 uri.Host = string(buf[0:amt])    
		case 63, 73:
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
		case 62, 72:
//line uri_parse.rl:37
 amt = 0                          
//line uri_parse.rl:50

      if amt > 0 {
        port, err := strconv.ParseUint(string(buf[0:amt]), 10, 16)
        if err != nil { goto fail }
        uri.Port = uint16(port)
      }
    
		case 57, 67:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
		case 59, 69:
//line uri_parse.rl:47
 b1 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
		case 58, 68:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:58

      if uri.Params == nil {
        uri.Params = Params{}
      }
      uri.Params[b1] = b2
    
		case 61, 71:
//line uri_parse.rl:48
 b2 = string(buf[0:amt]); amt = 0 
//line uri_parse.rl:65

      if uri.Headers == nil {
        uri.Headers = URIHeaders{}
      }
      uri.Headers[b1] = b2
    
//line uri_parse.go:3377
		}
	}

	_out: {}
	}

//line uri_parse.rl:128

  if cs < uri_first_final {
    if p == pe {
      return nil, errors.New(fmt.Sprintf("Unexpected EOF: %s", data))
    } else {
      return nil, errors.New(fmt.Sprintf("Error in URI at pos %d: %s", p, data))
    }
  }
  return uri, nil

fail:
  return nil, errors.New(fmt.Sprintf("Bad URI: %s", data))
}
