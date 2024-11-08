package xmlutils

import (
	"encoding/xml"
	"io"
	"strings"
)

func IsValid(input string) bool {

	err := xml.Unmarshal([]byte(input), new(interface{}))
	if err != nil {
		return false
	}

	decoder := xml.NewDecoder(strings.NewReader(input))

	val := new(interface{})
	for {
		err := decoder.Decode(val)
		if err != nil {
			return err == io.EOF
		}
	}
}

// Decide whether the given rune is in the XML Character Range, per
// the Char production of https://www.xml.com/axml/testaxml.htm,
// Section 2.2 Characters.
func IsInCharacterRange(r rune) (inrange bool) {
	return r == 0x09 ||
		r == 0x0A ||
		r == 0x0D ||
		r >= 0x20 && r <= 0xD7FF ||
		r >= 0xE000 && r <= 0xFFFD ||
		r >= 0x10000 && r <= 0x10FFFF
}
