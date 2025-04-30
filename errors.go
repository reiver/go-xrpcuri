package xrpcuri

import (
	"github.com/reiver/go-erorr"
)

const (
	errEmptyHost = erorr.Error("xrpcuri: empty host")
	errNilURL    = erorr.Error("xrpcuri: nil url")
)
