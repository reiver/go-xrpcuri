package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestURLString(t *testing.T) {

	tests := []struct{
		URL xrpcuri.URL
		Expected string
	}{
		{
			URL: xrpcuri.URL{
				Host:     "example.com",
				NSID:                "net.something.fooBar",
			},
			Expected: "xrpc://example.com/net.something.fooBar",
		},
		{
			URL: xrpcuri.URL{
				Unencrypted: true,
				Host:    "example.com",
				NSID:                            "net.something.fooBar",
			},
			Expected: "xrpc-unencrypted://example.com/net.something.fooBar",
		},



		{
			URL: xrpcuri.URL{
				Host:     "public.api.bsky.app",
				NSID:                        "app.bsky.actor.getProfile",
			},
			Expected: "xrpc://public.api.bsky.app/app.bsky.actor.getProfile",
		},
		{
			URL: xrpcuri.URL{
				Unencrypted: true,
				Host:                "public.api.bsky.app",
				NSID:                                    "app.bsky.actor.getProfile",
			},
			Expected: "xrpc-unencrypted://public.api.bsky.app/app.bsky.actor.getProfile",
		},



		{
			URL: xrpcuri.URL{
				Host:     "public.api.bsky.app",
				NSID:                        "app.bsky.actor.getProfile",
				Query:                                                 "actor=reiver.bsky.social",
			},
			Expected: "xrpc://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social",
		},
		{
			URL: xrpcuri.URL{
				Unencrypted: true,
				Host:                "public.api.bsky.app",
				NSID:                                    "app.bsky.actor.getProfile",
				Query:                                                             "actor=reiver.bsky.social",
			},
			Expected: "xrpc-unencrypted://public.api.bsky.app/app.bsky.actor.getProfile?actor=reiver.bsky.social",
		},
	}

	for testNumber, test := range tests {

		actual := test.URL.String()

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual resolved-url is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("URL: %#v", test.URL)
			continue
		}
	}
}
