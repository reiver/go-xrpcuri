package xrpcuri_test

import (
	"testing"

	"net/url"

	"github.com/reiver/go-xrpcuri"
)

func TestSplitUnencrypted(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedHost string
		ExpectedCollection string
		ExpectedQuery string
		ExpectedFragment string
	}{
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
		},
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar?actor=joeblow&sort=desc`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
		},
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar#wXyZ123`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedFragment:                                    "wXyZ123",
		},
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar?actor=joeblow&sort=desc#wXyZ123`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
			ExpectedFragment:                                                            "wXyZ123",
		},



		{
			URI:            `xrpc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrpC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrPc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrPC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRpc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRpC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRPc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRPC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `Xrpc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrpC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrPc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrPC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRpc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRpC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRPc-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRPC-unencrypted://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
	}

	for testNumber, test := range tests {

		actualHost, actualCollection, actualQuery, actualFragment, err := xrpcuri.SplitUnencrypted(test.URI)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}

		{
			expected := test.ExpectedHost
			actual   :=        actualHost

			if expected != actual {
				t.Errorf("For test #%d, the actual 'host' is not what was expected.", testNumber)
				t.Logf("EXPECTED-HOST: %s", expected)
				t.Logf("  ACTUAL-HOST: %s", actual)
				t.Logf("URI: %s", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedCollection
			actual   :=        actualCollection

			if expected != actual {
				t.Errorf("For test #%d, the actual 'host' is not what was expected.", testNumber)
				t.Logf("EXPECTED-COLLECTION: %s", expected)
				t.Logf("  ACTUAL-COLLECTION: %s", actual)
				t.Logf("URI: %s", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedQuery
			actual   :=        actualQuery

			if expected != actual {
				t.Errorf("For test #%d, the actual 'query' is not what was expected.", testNumber)
				t.Logf("EXPECTED-QUERY: %s", expected)
				t.Logf("  ACTUAL-QUERY: %s", actual)
				t.Logf("URI: %s", test.URI)
				continue
			}
		}

		{
			expected := test.ExpectedFragment
			actual   :=        actualFragment

			if expected != actual {
				t.Errorf("For test #%d, the actual 'fragment' is not what was expected.", testNumber)
				t.Logf("EXPECTED-FRAGMENT: %s", expected)
				t.Logf("  ACTUAL-FRAGMENT: %s", actual)
				t.Logf("URI: %s", test.URI)
				continue
			}
		}
	}
}

func TestSplitUnencrypted_fail(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: `xrpcuri: empty uri`,
		},



		{
			URI: "http://example.com",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was "http"`,
		},



		{
			URI: "x",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xr",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrp",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},



		{
			URI: "xrpc:",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was "xrpc"`,
		},



		{
			URI: "xrpc-",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-u",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-un",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-une",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unen",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unenc",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unencr",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unencry",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unencryp",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unencrypt",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unencrypte",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "xrpc-unencrypted",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},



		{
			URI: ":",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was ""`,
		},
		{
			URI: "x:",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was "x"`,
		},
		{
			URI: "xr:",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was "xr"`,
		},
		{
			URI: "xrp:",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was "xrp"`,
		},



		{
			URI: "xrpc-:",
			ExpectedError: `xrpc: expected scheme to be "xrpc-unencrypted" but actually was "xrpc-"`,
		},
	}

	for testNumber, test := range tests {

		_, _, _, _, err := xrpcuri.SplitUnencrypted(test.URI)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("URI: %s", test.URI)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED-ERROR: %s", expected)
			t.Logf("  ACTUAL-ERROR: %s", actual)
			t.Logf("ERROR-TYPE: %T", err)
			if e, casted := err.(*url.Error); casted {
				t.Logf("ERROR-OP: %q", e.Op)
				t.Logf("ERROR-URL: %q", e.URL)
			}
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}
