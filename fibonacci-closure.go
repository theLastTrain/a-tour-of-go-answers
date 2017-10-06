package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	f0, f1, fn := 0, 1, 1
	return func() int {
		fn = f0 + f1
		f0 = f1
		f1 = fn
		return f0
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
