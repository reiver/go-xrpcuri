package xrpcurienc_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri/enc"
)

func TestValidatePrefix(t *testing.T) {

	tests := []struct{
		URI string
	}{
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

		err := xrpcurienc.ValidatePrefix(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}

func TestValidatePrefix_fail(t *testing.T) {

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
			URI: "xrpc:",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xrpC:",
			ExpectedError: `xrpcuri: XRPC-URI "xrpC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xrPc:",
			ExpectedError: `xrpcuri: XRPC-URI "xrPc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xrPC:",
			ExpectedError: `xrpcuri: XRPC-URI "xrPC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xRpc:",
			ExpectedError: `xrpcuri: XRPC-URI "xRpc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xRpC:",
			ExpectedError: `xrpcuri: XRPC-URI "xRpC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xRPc:",
			ExpectedError: `xrpcuri: XRPC-URI "xRPc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "xRPC:",
			ExpectedError: `xrpcuri: XRPC-URI "xRPC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "Xrpc:",
			ExpectedError: `xrpcuri: XRPC-URI "Xrpc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XrpC:",
			ExpectedError: `xrpcuri: XRPC-URI "XrpC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XrPc:",
			ExpectedError: `xrpcuri: XRPC-URI "XrPc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XrPC:",
			ExpectedError: `xrpcuri: XRPC-URI "XrPC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XRpc:",
			ExpectedError: `xrpcuri: XRPC-URI "XRpc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XRpC:",
			ExpectedError: `xrpcuri: XRPC-URI "XRpC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XRPc:",
			ExpectedError: `xrpcuri: XRPC-URI "XRPc:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},
		{
			URI: "XRPC:",
			ExpectedError: `xrpcuri: XRPC-URI "XRPC:" is not valid because it does not have "//" after "xrpc:" — too short`,
		},



		{
			URI: "http://example.com",
			ExpectedError: `xrpcuri: URI "http://example.com" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
	}

	for testNumber, test := range tests {

		err := xrpcurienc.ValidatePrefix(test.URI)
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
