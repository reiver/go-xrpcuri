package xrpcuripln_test

import (
        "fmt"

        "github.com/reiver/go-xrpcuri/pln"
)

func ExampleValidateUnencrypted() {

        var xrpcURI string = "xrpc-unencrypted://archive.org/COM.Example.fooBar?actor=JoeBlow"

        err := xrpcuripln.Validate(xrpcURI)

        fmt.Printf("XRPC-unencrypted-URI: %s\n", xrpcURI)
        fmt.Printf("validation error: %s\n", err)

        // Output:
        // XRPC-unencrypted-URI: xrpc-unencrypted://archive.org/COM.Example.fooBar?actor=JoeBlow
        // validation error: xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://archive.org/COM.Example.fooBar?actor=JoeBlow" has an id "COM.Example.fooBar" that is not a valid NSID: nsid: character №0 ('C') (U+0043) of domain-authority part №0 ("COM") of domain-authority ("COM.Example") of nsid ("COM.Example.fooBar") is not a digit ('0'-'9'), lower-case letter ('a'-'z'), or a hyphen ('-')
}
