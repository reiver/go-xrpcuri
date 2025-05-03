package xrpcuri_internal

import (
	"github.com/reiver/go-nsid"
)

// NormalizeID returns the normalized form of an XRPC-URI 'id', as defined in:
// https://atproto.com/specs/xrpc
//
// An XRPC-URI 'id' is an NSID, as defined at:
// https://atproto.com/specs/nsid
//
// The XRPC-URI 'id' (such as "com.example.fooBar") is part of an XRPC-URI (such as "xrpc://archive.org/com.example.fooBar").
//
// An example of a non-normalized XRPC-URI 'id' would be "COM.Example.fooBar".
// Normalizing that non-normalized XRPC-URI 'id' would result in "com.example.fooBar".
//
// Note that if you want to normalize a whole XRPC-URI rather than just a 'id', then instead use [Normalize].
func NormalizeID(value string) string {
	return nsid.Normalize(value)
}
