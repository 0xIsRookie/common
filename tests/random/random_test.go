package random

import (
	"errors"
	"github.com/0xIsRookie/common/random"
	"testing"
)

func TestGetRandomChild(t *testing.T) {
	t.Run("空列表", func(t *testing.T) {
		var emptyList []int
		_, err := random.GetRandomChild(emptyList)
		if !errors.Is(err, random.ErrorEmptyList) {
			t.Errorf("GetRandomChild(%v) = _, %v; 希望返回 _, %v", emptyList, err, random.ErrorEmptyList)
		}
	})

	t.Run("只有一个元素的列表", func(t *testing.T) {
		list := []string{"foo"}
		element, err := random.GetRandomChild(list)
		if err != nil {
			t.Errorf("GetRandomChild(%v) = _, %v; 希望返回列表中唯一的元素", list, err)
		}
		if element != "foo" {
			t.Errorf("GetRandomChild(%v) = %v, _; 希望返回 %v", list, element, "foo")
		}
	})

	t.Run("多个元素的列表", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		element, err := random.GetRandomChild(list)
		if err != nil {
			t.Errorf("GetRandomChild(%v) = _, %v; 希望返回列表中的一个随机元素", list, err)
		}
		if !contains(list, element) {
			t.Errorf("GetRandomChild(%v) = %v, _; 希望返回列表中的一个随机元素", list, element)
		}
	})
}

func TestGetRandomChildren(t *testing.T) {
	t.Run("空列表", func(t *testing.T) {
		var emptyList []int
		_, err := random.GetRandomChildren(emptyList, 1)
		if !errors.Is(err, random.ErrorEmptyList) {
			t.Errorf("GetRandomChildren(%v, 1) = _, %v; 希望返回 _, %v", emptyList, err, random.ErrorEmptyList)
		}
	})

	t.Run("请求长度大于列表长度", func(t *testing.T) {
		list := []string{"foo", "bar"}
		_, err := random.GetRandomChildren(list, 3)
		if !errors.Is(err, random.ErrorInvalidLength) {
			t.Errorf("GetRandomChildren(%v, 3) = _, %v; 希望返回 _, %v", list, err, random.ErrorInvalidLength)
		}
	})

	t.Run("请求长度等于列表长度", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5}
		result, err := random.GetRandomChildren(list, len(list))
		if err != nil {
			t.Errorf("GetRandomChildren(%v, %v) = _, %v; 希望返回列表中的所有元素", list, len(list), err)
		}
		if len(list) != len(result) {
			t.Errorf("GetRandomChildren(%v, %v) = %v, _; 希望返回列表中的所有元素", list, len(list), result)
		}
	})

	t.Run("请求长度小于列表长度", func(t *testing.T) {
		list := []string{"foo", "bar", "baz", "qux", "quux"}
		result, err := random.GetRandomChildren(list, 3)
		if err != nil {
			t.Errorf("GetRandomChildren(%v, %v) = _, %v; 希望返回长度为 %v 的随机子列表", list, 3, err, 3)
		}
		if len(result) != 3 {
			t.Errorf("GetRandomChildren(%v, %v) = %v, _; 希望返回长度为 %v 的随机子列表", list, 3, result, 3)
		}
		for _, e := range result {
			if !contains(list, e) {
				t.Errorf("GetRandomChildren(%v, %v) 返回了不存在于列表中的元素 %v", list, 3, e)
			}
		}
	})
}
func contains[T comparable](list []T, element T) bool {
	for _, e := range list {
		if e == element {
			return true
		}
	}
	return false
}
