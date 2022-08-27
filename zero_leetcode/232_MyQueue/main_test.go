package _32_MyQueue

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	q := Constructor()
	q.Push(1)
	q.Push(2)
	peek := q.Peek()
	fmt.Printf("peek:%d\n", peek)
	pop := q.Pop()
	fmt.Printf("pop:%d\n", pop)
	peek2 := q.Peek()
	fmt.Printf("peek2:%d\n", peek2)
}
