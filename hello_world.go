package main

import (
	"fmt"
)

func CalcAbsDiff(x uint, y uint) uint {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}

func main() {
	fmt.Println("it works!")
	fmt.Println(CalcAbsDiff(10, 412))
}

