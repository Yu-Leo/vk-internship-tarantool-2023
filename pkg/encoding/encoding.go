package encoding

import (
	"encoding/base64"
)

func Encode(input string) string {
	return base64.StdEncoding.EncodeToString([]byte(input))
}

func Decode(input string) (string, error) {
	data, err := base64.StdEncoding.DecodeString(input)
	return string(data), err
}
