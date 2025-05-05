package xrpcurienc_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri/enc"
)

func ExampleValidatePrefix() {

	var uri string = "xrpc:example.com"

	err := xrpcurienc.ValidatePrefix(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: xrpcuri: XRPC-URI "xrpc:example.com" is not valid because it does not have "//" after "xrpc:"
}
