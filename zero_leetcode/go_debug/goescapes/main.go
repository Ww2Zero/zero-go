package main

import "fmt"

func foo() *int {
	t := 3
	return &t
}

//build:+ go build -gcflags "-m -l"
func main() {
	x := foo()
	fmt.Println(*x)
}

// go build -gcflags "-m -l" main.go                                                                            ─╯
//# command-line-arguments
//./main.go:6:2: moved to heap: t
//./main.go:13:13: ... argument does not escape
//./main.go:13:14: *x escapes to heap

//go tool compile -S main.go >main.s.txt

//0x001c 00028 (main.go:5)	FUNCDATA	ZR, gclocals·2a5305abe05176240e61b8620e19a815(SB)
//0x001c 00028 (main.go:5)	FUNCDATA	$1, gclocals·33cdeccccebe80329f1fdbee7f5874cb(SB)
//0x001c 00028 (main.go:6)	MOVD	$type.int(SB), R0
//0x0024 00036 (main.go:6)	MOVD	R0, 8(RSP)
//0x0028 00040 (main.go:6)	PCDATA	$1, ZR
//0x0028 00040 (main.go:6)	CALL	runtime.newobject(SB)      <----- moved to heap
//0x002c 00044 (main.go:6)	MOVD	16(RSP), R0
//0x0030 00048 (main.go:6)	MOVD	$3, R1
//0x0034 00052 (main.go:6)	MOVD	R1, (R0)
//0x0038 00056 (main.go:7)	MOVD	R0, "".~r0(FP)
//0x003c 00060 (main.go:7)	MOVD	-8(RSP), R29
//0x0040 00064 (main.go:7)	MOVD.P	48(RSP), R30
