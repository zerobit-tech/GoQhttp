package xmlutils

import (
	"bytes"
	"strings"
)

var removeNonUTF = func(r rune) rune {
	//fmt.Println("valid rune ", r, utf8.ValidRune(r), IsInCharacterRange(r))
	if !IsInCharacterRange(r) {
		return -1
	}
	return r
}

// RemoveNonUTF8Strings removes strings that isn't UTF-8 encoded
func RemoveNonUTF8Strings(stringX string) string {
	return strings.Map(removeNonUTF, stringX)
}

// RemoveNonUTF8Bytes removes bytes that isn't UTF-8 encoded
func RemoveNonUTF8Bytes(dataX []byte) []byte {
	return bytes.Map(removeNonUTF, dataX)
}
