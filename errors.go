package xrpcuri

import (
	"github.com/reiver/go-erorr"
)

const (
	errEmptyURI = erorr.Error("xrpcuri: empty uri")
	errNilURL   = erorr.Error("xrpcuri: nil url")
)
