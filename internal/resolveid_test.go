package xrpcuri_internal_test

import (
	"testing"

	"github.com/reiver/go-xrpcuri/internal"
)

func TestResolveID(t *testing.T) {

	tests := []struct{
		ID string
		Expected string
	}{
		{
		},



		{
			ID:             "com.example.fooBar",
			Expected: "/xrpc/com.example.fooBar",
		},
		{
			ID:             "COM.Example.fooBar",
			Expected: "/xrpc/com.example.fooBar",
		},
	}

	for testNumber, test := range tests {

		actual := xrpcuri_internal.ResolveID(test.ID)

		expected := test.Expected

		if expected != actual {
			t.Errorf("For test #%d, the actual HTTP-path is not what was expected.", testNumber)
			t.Logf("EXPECTED: %s", expected)
			t.Logf("ACTUAL:   %s", actual)
			t.Logf("ID:       %s", test.ID)
			continue
		}
	}
}
