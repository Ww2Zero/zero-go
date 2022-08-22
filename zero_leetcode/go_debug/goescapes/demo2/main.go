package main

type S struct{}

func main() {
	var x S
	y := &x
	_ = *identity(y)
}

func identity(z *S) *S {
	return z
}

//	Output:
// # command-line-arguments
//./main.go:11:15: leaking param: z to result ~r1 level=0
