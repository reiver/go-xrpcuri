package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestValidateUnencrypted(t *testing.T) {

	tests := []struct{
		URI string
		ExpectedError string
	}{
		{
			URI: "",
			ExpectedError: "xrpcuri: empty uri",
		},



                {
                        URI: "xrpc-unencrypted://foo.com/example/123",
                        ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://foo.com/example/123" has an id "example" that is not a valid NSID: nsid: nsid ("example") should have at least 3 segments but actually has 1`,
                },



		{
			URI: "xrpc-unencrypted://",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://" has an empty 'authority'`,
		},
		{
			URI: "xrpc-unencrypted:///",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted:///" has an empty 'authority'`,
		},
		{
			URI: "xrpc-unencrypted://?",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://?" has an empty 'authority'`,
		},
		{
			URI: "xrpc-unencrypted://#",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://#" has an empty 'authority'`,
		},
		{
			URI: "xrpc-unencrypted://?#",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://?#" has an empty 'authority'`,
		},



		{
			URI: "xrpc-unencrypted://user:pass@foo.com",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://user:pass@foo.com" may not have an "@" in its authority "user:pass@foo.com"`,
		},
		{
			URI: "xrpc-unencrypted://user:pass@Foo.COM",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://user:pass@Foo.COM" may not have an "@" in its authority "user:pass@Foo.COM"`,
		},
		{
			URI: "xrpc-unencrypted://@",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://@" may not have an "@" in its authority "@"`,
		},
		{
			URI: "xrpc-unencrypted://@example",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://@example" may not have an "@" in its authority "@example"`,
		},
		{
			URI: "xrpc-unencrypted://@example.com",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://@example.com" may not have an "@" in its authority "@example.com"`,
		},
	}

	for testNumber, test := range tests {

		err := xrpcuri.ValidateUnencrypted(test.URI)
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
