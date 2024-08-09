package winrm

import (
	"encoding/base64"
	"fmt"
	"unicode/utf16"
)

func Powershell(psCmd string) string {
	// Encode the string to UTF-16LE
	encoded := utf16.Encode([]rune(psCmd))

	// Convert UTF-16LE encoded []uint16 to []byte
	var buf []byte
	for _, r := range encoded {
		buf = append(buf, byte(r&0xff), byte(r>>8))
	}

	// Base64 encode the UTF-16LE bytes
	encodedCmd := base64.StdEncoding.EncodeToString(buf)

	// Create the powershell.exe command line to execute the script
	return fmt.Sprintf("powershell.exe -EncodedCommand %s", encodedCmd)
}
