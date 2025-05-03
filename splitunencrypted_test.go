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
			URI: `xrpc-unencrypted:`,
		},
		{
			URI: `xrpC-unencrypted:`,
		},
		{
			URI: `xrPc-unencrypted:`,
		},
		{
			URI: `xrPC-unencrypted:`,
		},
		{
			URI: `xRpc-unencrypted:`,
		},
		{
			URI: `xRpC-unencrypted:`,
		},
		{
			URI: `xRPc-unencrypted:`,
		},
		{
			URI: `xRPC-unencrypted:`,
		},
		{
			URI: `Xrpc-unencrypted:`,
		},
		{
			URI: `XrpC-unencrypted:`,
		},
		{
			URI: `XrPc-unencrypted:`,
		},
		{
			URI: `XrPC-unencrypted:`,
		},
		{
			URI: `XRpc-unencrypted:`,
		},
		{
			URI: `XRpC-unencrypted:`,
		},
		{
			URI: `XRPc-unencrypted:`,
		},
		{
			URI: `XRPC-unencrypted:`,
		},



		{
			URI: `xrpc-unencrypted://`,
		},
		{
			URI: `xrpC-unencrypted://`,
		},
		{
			URI: `xrPc-unencrypted://`,
		},
		{
			URI: `xrPC-unencrypted://`,
		},
		{
			URI: `xRpc-unencrypted://`,
		},
		{
			URI: `xRpC-unencrypted://`,
		},
		{
			URI: `xRPc-unencrypted://`,
		},
		{
			URI: `xRPC-unencrypted://`,
		},
		{
			URI: `Xrpc-unencrypted://`,
		},
		{
			URI: `XrpC-unencrypted://`,
		},
		{
			URI: `XrPc-unencrypted://`,
		},
		{
			URI: `XrPC-unencrypted://`,
		},
		{
			URI: `XRpc-unencrypted://`,
		},
		{
			URI: `XRpC-unencrypted://`,
		},
		{
			URI: `XRPc-unencrypted://`,
		},
		{
			URI: `XRPC-unencrypted://`,
		},



		{
			URI: `xrpc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xrpC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xrPc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xrPC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xRpc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xRpC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xRPc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `xRPC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `Xrpc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XrpC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XrPc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XrPC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XRpc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XRpC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XRPc-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},
		{
			URI: `XRPC-unencrypted://Example.COM`,
			ExpectedHost:           "Example.COM",
		},



		{
			URI: `xrpc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xrpC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xrPc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xrPC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xRpc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xRpC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xRPc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `xRPC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `Xrpc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XrpC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XrPc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XrPC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XRpc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XRpC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XRPc-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},
		{
			URI: `XRPC-unencrypted://Example.COM/APP.Cherry.fooBar`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
		},



		{
			URI: `xrpc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xrpC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xrPc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xrPC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xRpc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xRpC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xRPc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `xRPC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `Xrpc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XrpC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XrPc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XrPC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XRpc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XRpC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XRPc-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},
		{
			URI: `XRPC-unencrypted://Example.COM/APP.Cherry.fooBar?actor=JoeBlow&sort=desc`,
			ExpectedHost:           "Example.COM",
			ExpectedCollection:                 "APP.Cherry.fooBar",
			ExpectedQuery:                                        "actor=JoeBlow&sort=desc",
		},



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



		{
			URI:            `xrpc-unencrypted://Host.EXAMPLE/apple.BANANA.Cherry.dAtE/wxyZ`,
			ExpectedHost:                      "Host.EXAMPLE",
			ExpectedCollection:                             "apple.BANANA.Cherry.dAtE",
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
			ExpectedError: `xrpcuri: URI "http://example.com" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},



		{
			URI: "x",
			ExpectedError: `xrpcuri: URI "x" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xr",
			ExpectedError: `xrpcuri: URI "xr" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrp",
			ExpectedError: `xrpcuri: URI "xrp" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc",
			ExpectedError: `xrpcuri: URI "xrpc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-",
			ExpectedError: `xrpcuri: URI "xrpc-" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-u",
			ExpectedError: `xrpcuri: URI "xrpc-u" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-un",
			ExpectedError: `xrpcuri: URI "xrpc-un" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-une",
			ExpectedError: `xrpcuri: URI "xrpc-une" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unen",
			ExpectedError: `xrpcuri: URI "xrpc-unen" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unenc",
			ExpectedError: `xrpcuri: URI "xrpc-unenc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencr",
			ExpectedError: `xrpcuri: URI "xrpc-unencr" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencry",
			ExpectedError: `xrpcuri: URI "xrpc-unencry" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencryp",
			ExpectedError: `xrpcuri: URI "xrpc-unencryp" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencrypt",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypt" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencrypte",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypte" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencrypted",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypted" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},



		{
			URI: ":",
			ExpectedError: `xrpcuri: URI ":" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "x:",
			ExpectedError: `xrpcuri: URI "x:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xr:",
			ExpectedError: `xrpcuri: URI "xr:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrp:",
			ExpectedError: `xrpcuri: URI "xrp:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc:",
			ExpectedError: `xrpcuri: URI "xrpc:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-:",
			ExpectedError: `xrpcuri: URI "xrpc-:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-u:",
			ExpectedError: `xrpcuri: URI "xrpc-u:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-un:",
			ExpectedError: `xrpcuri: URI "xrpc-un:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-une:",
			ExpectedError: `xrpcuri: URI "xrpc-une:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unen:",
			ExpectedError: `xrpcuri: URI "xrpc-unen:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unenc:",
			ExpectedError: `xrpcuri: URI "xrpc-unenc:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencr:",
			ExpectedError: `xrpcuri: URI "xrpc-unencr:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencry:",
			ExpectedError: `xrpcuri: URI "xrpc-unencry:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencryp:",
			ExpectedError: `xrpcuri: URI "xrpc-unencryp:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencrypt:",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypt:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpc-unencrypte:",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypte:" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
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
