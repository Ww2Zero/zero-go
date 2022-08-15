package _60_LFUCache

import (
	"testing"
)

func TestLFU(t *testing.T) {
	lfuCache := Constructor(3)
	lfuCache.Put(1, 1)
	lfuCache.Put(2, 2)
	lfuCache.Put(3, 3)
	lfuCache.Put(4, 4)
	v1 := lfuCache.Get(4)
	assertEqual(t, v1, 4)
	lfuCache.Put(4, 5)
	lfuCache.Put(4, 6)
	v2 := lfuCache.Get(2)
	assertEqual(t, v2, 2)
}

func assertEqual(t *testing.T, v1 int, i int) {
	if v1 != i {
		t.Errorf("assertEqual fail, %v %v", v1, i)
	}
}
