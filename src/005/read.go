/*
Write a program which reads information from a file and represents it in a slice of
structs. Assume that there is a text file which contains a series of names. Each line
of the text file has a first name and a last name, in that order, separated by a single
space on the line.

Your program will define a name struct which has two fields, fname for the first name,
and lname for the last name. Each field will be a string of size 20 (characters).

Your program should prompt the user for the name of the text file. Your program will
successively read each line of the text file and create a struct which contains the
first and last names found in the file. Each struct created will be added to a slice,
and after all lines have been read from the file, your program will have a slice
containing one struct for each line in the file. After reading all lines from the file,
your program should iterate through your slice of structs and print the first and last
names found in each struct.
*/

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	var tfile string

	fmt.Println("Enter file name:")
	fmt.Scanln(&tfile)

	file, err := os.Open(tfile)

	if err != nil {
		log.Fatalf("[read.go] Error while opening file: %s", err)
	}

	tscan := bufio.NewScanner(file)
	tscan.Split(bufio.ScanLines)

	sliName := make([]Name, 0)

	for tscan.Scan() {

		var n Name

		s := strings.Split(tscan.Text(), " ")

		n.fname = s[0]
		n.lname = s[1]

		sliName = append(sliName, n)

	}

	file.Close()

	// PRINT THE CONTENTS OF SLICE

	for _, eachline := range sliName {

		fmt.Println(eachline.fname + " " + eachline.lname)
	}

}
