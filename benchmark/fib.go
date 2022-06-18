package main

import (
	"fmt"
	"os"
	"strconv"
)

func Fib(n int64) int64 {
	if n < 2 {
		return n
	}

	return Fib(n-1) + Fib(n-2)
}

func main() {
	intv, err := strconv.ParseInt(os.Args[1], 10, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(Fib(intv))
}
