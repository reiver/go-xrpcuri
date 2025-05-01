package xrpcuri_test

import (
	"testing"

	"net/url"

	"github.com/reiver/go-xrpcuri"
)

func TestSplit(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedScheme string
		ExpectedHost string
		ExpectedCollection string
		ExpectedQuery string
		ExpectedFragment string
	}{
		{
			URI:            `xrpc://example.com/app.cherry.fooBar`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
		},
		{
			URI:            `xrpc://example.com/app.cherry.fooBar?actor=joeblow&sort=desc`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
		},
		{
			URI:            `xrpc://example.com/app.cherry.fooBar#wXyZ123`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedFragment:                                    "wXyZ123",
		},
		{
			URI:            `xrpc://example.com/app.cherry.fooBar?actor=joeblow&sort=desc#wXyZ123`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
			ExpectedFragment:                                                            "wXyZ123",
		},



		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar`,
			ExpectedScheme: "xrpc-unencrypted",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
		},
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar?actor=joeblow&sort=desc`,
			ExpectedScheme: "xrpc-unencrypted",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
		},
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar#wXyZ123`,
			ExpectedScheme: "xrpc-unencrypted",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedFragment:                                    "wXyZ123",
		},
		{
			URI:            `xrpc-unencrypted://example.com/app.cherry.fooBar?actor=joeblow&sort=desc#wXyZ123`,
			ExpectedScheme: "xrpc-unencrypted",
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
			ExpectedFragment:                                                            "wXyZ123",
		},



		{
			URI:            `xrpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `Xrpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedScheme: "xrpc",
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
	}

	for testNumber, test := range tests {

		actualScheme, actualHost, actualCollection, actualQuery, actualFragment, err := xrpcuri.Split(test.URI)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}

		{
			expected := test.ExpectedScheme
			actual   :=        actualScheme

			if expected != actual {
				t.Errorf("For test #%d, the actual 'scheme' is not what was expected.", testNumber)
				t.Logf("EXPECTED-SCHEME: %s", expected)
				t.Logf("  ACTUAL-SCHEME: %s", actual)
				t.Logf("URI: %s", test.URI)
				continue
			}
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

func TestSplit_fail(t *testing.T) {

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
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was "http"`,
		},



		{
			URI: "x",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xr",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrp",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},



		{
			URI: "xrpc-",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-u",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-un",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-une",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unen",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unenc",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unencr",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unencry",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unencryp",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unencrypt",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unencrypte",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "xrpc-unencrypted",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},



		{
			URI: ":",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was ""`,
		},
		{
			URI: "x:",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was "x"`,
		},
		{
			URI: "xr:",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was "xr"`,
		},
		{
			URI: "xrp:",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was "xrp"`,
		},



		{
			URI: "xrpc-:",
			ExpectedError: `xrpc: expected scheme to be "xrpc" or "xrpc-unencrypted" but was "xrpc-"`,
		},
	}

	for testNumber, test := range tests {

		_, _, _, _, _, err := xrpcuri.Split(test.URI)
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
