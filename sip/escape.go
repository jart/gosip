package sip

// escapeUser escapes a URI user, which can't use quoting.
func escapeUser(s []byte) []byte {
	return escape(s, userc)
}

// escapePass escapes a URI password, which can't use quoting.
func escapePass(s []byte) []byte {
	return escape(s, passc)
}

// escapeParam escapes a URI parameter, which can't use quoting.
func escapeParam(s []byte) []byte {
	return escape(s, paramc)
}

// escapeHeader escapes a URI header, which can't use quoting.
func escapeHeader(s []byte) []byte {
	return escape(s, headerc)
}

// escape provides arbitrary URI escaping.
func escape(s []byte, p func(byte) bool) []byte {
	hc := 0
	for i := 0; i < len(s); i++ {
		if !p(s[i]) {
			hc++
		}
	}
	if hc == 0 {
		return s
	}
	res := make([]byte, len(s)+2*hc)
	j := 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if p(c) {
			res[j] = c
			j++
		} else {
			res[j] = '%'
			res[j+1] = hexChars[c>>4]
			res[j+2] = hexChars[c%16]
			j += 3
		}
	}
	return res
}
