package random

import (
	"math/rand"
	"time"
)

// GetRandomChild 从指定的任意类型数组中随机选取一个元素
//
// 参数：
// - arr: 待选取的任意类型数组
//
// 返回值：
// - T: 随机选取的任意类型数组中的一个元素
// - error: 如果待选取的数组为空，则返回 ErrorEmptyList 错误信息；如果选取过程中出现错误，则返回相应的错误信息
func GetRandomChild[T any](arr []T) (T, error) {
	var result T
	if len(arr) == 0 {
		return result, ErrorEmptyList
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(arr))
	result = arr[randomIndex]
	return result, nil
}

// GetRandomChildren 从给定的切片中随机选择多个元素并返回
//
// 参数：
// - arr: 给定的切片
// - length: 随机选择的元素个数
//
// 返回值：
// - []T: 随机选择的元素列表
// - error: 如果给定的切片为空，则返回 ErrorEmptyList 错误；如果给定的切片长度小于所需的随机元素个数，则返回 ErrorInvalidLength 错误
func GetRandomChildren[T any](arr []T, length int) ([]T, error) {
	var result []T
	if len(arr) == 0 {
		return result, ErrorEmptyList
	}

	if len(arr) < length {
		return arr, ErrorInvalidLength
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(arr))
		result = append(result, arr[randomIndex])
	}
	return result, nil
}
