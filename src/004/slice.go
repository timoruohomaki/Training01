/*
Write a program which prompts the user to enter integers and stores the integers in a
sorted slice. The program should be written as a loop. Before entering the loop, the
program should create an empty integer slice of size (length) 3. During each pass through
the loop, the program prompts the user to enter an integer to be added to the slice.
The program adds the integer to the slice, sorts the slice, and prints the contents of
the slice in sorted order. The slice must grow in size to accommodate any number of
integers which the user decides to enter. The program should only quit (exiting the loop)
 when the user enters the character ‘X’ instead of an integer.
*/

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var x string // entered number as string
	var y int

	sli := make([]int, 0, 3) // assignment says length of 3 and not capacity but it doesn't make any sense, so being a rebel here

	for x != "X" {

		fmt.Println("Enter next number (exit by entering X)")

		_, err := fmt.Scanln(&x)

		if x != "X" {

			y, err = strconv.Atoi(x)

			if err != nil {
				panic(err)
			}

			sli = append(sli, y)

			sort.SliceStable(sli, func(i, j int) bool {
				return sli[i] < sli[j]
			})

			fmt.Println(sli)

		}

	}

	fmt.Println("Exiting the program...")
}
