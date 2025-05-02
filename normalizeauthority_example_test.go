package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleNormalizeAuthority() {

	var authority string = "Example.COM"

	normalizedAuthority := xrpcuri.NormalizeAuthority(authority)

	fmt.Printf("original authority:   %s\n", authority)
	fmt.Printf("normalized authority: %s\n", normalizedAuthority)

	// Output:
	// original authority:   Example.COM
	// normalized authority: example.com
}
