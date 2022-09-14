package main

import (
	"fmt"
	"unsafe"
)

type Person interface {
	growUp()
}

type Student struct {
	age int
}

func (p Student) growUp() {
	p.age += 1
	return
}

type iface struct {
	tab  *itab
	data unsafe.Pointer
}
type itab struct {
	inter uintptr
	_type uintptr
	link  uintptr
	hash  uint32
	_     [4]byte
	fun   [1]uintptr
}

func main() {
	var qcrao = Person(Student{age: 18})

	iface := (*iface)(unsafe.Pointer(&qcrao))
	fmt.Printf("iface=%p\n", iface)
	fmt.Printf("iface= %+v\n", iface)
	fmt.Printf("iface.tab.hash = %#x\n", iface.tab.hash)
	fmt.Println(iface)

	fmt.Println(iface.tab)
}
