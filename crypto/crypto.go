package crypto

import (
	"crypto/md5"
	"encoding/hex"
)

// GetMd5 计算输入字符串的MD5哈希值，并以十六进制字符串表示返回。
//
// 参数：
//   - str：待计算的字符串。
//   - bit：表示哈希值的长度。可以是32或16。如果不是这两个值之一，返回空字符串。
//
// 返回值：
//   - string：输入字符串的MD5哈希值的十六进制字符串表示。
func GetMd5(str string, bit int) string {
	hash := md5.Sum([]byte(str))
	if bit == 32 {
		return hex.EncodeToString(hash[:])
	} else if bit == 16 {
		return hex.EncodeToString(hash[4:12])
	} else {
		return ""
	}
}
