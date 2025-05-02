package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestValidate(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: "xrpcuri: empty uri",
		},



                {
                        URI: "xrpc://foo.com/example/123",
                        ExpectedError: `xrpcuri: XRPC-URI "xrpc://foo.com/example/123" has an id "example/123" that is not a valid NSID: nsid: nsid ("example/123") should have at least 3 segments but actually has 1`,
                },



		{
			URI: "xrpc://",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://" has an empty 'authority'`,
		},
		{
			URI: "xrpc:///",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc:///" has an empty 'authority'`,
		},
		{
			URI: "xrpc://?",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://?" has an empty 'authority'`,
		},
		{
			URI: "xrpc://#",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://#" has an empty 'authority'`,
		},
		{
			URI: "xrpc://?#",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://?#" has an empty 'authority'`,
		},



		{
			URI: "xrpc://user:pass@foo.com",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://user:pass@foo.com" may not have an "@" in its authority "user:pass@foo.com"`,
		},
		{
			URI: "xrpc://user:pass@Foo.COM",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://user:pass@Foo.COM" may not have an "@" in its authority "user:pass@Foo.COM"`,
		},
		{
			URI: "xrpc://@",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://@" may not have an "@" in its authority "@"`,
		},
		{
			URI: "xrpc://@example",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://@example" may not have an "@" in its authority "@example"`,
		},
		{
			URI: "xrpc://@example.com",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://@example.com" may not have an "@" in its authority "@example.com"`,
		},
	}

	for testNumber, test := range tests {

		err := xrpcuri.Validate(test.URI)
		if nil == err {
			t.Errorf("For test #%d, expected an error but didn't get one.", testNumber)
			t.Logf("URI: %q", test.URI)
			continue
		}

		expected := test.ExpectedError
		actual := err.Error()

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("URI: %q", test.URI)
			continue
		}
	}
}
