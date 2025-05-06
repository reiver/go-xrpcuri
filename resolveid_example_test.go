package xrpcuri_test

import (
	"fmt"

	"github.com/reiver/go-xrpcuri"
)

func ExampleResolveID() {

	var id string = "app.bsky.actor.getProfile"

	path := xrpcuri.ResolveID(id)

	fmt.Printf("id:         %s\n", id)
	fmt.Printf("path: %s\n", path)

	// Output:
	// id:         app.bsky.actor.getProfile
	// path: /xrpc/app.bsky.actor.getProfile
}
