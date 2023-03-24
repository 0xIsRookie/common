package random

import (
	"math/rand"
	"time"
)

// GetRandomChild 在任意数组中,随机一个值并返回
func GetRandomChild[T any](arr []T) (T, error) {
	var result T
	if len(arr) == 0 {
		return result, ErrorNilList
	}

	rand.Seed(time.Now().UnixNano())

	randomIndex := rand.Intn(len(arr))
	result = arr[randomIndex]
	return result, nil
}

// GetRandomChildren 在任意数组中,随机指定的一个列表
func GetRandomChildren[T any](arr []T, length int) ([]T, error) {
	var result []T
	if len(arr) == 0 {
		return result, ErrorNilList
	}

	if len(arr) < length {
		return arr, ErrorArraySite
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(arr))
		result = append(result, arr[randomIndex])
	}
	return result, nil
}
