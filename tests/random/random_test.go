package random

import (
	"github.com/0xIsRookie/Command/random"
	"reflect"
	"testing"
)

type Name struct{ Name string }

func TestGetRandomChild(t *testing.T) {
	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	i, err := random.GetRandomChild(arrInt)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(i) != reflect.TypeOf(0) {
		t.Errorf("GetRandomChild Int数组测试异常")
	}

	arrStr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	s, err := random.GetRandomChild(arrStr)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(s) != reflect.TypeOf("0") {
		t.Errorf("GetRandomChild 字符串数组测试异常")
	}

	arrStruct := []Name{
		{"1"}, {"2"}, {"3"},
	}
	_struct, err := random.GetRandomChild(arrStruct)
	if err != nil {
		t.Fatal(err)
	}
	if reflect.TypeOf(_struct) != reflect.TypeOf(Name{}) {
		t.Errorf("GetRandomChild 结构体测试异常")
	}

	var arrEmpty []int
	_, err = random.GetRandomChild(arrEmpty)
	if err != random.ErrorNilList {
		t.Errorf("GetRandomChild 空数组测试异常")
	}
}

func TestGetRandomChildren(t *testing.T) {
	arrInt := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0}
	i, err := random.GetRandomChildren(arrInt, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(i) != 2 || reflect.TypeOf(i) != reflect.TypeOf(arrInt) {
		t.Errorf("GetRandomChild Int数组测试异常 %v, %v != %v", i, reflect.TypeOf(i), reflect.TypeOf(arrInt))
	}

	arrStr := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	s, err := random.GetRandomChildren(arrStr, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(s) != 2 || reflect.TypeOf(s) != reflect.TypeOf(arrStr) {
		t.Errorf("GetRandomChild 字符串数组测试异常 %v, %v != %v", s, reflect.TypeOf(i), reflect.TypeOf(s))
	}

	arrStruct := []Name{{"1"}, {"2"}, {"3"}}
	_struct, err := random.GetRandomChildren(arrStruct, 2)
	if err != nil {
		t.Fatal(err)
	}
	if len(_struct) != 2 || reflect.TypeOf(_struct) != reflect.TypeOf(arrStruct) {
		t.Errorf("GetRandomChild 结构体测试异常 %v, %v != %v", _struct, reflect.TypeOf(_struct), reflect.TypeOf(arrStruct))
	}

	var arrEmpty []int
	_, err = random.GetRandomChildren(arrEmpty, 2)
	if err != random.ErrorNilList {
		t.Errorf("GetRandomChild 空数组测试异常")
	}

	arrSite := []string{"1"}
	_, err = random.GetRandomChildren(arrSite, 2)
	if err != random.ErrorArraySite {
		t.Errorf("GetRandomChild 空数组测试异常")
	}
}
