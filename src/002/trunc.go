package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x1 string  // input as string
	var x2 float64 // input as valid float
	var x3 int     // float rounded to nearest int

	fmt.Println("Enter your floating point number:")

	count, err := fmt.Scanln(&x1)

	fmt.Println("Thank you for entering", count, "value.")

	if err != nil {
		panic(err)
	}

	x2, err = strconv.ParseFloat(x1, 64)

	if err != nil {
		panic(err)
	}

	x3 = int(x2)

	fmt.Println("Truncated version of your input is ", x3)
}
