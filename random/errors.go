package random

import "errors"

var (
	ErrorNilList   = errors.New("数组不能为空")
	ErrorArraySite = errors.New("数组元素大小错误")
)
