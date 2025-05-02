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
	// error: xrpcuri: XRPC-URI "xrpc:example.com" is not valid because it does not have "//" after "xrpc:"
}
