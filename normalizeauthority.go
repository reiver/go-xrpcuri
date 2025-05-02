package xrpcuri

import (
	"strings"
)

// NormalizeAuthority returns the normalized form of an XRPC-URI 'authority'.
//
// The XRPC-URI 'authority' (such as "example.com") is part of an XRPC-URI (such as "xrpc://example.com").
//
// In simple language, you can think of an XRPC-URI 'authority' as being either an Internet domain-name (ex: "example.com").
//
// An example of a non-normalized XRPC-URI 'authority' would be "Example.COM".
// Normalizing that non-normalized XRPC-URI 'authority' would result in "example.com".
//
// NormalizeAuthority will leave any potential 'userinfo' in the 'authority' as is.
//
// Note that if you want to normalize a whole XRPC-URI rather than just an 'authority', then instead use [Normalize].
func NormalizeAuthority(value string) string {
	var str string = value

	var buffer [256]byte
	var p []byte = buffer[0:0]

	{
		var index int = strings.Index(str, "@")
		if 0 <= index {
			var userinfo string = str[:index]
			p = append(p, userinfo...)
			p = append(p, '@')

			str = str[index+1:]
		}
	}

	var length int = len(str)

	for i:=0; i<length; i++ {
		var b byte = str[i]

		switch {
		case 'A' <= b && b <= 'Z':
			p = append(p, b-'A'+'a')
		default:
			p = append(p, b)
		}
	}

	return string(p)
}
