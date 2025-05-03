package xrpcuri_internal

import (
	"strings"

	"golang.org/x/net/idna"
)

func Join(scheme string, authority string, id string, query string, fragment string) string {
	authority  = NormalizeAuthority(authority)
	id         = NormalizeID(id)

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, scheme...)
	p = append(p, "://"...)

	if "" != authority {
		{
			punycode, err := idna.ToASCII(authority)
			if nil == err {
				authority = punycode
			}
		}

		p = append(p, encodeAtSign(encodeSolidus(encodeQuestionMark(encodeNumberSign(encodePercentSign(authority)))))...)
	}

	if  "" != id {
		p = append(p, "/"...)

		if "" != id {
			p = append(p, encodeSolidus(encodeQuestionMark(encodeNumberSign(encodePercentSign(id))))...)
		}
	}

	if "" != query {
		p = append(p, "?"...)
		p = append(p, encodeNumberSign(encodePercentSign(query))...)
	}

	if "" != fragment {
		p = append(p, "#"...)
		p = append(p, fragment...)
	}

	return string(p)
}

func encodeAtSign(str string) string {
	return strings.ReplaceAll(str,  "@", "%40")
}

func encodeNumberSign(str string) string {
	return strings.ReplaceAll(str,  "#", "%23")
}

func encodePercentSign(str string) string {
	return strings.ReplaceAll(str,  "%", "%25")
}

func encodeQuestionMark(str string) string {
	return strings.ReplaceAll(str,  "?", "%3F")
}

func encodeSolidus(str string) string {
	return strings.ReplaceAll(str,  "/", "%2F")
}
