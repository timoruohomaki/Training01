// Write a program which prompts the user to enter a floating point number
// and prints the integer which is a truncated version of the floating point number that was entered.
// Truncation is the process of removing the digits to the right of the decimal place.
package main
import "fmt"

func main(){
	var num float32
	fmt.Println("Enter the floating point number:")
	fmt.Scanf("%float32", &num)

	var result int32
	result = int32(num)
	fmt.Println(result)
}