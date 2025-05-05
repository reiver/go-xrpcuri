package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestResolve(t *testing.T) {

	tests := []struct{
		RequestType string
		URI string
		Expected string
	}{
		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "http://example.com",
			Expected: "http://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "http://example.com",
			Expected: "http://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "http://example.com",
			Expected: "http://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "HTTP://example.com",
			Expected: "HTTP://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "HTTP://example.com",
			Expected: "HTTP://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "HTTP://example.com",
			Expected: "HTTP://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "https://example.com",
			Expected: "https://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "https://example.com",
			Expected: "https://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "https://example.com",
			Expected: "https://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "HTTPS://example.com",
			Expected: "HTTPS://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "HTTPS://example.com",
			Expected: "HTTPS://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "HTTPS://example.com",
			Expected: "HTTPS://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "ws://example.com",
			Expected: "ws://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "ws://example.com",
			Expected: "ws://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "ws://example.com",
			Expected: "ws://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "WS://example.com",
			Expected: "WS://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "WS://example.com",
			Expected: "WS://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "WS://example.com",
			Expected: "WS://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "wss://example.com",
			Expected: "wss://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "wss://example.com",
			Expected: "wss://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "wss://example.com",
			Expected: "wss://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:      "WSS://example.com",
			Expected: "WSS://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:      "WSS://example.com",
			Expected: "WSS://example.com",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:      "WSS://example.com",
			Expected: "WSS://example.com",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:       "xrpc://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected: "https://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:       "xrpc://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected: "https://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:       "xrpc://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:   "wss://host.example/xrpc/com.atproto.repo.listRecords",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI:       "XRPC://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected: "https://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI:       "XRPC://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected: "https://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI:       "XRPC://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:   "wss://host.example/xrpc/com.atproto.repo.listRecords",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI: "xrpc-unencrypted://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:        "http://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI: "xrpc-unencrypted://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:        "http://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI: "xrpc-unencrypted://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:          "ws://host.example/xrpc/com.atproto.repo.listRecords",
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI: "XRPC-UNENCRYPTED://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:        "http://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI: "XRPC-UNENCRYPTED://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:        "http://host.example/xrpc/com.atproto.repo.listRecords",
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI: "XRPC-UNENCRYPTED://Host.EXAMPLE/com.atproto.repo.listRecords",
			Expected:          "ws://host.example/xrpc/com.atproto.repo.listRecords",
		},
	}

	for testNumber, test := range tests {

		actual, err := xrpcuri.Resolve(test.URI, test.RequestType)
		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			t.Logf("REQUEST-TYPE: %s", test.RequestType)
			continue
		}

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual resolved-URI is not what was expected.", testNumber)
			t.Logf("EXPECTED: %q", expected)
			t.Logf("ACTUAL:   %q", actual)
			t.Logf("URI:      %q", test.URI)
			t.Logf("REQUEST-TYPE: %s", test.RequestType)
			continue
		}
	}
}

func TestResolve_fail(t *testing.T) {

	tests := []struct{
		RequestType string
		URI string
		ExpectedError string
	}{
		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI: "banana://example.com",
			ExpectedError: `xrpcuri: URI "banana://example.com" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI: "banana://example.com",
			ExpectedError: `xrpcuri: URI "banana://example.com" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI: "banana://example.com",
			ExpectedError: `xrpcuri: URI "banana://example.com" is not an XRPC-URI because it does not begin with "xrpc:"`,
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI: "xrpc://example.com",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://example.com" has an id "" that is not a valid NSID: nsid: empty NSID`,
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI: "xrpc://example.com",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://example.com" has an id "" that is not a valid NSID: nsid: empty NSID`,
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI: "xrpc://example.com",
			ExpectedError: `xrpcuri: XRPC-URI "xrpc://example.com" has an id "" that is not a valid NSID: nsid: empty NSID`,
		},



		{
			RequestType: xrpcuri.RequestTypeExecute,
			URI: "xrpc-unencrypted://example.com",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://example.com" has an id "" that is not a valid NSID: nsid: empty NSID`,
		},
		{
			RequestType: xrpcuri.RequestTypeQuery,
			URI: "xrpc-unencrypted://example.com",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://example.com" has an id "" that is not a valid NSID: nsid: empty NSID`,
		},
		{
			RequestType: xrpcuri.RequestTypeSubscribe,
			URI: "xrpc-unencrypted://example.com",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted://example.com" has an id "" that is not a valid NSID: nsid: empty NSID`,
		},
	}

	for testNumber, test := range tests {

		_, err := xrpcuri.Resolve(test.URI, test.RequestType)
		if nil == err {
			t.Errorf("For test #%d, expected an error but did not actually get one.", testNumber)
			t.Logf("URI: %s", test.URI)
			t.Logf("REQUEST-TYPE: %s", test.RequestType)
			continue
		}

		actual := err.Error()
		expected := test.ExpectedError

		if expected != actual {
			t.Errorf("For test #%d, the actual error is not what was expected.", testNumber)
			t.Logf("EXPECTED-ERROR: %s", expected)
			t.Logf("ACTUAL-ERROR:   %s", actual)
			t.Logf("URI:      %q", test.URI)
			t.Logf("REQUEST-TYPE: %s", test.RequestType)
			continue
		}
	}
}
