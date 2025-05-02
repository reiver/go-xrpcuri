package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestJoinUnencrypted(t *testing.T) {

	tests := []struct{
		Host string
		Collection string
		Query string
		Fragment string
		Expected string
	}{
		{
			Host:                        "example.com",
			Collection:                              "app.cherry.fooBar",
			Expected: `xrpc-unencrypted://example.com/app.cherry.fooBar`,
		},
		{
			Host:                        "example.com",
			Collection:                              "app.cherry.fooBar",
			Query:                                                     "actor=joeblow&sort=desc",
			Expected: `xrpc-unencrypted://example.com/app.cherry.fooBar?actor=joeblow&sort=desc`,
		},
		{
			Host:                        "example.com",
			Collection:                              "app.cherry.fooBar",
			Fragment:                                                  "wXyZ123",
			Expected: `xrpc-unencrypted://example.com/app.cherry.fooBar#wXyZ123`,
		},
		{
			Host:                        "example.com",
			Collection:                              "app.cherry.fooBar",
			Query:                                                     "actor=joeblow&sort=desc",
			Fragment:                                                                          "wXyZ123",
			Expected: `xrpc-unencrypted://example.com/app.cherry.fooBar?actor=joeblow&sort=desc#wXyZ123`,
		},



		{
			Host:                        "host.example",
			Collection:                               "once.twice.thrice.fource.someThing",
			Expected: `xrpc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
		},
	}

	for testNumber, test := range tests {

		actual := xrpcuri.JoinUnencrypted(test.Host, test.Collection, test.Query, test.Fragment)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual XRPC-URI is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("HOST:       %q", test.Host)
			t.Logf("COLLECTION: %q", test.Collection)
			t.Logf("QUERY:      %q", test.Query)
			t.Logf("FRAGMENT:   %q", test.Fragment)
			continue
		}
	}
}
