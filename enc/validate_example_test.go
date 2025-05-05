package xrpcurienc_test

import (
        "fmt"

        "github.com/reiver/go-xrpcuri/enc"
)

func ExampleValidate() {

        var xrpcURI string = "xrpc://archive.org/COM.Example.fooBar?actor=JoeBlow"

        err := xrpcurienc.Validate(xrpcURI)

        fmt.Printf("XRPC-URI: %s\n", xrpcURI)
        fmt.Printf("validation error: %s\n", err)

        // Output:
        // XRPC-URI: xrpc://archive.org/COM.Example.fooBar?actor=JoeBlow
        // validation error: xrpcuri: XRPC-URI "xrpc://archive.org/COM.Example.fooBar?actor=JoeBlow" has an id "COM.Example.fooBar" that is not a valid NSID: nsid: character №0 ('C') (U+0043) of domain-authority part №0 ("COM") of domain-authority ("COM.Example") of nsid ("COM.Example.fooBar") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')
}
