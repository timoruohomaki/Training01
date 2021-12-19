/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON
object from the map, and then your program should print the JSON object.
*/

package main

import (
	"encoding/json"
	"fmt"
)

var name string
var address string

func main() {
	fmt.Println("Enter name:")
	_, err1 := fmt.Scanln(&name)
	if err1 != nil {
		panic(err1)
	}
	fmt.Println("Enter address:")
	_, err2 := fmt.Scanln(&address)
	if err2 != nil {
		panic(err2)
	}

	pMap := make(map[string]string)

	pMap["name"] = name
	pMap["address"] = address

	barr, err3 := json.Marshal(pMap)

	if err3 != nil {
		panic(err3)
	}

	fmt.Println("")
	fmt.Println("Thank you. Person object in JSON notation is: ")
	fmt.Println("")
	fmt.Println(string(barr))
}
