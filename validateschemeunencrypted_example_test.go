package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleValidateSchemeUnencrypted() {

	var uri string = "http://example.com/once/twice/thrice/fource.html"

	err := xrpcuri.ValidateSchemeUnencrypted(uri)

	fmt.Printf("error: %s\n", err)

	// Output:
	// error: xrpcuri: URI "http://example.com/once/twice/thrice/fource.html" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"
}
