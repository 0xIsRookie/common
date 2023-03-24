package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5 计算输入字符串的 MD5 值，并返回十六进制字符串表示。
//
// 参数：
//   - v: 待计算的字符串。
//
// 返回值：
//   - string: 输入字符串的 MD5 值，以十六进制字符串表示。
func GetMd5(v string) string {
	d := []byte(v)
	h := md5.New()
	h.Write(d)
	return hex.EncodeToString(h.Sum(nil))
}
