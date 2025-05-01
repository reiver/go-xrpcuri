package xrpcuri

import (
	"github.com/reiver/go-erorr"
)

const (
	errEmptyHost = erorr.Error("xrpcuri: empty host")
	errEmptyURI  = erorr.Error("xrpcuri: empty uri")
	errNilURL    = erorr.Error("xrpcuri: nil url")
)
