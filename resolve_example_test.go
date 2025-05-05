package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleResolve() {

	var xrpcURI string = "xrpc://example.com/com.atproto.repo.listRecords"

	httpURI, err := xrpcuri.Resolve(xrpcURI, "query")
	if nil != err {
		fmt.Printf("ERROR: could not resolve XRPC-URI %q: %s\n", xrpcURI, err)
		return
	}

	fmt.Printf("http-uri: %s\n", httpURI)

	// Output:
	// http-uri: https://example.com/xrpc/com.atproto.repo.listRecords
}
