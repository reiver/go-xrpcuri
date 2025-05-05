package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleValidatePrefix() {

	var uri string = "xrpc:example.com"

	err := xrpcuri.ValidatePrefix(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: xrpcuri: URI "xrpc:example.com" is not an XRPC-URI or an XRPC-unencrypted-URI because it does not begin with "xrpc://" or "xrpc-unencrypted://"
}
