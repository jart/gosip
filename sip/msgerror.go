package sip

import (
	"bytes"
	"fmt"
)

type MsgIncompleteError struct {
	Msg []byte
}

func (err MsgIncompleteError) Error() string {
	return fmt.Sprintf("Incomplete SIP message:\n%s", err.Msg)
}

type MsgParseError struct {
	Msg    []byte
	Offset int
}

func (err MsgParseError) Error() string {
	lines := bytes.Split(err.Msg, []byte("\r\n"))
	var b bytes.Buffer
	o := 0
	line := 0
	lineOffset := 0
	for i := 0; i < len(lines); i++ {
		b.Write(lines[i])
		if o <= err.Offset && err.Offset < o+len(lines[i]) {
			line = i + 1
			lineOffset = err.Offset - o
			b.WriteByte('\n')
			for j := 0; j < lineOffset; j++ {
				b.WriteByte(' ')
			}
			b.WriteByte('^')
		}
		o += len(lines[i]) + 2
		if i < len(lines)-1 {
			b.WriteByte('\n')
		}
	}
	return fmt.Sprintf("Error in SIP message at line %d offset %d:\n%s", line, lineOffset, b.String())
}
