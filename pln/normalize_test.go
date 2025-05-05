package xrpcuripln_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri/pln"
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
			URI:      "xrpc-unencrypted",
			Expected: "xrpc-unencrypted",
		},
		{
			URI:      "xrpC-unencrypted",
			Expected: "xrpC-unencrypted",
		},
		{
			URI:      "xrPc-unencrypted",
			Expected: "xrPc-unencrypted",
		},
		{
			URI:      "xrPC-unencrypted",
			Expected: "xrPC-unencrypted",
		},
		{
			URI:      "xRpc-unencrypted",
			Expected: "xRpc-unencrypted",
		},
		{
			URI:      "xRpC-unencrypted",
			Expected: "xRpC-unencrypted",
		},
		{
			URI:      "xRPc-unencrypted",
			Expected: "xRPc-unencrypted",
		},
		{
			URI:      "xRPC-unencrypted",
			Expected: "xRPC-unencrypted",
		},
		{
			URI:      "Xrpc-unencrypted",
			Expected: "Xrpc-unencrypted",
		},
		{
			URI:      "XrpC-unencrypted",
			Expected: "XrpC-unencrypted",
		},
		{
			URI:      "XrPc-unencrypted",
			Expected: "XrPc-unencrypted",
		},
		{
			URI:      "XrPC-unencrypted",
			Expected: "XrPC-unencrypted",
		},
		{
			URI:      "XRpc-unencrypted",
			Expected: "XRpc-unencrypted",
		},
		{
			URI:      "XRpC-unencrypted",
			Expected: "XRpC-unencrypted",
		},
		{
			URI:      "XRPc-unencrypted",
			Expected: "XRPc-unencrypted",
		},
		{
			URI:      "XRPC-unencrypted",
			Expected: "XRPC-unencrypted",
		},



		{
			URI:      "xrpc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrpC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "Xrpc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrpC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPc-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPC-unencrypted:",
			Expected: "xrpc-unencrypted://",
		},



		{
			URI:      "xrpc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrpC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "Xrpc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrpC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPc-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPC-unencrypted://",
			Expected: "xrpc-unencrypted://",
		},



		{
			URI:      "xrpc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrpC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "Xrpc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrpC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPc-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPC-unencrypted:///",
			Expected: "xrpc-unencrypted://",
		},



		{
			URI:      "xrpc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrpC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xrPC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRpC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "xRPC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "Xrpc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrpC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XrPC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRpC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPc-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},
		{
			URI:      "XRPC-unencrypted:////",
			Expected: "xrpc-unencrypted://",
		},



		{
			URI:      "xrpc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrpC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "Xrpc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrpC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPc-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPC-unencrypted://example.com",
			Expected: "xrpc-unencrypted://example.com",
		},



		{
			URI:      "xrpc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrpC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "Xrpc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrpC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPc-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPC-unencrypted://Example.COM",
			Expected: "xrpc-unencrypted://example.com",
		},



		{
			URI:      "xrpc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrpC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "Xrpc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrpC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPc-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPC-unencrypted://Example.COM/",
			Expected: "xrpc-unencrypted://example.com",
		},



		{
			URI:      "xrpc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrpC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "Xrpc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrpC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPc-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPC-unencrypted://Example.COM?",
			Expected: "xrpc-unencrypted://example.com",
		},



		{
			URI:      "xrpc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrpC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xrPC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRpC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "xRPC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "Xrpc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrpC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XrPC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRpC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPc-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},
		{
			URI:      "XRPC-unencrypted://Example.COM#",
			Expected: "xrpc-unencrypted://example.com",
		},



		{
			URI:      "xrpc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xrpC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "Xrpc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XrpC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPc-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPC-unencrypted://TEST.Example/COM.Example.fooBar",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},



		{
			URI:      "xrpc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xrpC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xrPC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRpC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "xRPC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "Xrpc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XrpC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XrPC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRpC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPc-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
		{
			URI:      "XRPC-unencrypted://TEST.Example/COM.Example.fooBar/wXYZ",
			Expected: "xrpc-unencrypted://test.example/com.example.fooBar",
		},
	}

	for testNumber, test := range tests {

		actual := xrpcuripln.Normalize(test.URI)

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
