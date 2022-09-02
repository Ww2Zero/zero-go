package xxx_skiplist

import (
	"testing"
)

func TestSkiplist2(t *testing.T) {
	n := 20
	sl := New()
	for i := 0; i < n; i++ {
		sl.Insert(i, i)
	}
	sl.PrintAll()
}
