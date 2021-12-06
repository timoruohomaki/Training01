package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x1 string
	var x2 float64
	var x3 int

	fmt.Printf("Enter your floating point number:")

	num, err := fmt.Scanln(&x1)

	x2, err = strconv.ParseFloat(x1, 64)

	x3 = int(x2)

	fmt.Printf("Truncated version of your input is ", x3)
}
