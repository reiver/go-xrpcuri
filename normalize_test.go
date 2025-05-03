package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestNormalize(t *testing.T) {

	tests := []struct{
		URI string
		Expected string
	}{
		{
			URI:      "",
			Expected: "",
		},



		{
			URI:      "http://example.com?wow#",
			Expected: "http://example.com?wow#",
		},
		{
			URI:      "HTTP://Example.COM?wow#",
			Expected: "HTTP://Example.COM?wow#",
		},



		{
			URI:      "x",
			Expected: "x",
		},
		{
			URI:      "X",
			Expected: "X",
		},



		{
			URI:      "xr",
			Expected: "xr",
		},
		{
			URI:      "xR",
			Expected: "xR",
		},
		{
			URI:      "Xr",
			Expected: "Xr",
		},
		{
			URI:      "XR",
			Expected: "XR",
		},



		{
			URI:      "xrp",
			Expected: "xrp",
		},
		{
			URI:      "xrP",
			Expected: "xrP",
		},
		{
			URI:      "xRp",
			Expected: "xRp",
		},
		{
			URI:      "xRP",
			Expected: "xRP",
		},
		{
			URI:      "Xrp",
			Expected: "Xrp",
		},
		{
			URI:      "XrP",
			Expected: "XrP",
		},
		{
			URI:      "XRp",
			Expected: "XRp",
		},
		{
			URI:      "XRP",
			Expected: "XRP",
		},



		{
			URI:      "xrpc",
			Expected: "xrpc",
		},
		{
			URI:      "xrpC",
			Expected: "xrpC",
		},
		{
			URI:      "xrPc",
			Expected: "xrPc",
		},
		{
			URI:      "xrPC",
			Expected: "xrPC",
		},
		{
			URI:      "xRpc",
			Expected: "xRpc",
		},
		{
			URI:      "xRpC",
			Expected: "xRpC",
		},
		{
			URI:      "xRPc",
			Expected: "xRPc",
		},
		{
			URI:      "xRPC",
			Expected: "xRPC",
		},
		{
			URI:      "Xrpc",
			Expected: "Xrpc",
		},
		{
			URI:      "XrpC",
			Expected: "XrpC",
		},
		{
			URI:      "XrPc",
			Expected: "XrPc",
		},
		{
			URI:      "XrPC",
			Expected: "XrPC",
		},
		{
			URI:      "XRpc",
			Expected: "XRpc",
		},
		{
			URI:      "XRpC",
			Expected: "XRpC",
		},
		{
			URI:      "XRPc",
			Expected: "XRPc",
		},
		{
			URI:      "XRPC",
			Expected: "XRPC",
		},



		{
			URI:      "xrpc:",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpC:",
			Expected: "xrpc://",
		},
		{
			URI:      "xrPc:",
			Expected: "xrpc://",
		},
		{
			URI:      "xrPC:",
			Expected: "xrpc://",
		},
		{
			URI:      "xRpc:",
			Expected: "xrpc://",
		},
		{
			URI:      "xRpC:",
			Expected: "xrpc://",
		},
		{
			URI:      "xRPc:",
			Expected: "xrpc://",
		},
		{
			URI:      "xRPC:",
			Expected: "xrpc://",
		},
		{
			URI:      "Xrpc:",
			Expected: "xrpc://",
		},
		{
			URI:      "XrpC:",
			Expected: "xrpc://",
		},
		{
			URI:      "XrPc:",
			Expected: "xrpc://",
		},
		{
			URI:      "XrPC:",
			Expected: "xrpc://",
		},
		{
			URI:      "XRpc:",
			Expected: "xrpc://",
		},
		{
			URI:      "XRpC:",
			Expected: "xrpc://",
		},
		{
			URI:      "XRPc:",
			Expected: "xrpc://",
		},
		{
			URI:      "XRPC:",
			Expected: "xrpc://",
		},



		{
			URI:      "xrpc://",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpC://",
			Expected: "xrpc://",
		},
		{
			URI:      "xrPc://",
			Expected: "xrpc://",
		},
		{
			URI:      "xrPC://",
			Expected: "xrpc://",
		},
		{
			URI:      "xRpc://",
			Expected: "xrpc://",
		},
		{
			URI:      "xRpC://",
			Expected: "xrpc://",
		},
		{
			URI:      "xRPc://",
			Expected: "xrpc://",
		},
		{
			URI:      "xRPC://",
			Expected: "xrpc://",
		},
		{
			URI:      "Xrpc://",
			Expected: "xrpc://",
		},
		{
			URI:      "XrpC://",
			Expected: "xrpc://",
		},
		{
			URI:      "XrPc://",
			Expected: "xrpc://",
		},
		{
			URI:      "XrPC://",
			Expected: "xrpc://",
		},
		{
			URI:      "XRpc://",
			Expected: "xrpc://",
		},
		{
			URI:      "XRpC://",
			Expected: "xrpc://",
		},
		{
			URI:      "XRPc://",
			Expected: "xrpc://",
		},
		{
			URI:      "XRPC://",
			Expected: "xrpc://",
		},



		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:///",
			Expected: "xrpc://",
		},



		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},
		{
			URI:      "xrpc:////",
			Expected: "xrpc://",
		},



		{
			URI:      "xrpc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrpC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "Xrpc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrpC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpC://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPc://example.com",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPC://example.com",
			Expected: "xrpc://example.com",
		},



		{
			URI:      "xrpc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrpC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "Xrpc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrpC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpC://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPc://Example.COM",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPC://Example.COM",
			Expected: "xrpc://example.com",
		},



		{
			URI:      "xrpc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrpC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "Xrpc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrpC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpC://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPc://Example.COM/",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPC://Example.COM/",
			Expected: "xrpc://example.com",
		},



		{
			URI:      "xrpc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrpC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "Xrpc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrpC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpC://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPc://Example.COM?",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPC://Example.COM?",
			Expected: "xrpc://example.com",
		},



		{
			URI:      "xrpc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrpC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xrPC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRpC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "xRPC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "Xrpc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrpC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XrPC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRpC://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPc://Example.COM#",
			Expected: "xrpc://example.com",
		},
		{
			URI:      "XRPC://Example.COM#",
			Expected: "xrpc://example.com",
		},



		{
			URI:      "xrpc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xrpC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "Xrpc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XrpC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPc://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPC://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc://test.example/com.example.fooBar",
		},



		{
			URI:      "xrpc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xrpC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "Xrpc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XrpC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPc://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPC://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc://test.example/com.example.fooBar",
		},
	}

	for testNumber, test := range tests {

		actual := xrpcuri.Normalize(test.URI)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual normalized XRPC-URI is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("XRPC-URI: %s", test.URI)
			continue
		}
	}
}
