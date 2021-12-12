package main

import (
	"fmt"
	"strings"
)

func main() {

	var x1 string // source string

	fmt.Println("Enter your string:")

	count, err := fmt.Scanln(&x1)

	fmt.Println(" Thank you for entering ", count, "value.")

	if err != nil {
		panic(err)
	}

	res1 := strings.Index(strings.ToLower(x1), "i") == 0
	res2 := strings.IndexAny(strings.ToLower(x1), "a") > 0
	res3 := strings.LastIndex(strings.ToLower(x1), "n") == len(x1)-1

	if res1 && res2 && res3 {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not found!")
	}

}
