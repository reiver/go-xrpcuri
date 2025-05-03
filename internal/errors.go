package xrpcuri_internal

import (
	"github.com/reiver/go-erorr"
)

const (
	ErrEmptyURI = erorr.Error("xrpcuri: empty uri")
	ErrNilURL   = erorr.Error("xrpcuri: nil url")
)
