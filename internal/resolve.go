package xrpcuri_internal

import (
	"github.com/reiver/go-erorr"
)

const (
	errNilFunc = erorr.Error("xrpcuri: nil func")
)

func Resolve(uri string, fn func(string)(string,string,string,string,error), scheme string) (string, error) {

	if nil == fn {
		const nada string = ""
		return nada, errNilFunc
	}

	authority, id, query, fragment, err := fn(uri)
	if nil != err {
		const nada string = ""
		return nada, err
	}

	var buffer [256]byte
	var p []byte = buffer[0:0]

	p = append(p, scheme...)
	p = append(p, "://"...)

	p = append(p, authority...)

	p = append(p, "/xrpc/"...)
	p = append(p, id...)

	if "" != query {
		p = append(p, '?')
		p = append(p, query...)
	}

	if "" != fragment {
		p = append(p, '#')
		p = append(p, fragment...)
	}

	return string(p), nil
}
