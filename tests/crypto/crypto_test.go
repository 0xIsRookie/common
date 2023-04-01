package crypto

import (
	"github.com/0xIsRookie/common/crypto"
	"testing"
)

func TestGetMd5(t *testing.T) {
	testCases := []struct {
		input string
		bit   int
		want  string
	}{
		{input: "hello world", bit: 16, want: "e01eeed093cb22bb"},
		{input: "hello world", bit: 32, want: "5eb63bbbe01eeed093cb22bb8f5acdc3"},
		{input: "", bit: 32, want: "d41d8cd98f00b204e9800998ecf8427e"},
		{input: "foo", bit: 32, want: "acbd18db4cc2f85cedef654fccc4a4d8"},
	}

	for _, tc := range testCases {
		got := crypto.GetMd5(tc.input, tc.bit)
		if got != tc.want {
			t.Errorf("getMd5(%q, %d) = %q; want %q", tc.input, tc.bit, got, tc.want)
		}
	}
}
