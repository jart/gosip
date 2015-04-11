// Charset implementation using four int64 bitmasks
//
// Each charset mask is 32 bytes, which fits into a 64 byte cache line.

package sip

var tokencMask [4]uint64
var qdtextcMask [4]uint64

func init() {
	charsetAddRange(&tokencMask, 'a', 'z')
	charsetAddRange(&tokencMask, 'A', 'Z')
	charsetAddRange(&tokencMask, '0', '9')
	charsetAdd(&tokencMask, '-')
	charsetAdd(&tokencMask, '.')
	charsetAdd(&tokencMask, '!')
	charsetAdd(&tokencMask, '%')
	charsetAdd(&tokencMask, '*')
	charsetAdd(&tokencMask, '_')
	charsetAdd(&tokencMask, '+')
	charsetAdd(&tokencMask, '`')
	charsetAdd(&tokencMask, '\'')
	charsetAdd(&tokencMask, '~')

	charsetAdd(&qdtextcMask, '\r')
	charsetAdd(&qdtextcMask, '\n')
	charsetAdd(&qdtextcMask, '\t')
	charsetAdd(&qdtextcMask, ' ')
	charsetAdd(&qdtextcMask, '!')
	charsetAddRange(&qdtextcMask, 0x23, 0x5B)
	charsetAddRange(&qdtextcMask, 0x5D, 0x7E)
}

func tokenc(c byte) bool {
	return charsetContains(&tokencMask, c)
}

func qdtextc(c byte) bool {
	return charsetContains(&qdtextcMask, c)
}

func qdtextesc(c byte) bool {
	return 0x00 <= c && c <= 0x09 ||
		0x0B <= c && c <= 0x0C ||
		0x0E <= c && c <= 0x7F
}

func whitespacec(c byte) bool {
	return c == ' ' || c == '\t' || c == '\r' || c == '\n'
}

func charsetContains(mask *[4]uint64, i byte) bool {
	return mask[i/64]&(1<<(i%64)) != 0
}

func charsetAdd(mask *[4]uint64, i byte) {
	mask[i/64] |= 1 << (i % 64)
}

func charsetAddRange(mask *[4]uint64, a, b byte) {
	for i := a; i <= b; i++ {
		charsetAdd(mask, i)
	}
	// var m uint64
	// i := a
	// j := i / 64
	// for i <= b {
	// 	m &= 1 << (i % 64)
	// 	i++
	// 	if i%64 == 0 {
	// 		mask[j] &= m
	// 		j = i / 64
	// 		m = 0
	// 	}
	// }
	// if m != 0 {
	// 	mask[j] |= m
	// }
}
