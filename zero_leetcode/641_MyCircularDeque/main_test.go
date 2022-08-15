package _41_MyCircularDeque

import (
	"runtime"
	"testing"
)

func TestConstructor(t *testing.T) {
	circularDeque := Constructor(3)   // 设置容量大小为3
	b1 := circularDeque.InsertLast(1) // 返回 true
	assertBool(t, b1, true)
	b2 := circularDeque.InsertLast(2) // 返回 true
	assertBool(t, b2, true)
	b3 := circularDeque.InsertFront(3) // 返回 true
	assertBool(t, b3, true)
	b4 := circularDeque.InsertFront(4) // 已经满了，返回 false
	assertBool(t, b4, false)
	i1 := circularDeque.GetRear() // 返回 2
	assertInt(t, i1, 2)
	b5 := circularDeque.IsFull() // 返回 true
	assertBool(t, b5, true)
	b6 := circularDeque.DeleteLast() // 返回 true
	assertBool(t, b6, true)
	b7 := circularDeque.InsertFront(4) // 返回 true
	assertBool(t, b7, true)
	i2 := circularDeque.GetFront() // 返回 4
	assertInt(t, i2, 4)

}
func assertBool(t *testing.T, i, j bool) {
	if i != j {
		_, _, lineNo1, _ := runtime.Caller(1)
		t.Errorf("assertBool fail,lineNo=%v %v != %v", lineNo1, i, j)
	}
}
func assertInt(t *testing.T, i, j int) {
	if i != j {
		_, _, lineNo1, _ := runtime.Caller(1)
		t.Errorf("assertInt fail,lineNo=%v %v != %v", lineNo1, i, j)
	}
}
