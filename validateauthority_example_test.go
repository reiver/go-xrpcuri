package xrpcuri_test

import (
        "fmt"

        "github.com/reiver/go-xrpcuri"
)

func ExampleValidateAuthority() {

        var authority string = "user:pass@archive.org"

        err := xrpcuri.ValidateAuthority(authority)

        fmt.Printf("authority: %s\n", authority)
        fmt.Printf("validation error: %s\n", err)

        // Output:
        // authority: user:pass@archive.org
        // validation error: xrpcuri: authority may not have an "@" in it
}
