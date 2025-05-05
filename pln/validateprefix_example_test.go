package xrpcuripln_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri/pln"
)

func ExampleValidatePrefixUnencrypted() {

	var uri string = "xrpc-unencrypted:example.com"

	err := xrpcuripln.ValidatePrefix(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted:example.com" is not valid because it does not have "//" after "xrpc-unencrypted:"
}
