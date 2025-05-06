package xrpcuri_test

import (
	"fmt"
	"net/url"

	"github.com/reiver/go-xrpcuri"
)

func ExampleJoin() {

	var authority string = "public.api.bsky.app"
	var id        string = "app.bsky.actor.getProfile"
	var actor     string = "reiver.bsky.social"

	query := "actor=" + url.QueryEscape(actor)

	uri := xrpcuri.Join(authority, id, query, "")

	fmt.Printf("authority:   %s\n", authority)
	fmt.Printf("id:          %s\n", id)
	fmt.Printf("actor:       %s\n", actor)
	fmt.Println()
	fmt.Printf("query: %s\n", query)
	fmt.Println()
	fmt.Printf("uri: %s\n", uri)

	// Output:
	// authority:   public.api.bsky.app
	// id:          app.bsky.actor.getProfile
	// actor:       reiver.bsky.social
	//
	// query: actor=reiver.bsky.social
	//
	// uri: xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social
}
