package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestValidateScheme(t *testing.T) {

	tests := []struct{
		URI string
	}{
		{
			URI: "xrpc:",
		},
		{
			URI: "xrpC:",
		},
		{
			URI: "xrPc:",
		},
		{
			URI: "xrPC:",
		},
		{
			URI: "xRpc:",
		},
		{
			URI: "xRpC:",
		},
		{
			URI: "xRPc:",
		},
		{
			URI: "xRPC:",
		},
		{
			URI: "Xrpc:",
		},
		{
			URI: "XrpC:",
		},
		{
			URI: "XrPc:",
		},
		{
			URI: "XrPC:",
		},
		{
			URI: "XRpc:",
		},
		{
			URI: "XRpC:",
		},
		{
			URI: "XRPc:",
		},
		{
			URI: "XRPC:",
		},



		{
			URI: "xrpc://",
		},
		{
			URI: "xrpC://",
		},
		{
			URI: "xrPc://",
		},
		{
			URI: "xrPC://",
		},
		{
			URI: "xRpc://",
		},
		{
			URI: "xRpC://",
		},
		{
			URI: "xRPc://",
		},
		{
			URI: "xRPC://",
		},
		{
			URI: "Xrpc://",
		},
		{
			URI: "XrpC://",
		},
		{
			URI: "XrPc://",
		},
		{
			URI: "XrPC://",
		},
		{
			URI: "XRpc://",
		},
		{
			URI: "XRpC://",
		},
		{
			URI: "XRPc://",
		},
		{
			URI: "XRPC://",
		},



		{
			URI: "xrpc://VIDEO.Archive.ORG",
		},
		{
			URI: "xrpC://VIDEO.Archive.ORG",
		},
		{
			URI: "xrPc://VIDEO.Archive.ORG",
		},
		{
			URI: "xrPC://VIDEO.Archive.ORG",
		},
		{
			URI: "xRpc://VIDEO.Archive.ORG",
		},
		{
			URI: "xRpC://VIDEO.Archive.ORG",
		},
		{
			URI: "xRPc://VIDEO.Archive.ORG",
		},
		{
			URI: "xRPC://VIDEO.Archive.ORG",
		},
		{
			URI: "Xrpc://VIDEO.Archive.ORG",
		},
		{
			URI: "XrpC://VIDEO.Archive.ORG",
		},
		{
			URI: "XrPc://VIDEO.Archive.ORG",
		},
		{
			URI: "XrPC://VIDEO.Archive.ORG",
		},
		{
			URI: "XRpc://VIDEO.Archive.ORG",
		},
		{
			URI: "XRpC://VIDEO.Archive.ORG",
		},
		{
			URI: "XRPc://VIDEO.Archive.ORG",
		},
		{
			URI: "XRPC://VIDEO.Archive.ORG",
		},



		{
			URI: "xrpc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xrpC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xrPc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xrPC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRpc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRpC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRPc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRPC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "Xrpc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XrpC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XrPc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XrPC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRpc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRpC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRPc://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRPC://VIDEO.Archive.ORG/COM.Example.fooBar",
		},



		{
			URI: "xrpc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xrpC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xrPc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xrPC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRpc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRpC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRPc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRPC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "Xrpc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XrpC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XrPc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XrPC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRpc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRpC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRPc://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRPC://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},



		{
			URI: "XRPC://user:pass@SOMETHING.Example/banana/wxyz",
		},
	}

	for testNumber, test := range tests {

		err := xrpcuri.ValidateScheme(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}

func TestValidateScheme_fail(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: `xrpcuri: empty uri`,
		},



		{
			URI: "a",
			ExpectedError: `xrpcuri: URI "a" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "xrpc",
			ExpectedError: `xrpcuri: URI "xrpc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrpC",
			ExpectedError: `xrpcuri: URI "xrpC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrPc",
			ExpectedError: `xrpcuri: URI "xrPc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xrPC",
			ExpectedError: `xrpcuri: URI "xrPC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xRpc",
			ExpectedError: `xrpcuri: URI "xRpc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xRpC",
			ExpectedError: `xrpcuri: URI "xRpC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xRPc",
			ExpectedError: `xrpcuri: URI "xRPc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "xRPC",
			ExpectedError: `xrpcuri: URI "xRPC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "Xrpc",
			ExpectedError: `xrpcuri: URI "Xrpc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XrpC",
			ExpectedError: `xrpcuri: URI "XrpC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XrPc",
			ExpectedError: `xrpcuri: URI "XrPc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XrPC",
			ExpectedError: `xrpcuri: URI "XrPC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XRpc",
			ExpectedError: `xrpcuri: URI "XRpc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XRpC",
			ExpectedError: `xrpcuri: URI "XRpC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XRPc",
			ExpectedError: `xrpcuri: URI "XRPc" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			URI: "XRPC",
			ExpectedError: `xrpcuri: URI "XRPC" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "xrpc-unencrypted:",
			ExpectedError: `xrpcuri: URI "xrpc-unencrypted:" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			URI: "http://example.com",
			ExpectedError: `xrpcuri: URI "http://example.com" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
	}

	for testNumber, test := range tests {

		err := xrpcuri.ValidateScheme(test.URI)
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
			t.Logf("ACTUAL-ERROR:   %s", actual)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}
