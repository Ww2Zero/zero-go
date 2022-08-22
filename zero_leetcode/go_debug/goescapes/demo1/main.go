package main

type S struct{}

func main() {
	var x S
	_ = identity(x)
}

func identity(x S) S {
	return x
}

// Output:
// # command-line-arguments
// 无逃逸参数
