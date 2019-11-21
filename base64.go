package tool

import "encoding/base64"

var Base64Coder = base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")

func Base64Encode(src string) string {
	return Base64Coder.EncodeToString([]byte(src))
}

func Base64Decode(src string) string {
	data, _ := Base64Coder.DecodeString(src)
	return string(data)
}
