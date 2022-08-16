package _656_OrderedStream

import (
	"reflect"
	"runtime"
	"testing"
)

func TestConstructor(t *testing.T) {
	os := Constructor(5)
	i1 := os.Insert(3, "ccccc") // 插入 (3, "ccccc")，返回 []
	assertIntArray(t, i1, []string{})
	i2 := os.Insert(1, "aaaaa") // 插入 (1, "aaaaa")，返回 ["aaaaa"]
	assertIntArray(t, i2, []string{"aaaaa"})
	i3 := os.Insert(2, "bbbbb") // 插入 (2, "bbbbb")，返回 ["bbbbb", "ccccc"]
	assertIntArray(t, i3, []string{"bbbbb", "ccccc"})
	i4 := os.Insert(5, "eeeee") // 插入 (5, "eeeee")，返回 []
	assertIntArray(t, i4, []string{})
	i5 := os.Insert(4, "ddddd") // 插入 (4, "ddddd")，返回 ["ddddd", "eeeee"]
	assertIntArray(t, i5, []string{"ddddd", "eeeee"})
}

func assertIntArray(t *testing.T, n1, n2 []string) {
	if !reflect.DeepEqual(n1, n2) {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("assertIntArray fail, [%v] %v != %v", line, n1, n2)
	}
}
