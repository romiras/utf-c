package utfc

import (
	"strconv"
	"testing"
)

var testStrings []string = []string{
	"Tiếng Việt",
}

func hexString(buf []byte) string {
	s := ""
	for _, v := range buf {
		if v < 16 {
			s += "0"
		}
		s += strconv.FormatInt(int64(v), 16)
		s += " "
	}
	return s
}

const MaxLength = 1000

func TestUTFC(t *testing.T) {
	for _, test := range testStrings {
		name := test
		if len(name) > MaxLength {
			name = name[:MaxLength] + "…"
		}

		t.Run(name, func(t *testing.T) {
			utfc := Encode(test)
			ctrl := Decode(utfc)
			if ctrl != test {
				t.Errorf("String '%v' decoded back as '%v', bytes: %v", test, ctrl, hexString(utfc))
			}
		})
	}
}
