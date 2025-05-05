package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleNormalize() {

	var id string = "xrpcuri://Host.EXAMPLE/COM.Example.fooBar"

	normalizedID := xrpcuri.NormalizeID(id)

	fmt.Printf("original xrpc-uri:   %s\n", id)
	fmt.Printf("normalized xrpc-uri: %s\n", normalizedID)

	// Output:
	// original xrpc-uri:   xrpcuri://Host.EXAMPLE/COM.Example.fooBar
	// normalized xrpc-uri: xrpcuri://host.example/com.example.fooBar
}
