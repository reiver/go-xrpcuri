package xrpcuri_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri"
)

func TestValidatePrefixUnencrypted(t *testing.T) {

	tests := []struct{
		URI string
	}{
		{
			URI: "xrpc-unencrypted://",
		},
		{
			URI: "xrpC-unencrypted://",
		},
		{
			URI: "xrPc-unencrypted://",
		},
		{
			URI: "xrPC-unencrypted://",
		},
		{
			URI: "xRpc-unencrypted://",
		},
		{
			URI: "xRpC-unencrypted://",
		},
		{
			URI: "xRPc-unencrypted://",
		},
		{
			URI: "xRPC-unencrypted://",
		},
		{
			URI: "Xrpc-unencrypted://",
		},
		{
			URI: "XrpC-unencrypted://",
		},
		{
			URI: "XrPc-unencrypted://",
		},
		{
			URI: "XrPC-unencrypted://",
		},
		{
			URI: "XRpc-unencrypted://",
		},
		{
			URI: "XRpC-unencrypted://",
		},
		{
			URI: "XRPc-unencrypted://",
		},
		{
			URI: "XRPC-unencrypted://",
		},

		{
			URI: "xrpc-UnEnCrYpTeD://",
		},
		{
			URI: "xrpC-UnEnCrYpTeD://",
		},
		{
			URI: "xrPc-UnEnCrYpTeD://",
		},
		{
			URI: "xrPC-UnEnCrYpTeD://",
		},
		{
			URI: "xRpc-UnEnCrYpTeD://",
		},
		{
			URI: "xRpC-UnEnCrYpTeD://",
		},
		{
			URI: "xRPc-UnEnCrYpTeD://",
		},
		{
			URI: "xRPC-UnEnCrYpTeD://",
		},
		{
			URI: "Xrpc-UnEnCrYpTeD://",
		},
		{
			URI: "XrpC-UnEnCrYpTeD://",
		},
		{
			URI: "XrPc-UnEnCrYpTeD://",
		},
		{
			URI: "XrPC-UnEnCrYpTeD://",
		},
		{
			URI: "XRpc-UnEnCrYpTeD://",
		},
		{
			URI: "XRpC-UnEnCrYpTeD://",
		},
		{
			URI: "XRPc-UnEnCrYpTeD://",
		},
		{
			URI: "XRPC-UnEnCrYpTeD://",
		},

		{
			URI: "xrpc-UNENCRYPTED://",
		},
		{
			URI: "xrpC-UNENCRYPTED://",
		},
		{
			URI: "xrPc-UNENCRYPTED://",
		},
		{
			URI: "xrPC-UNENCRYPTED://",
		},
		{
			URI: "xRpc-UNENCRYPTED://",
		},
		{
			URI: "xRpC-UNENCRYPTED://",
		},
		{
			URI: "xRPc-UNENCRYPTED://",
		},
		{
			URI: "xRPC-UNENCRYPTED://",
		},
		{
			URI: "Xrpc-UNENCRYPTED://",
		},
		{
			URI: "XrpC-UNENCRYPTED://",
		},
		{
			URI: "XrPc-UNENCRYPTED://",
		},
		{
			URI: "XrPC-UNENCRYPTED://",
		},
		{
			URI: "XRpc-UNENCRYPTED://",
		},
		{
			URI: "XRpC-UNENCRYPTED://",
		},
		{
			URI: "XRPc-UNENCRYPTED://",
		},
		{
			URI: "XRPC-UNENCRYPTED://",
		},



		{
			URI: "xrpc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xrpC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xrPc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xrPC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xRpc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xRpC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xRPc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "xRPC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "Xrpc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XrpC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XrPc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XrPC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XRpc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XRpC-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XRPc-unencrypted://VIDEO.Archive.ORG",
		},
		{
			URI: "XRPC-unencrypted://VIDEO.Archive.ORG",
		},



		{
			URI: "xrpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xrpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xrPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xrPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "xRPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "Xrpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XrpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XrPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XrPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},
		{
			URI: "XRPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar",
		},



		{
			URI: "xrpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xrpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xrPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xrPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "xRPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "Xrpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XrpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XrPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XrPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRpc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRpC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRPc-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},
		{
			URI: "XRPC-unencrypted://VIDEO.Archive.ORG/COM.Example.fooBar?actor=JoeBlow",
		},



		{
			URI: "XRPC-UNENCRYPTED://user:pass@SOMETHING.Example/banana/wxyz",
		},
	}

	for testNumber, test := range tests {

		err := xrpcuri.ValidatePrefixUnencrypted(test.URI)

		if nil != err {
			t.Errorf("For test #%d, did not expect an error but actually got one.", testNumber)
			t.Logf("ERROR: %s", err)
			t.Logf("URI: %s", test.URI)
			continue
		}
	}
}

func TestValidatePrefixUnencrypted_fail(t *testing.T) {

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
			ExpectedError: `xrpcuri: URI "a" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},



		{
			URI: "xrpc",
			ExpectedError: `xrpcuri: URI "xrpc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrpC",
			ExpectedError: `xrpcuri: URI "xrpC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrPc",
			ExpectedError: `xrpcuri: URI "xrPc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xrPC",
			ExpectedError: `xrpcuri: URI "xrPC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xRpc",
			ExpectedError: `xrpcuri: URI "xRpc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xRpC",
			ExpectedError: `xrpcuri: URI "xRpC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xRPc",
			ExpectedError: `xrpcuri: URI "xRPc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "xRPC",
			ExpectedError: `xrpcuri: URI "xRPC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "Xrpc",
			ExpectedError: `xrpcuri: URI "Xrpc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XrpC",
			ExpectedError: `xrpcuri: URI "XrpC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XrPc",
			ExpectedError: `xrpcuri: URI "XrPc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XrPC",
			ExpectedError: `xrpcuri: URI "XrPC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XRpc",
			ExpectedError: `xrpcuri: URI "XRpc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XRpC",
			ExpectedError: `xrpcuri: URI "XRpC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XRPc",
			ExpectedError: `xrpcuri: URI "XRPc" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
		{
			URI: "XRPC",
			ExpectedError: `xrpcuri: URI "XRPC" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},



		{
			URI: "xrpc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrpC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrPc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrPc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrPC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrPC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRpc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRpc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRpC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRpC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRPc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRPc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRPC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRPC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "Xrpc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "Xrpc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrpC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrpC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrPc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrPc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrPC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrPC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRpc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRpc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRpC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRpC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRPc-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRPc-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRPC-unencrypted:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRPC-unencrypted:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},

		{
			URI: "xrpc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrpC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrPc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrPc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrPC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrPC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRpc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRpc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRpC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRpC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRPc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRPc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRPC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRPC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "Xrpc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "Xrpc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrpC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrpC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrPc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrPc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrPC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrPC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRpc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRpc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRpC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRpC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRPc-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRPc-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRPC-UnEnCrYpTeD:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRPC-UnEnCrYpTeD:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},

		{
			URI: "xrpc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrpC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrpC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrPc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrPc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xrPC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xrPC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRpc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRpc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRpC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRpC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRPc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRPc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "xRPC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "xRPC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "Xrpc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "Xrpc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrpC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrpC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrPc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrPc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XrPC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XrPC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRpc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRpc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRpC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRpC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRPc-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRPc-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},
		{
			URI: "XRPC-UNENCRYPTED:",
			ExpectedError: `xrpcuri: XRPC-unencrypted-URI "XRPC-UNENCRYPTED:" is not valid because it does not have "//" after "xrpc-unencrypted:" — too short`,
		},



		{
			URI: "http://example.com",
			ExpectedError: `xrpcuri: URI "http://example.com" is not an XRPC-unencrypted-URI because it does not begin with "xrpc-unencrypted:"`,
		},
	}

	for testNumber, test := range tests {

		err := xrpcuri.ValidatePrefixUnencrypted(test.URI)
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
