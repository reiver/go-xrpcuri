package xrpcurienc_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri/enc"
)

func ExampleValidateScheme() {

	var uri string = "http://example.com/once/twice/thrice/fource.html"

	err := xrpcurienc.ValidateScheme(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: xrpcuri: URI "http://example.com/once/twice/thrice/fource.html" is not an XRPC-URI because it does not begin with "xrpc:"
}
