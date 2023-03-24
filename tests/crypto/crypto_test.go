package crypto

import (
	"github.com/0xIsRookie/Command/crypto"
	"testing"
)

func TestGetMd5(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"", "d41d8cd98f00b204e9800998ecf8427e"},    // 空字符串的 MD5 值
		{"foo", "acbd18db4cc2f85cedef654fccc4a4d8"}, // "foo" 的 MD5 值
		{"bar", "37b51d194a7513e45b56f6524f2d51f2"}, // "bar" 的 MD5 值
		{"baz", "73feffa4b7f6bb68e44cf984c85f6e88"}, // "baz" 的 MD5 值
	}

	for _, tc := range testCases {
		output := crypto.GetMd5(tc.input)
		if output != tc.expected {
			t.Errorf("GetMd5(%q) = %q, 希望返回 %q", tc.input, output, tc.expected)
		}
	}
}
