package main

type S struct{}

func main() {
	var x S
	_ = *ref(x)
}

func ref(z S) *S {
	return &z
}

// Output:
// # command-line-arguments
//./main.go:10:10: moved to heap: z
