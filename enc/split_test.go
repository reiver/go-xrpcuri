package xrpcurienc_test

import (
	"testing"

	"net/url"

	"github.com/reiver/go-xrpcuri/enc"
)

func TestSplit(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedHost string
		ExpectedCollection string
		ExpectedQuery string
		ExpectedFragment string
	}{
		{
			URI: `xrpc:`,
		},
		{
			URI: `xrpC:`,
		},
		{
			URI: `xrPc:`,
		},
		{
			URI: `xrPC:`,
		},
		{
			URI: `xRpc:`,
		},
		{
			URI: `xRpC:`,
		},
		{
			URI: `xRPc:`,
		},
		{
			URI: `xRPC:`,
		},
		{
			URI: `Xrpc:`,
		},
		{
			URI: `XrpC:`,
		},
		{
			URI: `XrPc:`,
		},
		{
			URI: `XrPC:`,
		},
		{
			URI: `XRpc:`,
		},
		{
			URI: `XRpC:`,
		},
		{
			URI: `XRPc:`,
		},
		{
			URI: `XRPC:`,
		},



		{
			URI: `xrpc://`,
		},
		{
			URI: `xrpC://`,
		},
		{
			URI: `xrPc://`,
		},
		{
			URI: `xrPC://`,
		},
		{
			URI: `xRpc://`,
		},
		{
			URI: `xRpC://`,
		},
		{
			URI: `xRPc://`,
		},
		{
			URI: `xRPC://`,
		},
		{
			URI: `Xrpc://`,
		},
		{
			URI: `XrpC://`,
		},
		{
			URI: `XrPc://`,
		},
		{
			URI: `XrPC://`,
		},
		{
			URI: `XRpc://`,
		},
		{
			URI: `XRpC://`,
		},
		{
			URI: `XRPc://`,
		},
		{
			URI: `XRPC://`,
		},



		{
			URI:   `xrpc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xrpC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xrPc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xrPC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRpc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRpC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRPc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRPC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `Xrpc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XrpC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XrPc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XrPC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRpc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRpC://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRPc://Example.COM`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRPC://Example.COM`,
			ExpectedHost: "Example.COM",
		},



		{
			URI:   `xrpc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xrpC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xrPc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xrPC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRpc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRpC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRPc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `xRPC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `Xrpc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XrpC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XrPc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XrPC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRpc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRpC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRPc://Example.COM/`,
			ExpectedHost: "Example.COM",
		},
		{
			URI:   `XRPC://Example.COM/`,
			ExpectedHost: "Example.COM",
		},



		{
			URI:   `xrpc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xrpC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xrPc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xrPC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xRpc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xRpC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xRPc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `xRPC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `Xrpc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XrpC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XrPc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XrPC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XRpc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XRpC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XRPc://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},
		{
			URI:   `XRPC://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
		},



		{
			URI:   `xrpc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xrpC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xrPc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xrPC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xRpc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xRpC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xRPc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `xRPC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `Xrpc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XrpC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XrPc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XrPC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XRpc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XRpC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XRPc://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},
		{
			URI:   `XRPC://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost: "Example.COM",
			ExpectedCollection:       "APP.Cherry.fooBar",
			ExpectedQuery:                              "actor=JoeBlow&sort=desc",
		},



		{
			URI:            `xrpc://example.com/app.cherry.fooBar`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
		},
		{
			URI:            `xrpc://example.com/app.cherry.fooBar?actor=joeblow&sort=desc`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
		},
		{
			URI:            `xrpc://example.com/app.cherry.fooBar#wXyZ123`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedFragment:                                    "wXyZ123",
		},
		{
			URI:            `xrpc://example.com/app.cherry.fooBar?actor=joeblow&sort=desc#wXyZ123`,
			ExpectedHost:          "example.com",
			ExpectedCollection:                "app.cherry.fooBar",
			ExpectedQuery:                                       "actor=joeblow&sort=desc",
			ExpectedFragment:                                                            "wXyZ123",
		},



		{
			URI:            `xrpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xrPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `xRPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `Xrpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XrPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRpc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRpC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRPc://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},
		{
			URI:            `XRPC://host.example/once.twice.thrice.fource.someThing`,
			ExpectedHost:          "host.example",
			ExpectedCollection:                 "once.twice.thrice.fource.someThing",
		},



		{
			URI:            `xrpc://Host.EXAMPLE/apple.BANANA.Cherry.dAtE/wxyZ`,
			ExpectedHost:          "Host.EXAMPLE",
			ExpectedCollection:                 "apple.BANANA.Cherry.dAtE",
		},
	}

	for testNumber, test := range tests {

		actualHost, actualCollection, actualQuery, actualFragment, err := xrpcurienc.Split(test.URI)
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
			ExpectedError: `xrpcuri: URI "http://example.com" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "x",
			ExpectedError: `xrpcuri: URI "x" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xr",
			ExpectedError: `xrpcuri: URI "xr" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrp",
			ExpectedError: `xrpcuri: URI "xrp" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc",
			ExpectedError: `xrpcuri: URI "xrpc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "xrpc-",
			ExpectedError: `xrpcuri: URI "xrpc-" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-u",
			ExpectedError: `xrpcuri: URI "xrpc-u" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-un",
			ExpectedError: `xrpcuri: URI "xrpc-un" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-une",
			ExpectedError: `xrpcuri: URI "xrpc-une" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unen",
			ExpectedError: `xrpcuri: URI "xrpc-unen" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unenc",
			ExpectedError: `xrpcuri: URI "xrpc-unenc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unencr",
			ExpectedError: `xrpcuri: URI "xrpc-unencr" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unencry",
			ExpectedError: `xrpcuri: URI "xrpc-unencry" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unencryp",
			ExpectedError: `xrpcuri: URI "xrpc-unencryp" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unencrypt",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypt" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unencrypte",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypte" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpc-unencrypted",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypted" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "xrpc-unencrypted:",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypted:" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: ":",
			ExpectedError: `xrpcuri: URI ":" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "x:",
			ExpectedError: `xrpcuri: URI "x:" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xr:",
			ExpectedError: `xrpcuri: URI "xr:" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrp:",
			ExpectedError: `xrpcuri: URI "xrp:" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "xrpc-:",
			ExpectedError: `xrpcuri: URI "xrpc-:" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
	}

	for testNumber, test := range tests {

		_, _, _, _, err := xrpcurienc.Split(test.URI)
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
