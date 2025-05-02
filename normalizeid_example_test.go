package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleNormalizeID() {

	var id string = "COM.Example.fooBar"

	normalizedID := xrpcuri.NormalizeID(id)

	fmt.Printf("original id:   %s\n", id)
	fmt.Printf("normalized id: %s\n", normalizedID)

	// Output:
	// original id:   COM.Example.fooBar
	// normalized id: com.example.fooBar
}
