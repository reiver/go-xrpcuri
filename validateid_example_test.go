package xrpcuri_test

import (
        "fmt"

        "github.com/reiver/go-xrpcuri"
)

func ExampleValidateID() {

        var id string = "COM.Example.fooBar"

        err := xrpcuri.ValidateID(id)

        fmt.Printf("id (i.e., NSID): %s\n", id)
        fmt.Printf("validation error: %s\n", err)

        // Output:
	// id (i.e., NSID): COM.Example.fooBar
        // validation error: nsid: character №0 ('C') (U+0043) of domain-authority part №0 ("COM") of domain-authority ("COM.Example") of nsid ("COM.Example.fooBar") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')
}
