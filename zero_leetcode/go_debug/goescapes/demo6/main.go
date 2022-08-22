package main

type S struct {
	M *int
}

// go build -gcflags "-m -l"
func main() {
	var x S
	var i int
	ref(&i, &x)
}

func ref(y *int, z *S) {
	z.M = y
}

// Output:
//# command-line-arguments
//./main.go:14:10: leaking param: y
//./main.go:14:18: z does not escape
//./main.go:10:6: moved to heap: i
