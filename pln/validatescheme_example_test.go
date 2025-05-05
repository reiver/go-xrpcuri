package xrpcuripln_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri/pln"
)

func ExampleValidateScheme() {

	var uri string = "http://example.com/once/twice/thrice/fource.html"

	err := xrpcuripln.ValidateScheme(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: xrpcuri: URI "http://example.com/once/twice/thrice/fource.html" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"
}
