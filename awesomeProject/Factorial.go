/*
This program calculates the factorial of a given number from user input
 */
package main

import "fmt"

func main() {
	var number int

	fmt.Print("Enter a number under 100: ")
	fmt.Scan(&number)
	fmt.Print("Factorial is: ",Factorial(number))
}

func Factorial(number int) uint64 {
	var factorial uint64 = 1
	if number < 0{
		fmt.Println("Number may not be negative.")
	}else {
		for i := 1; i <= number; i++ {
			factorial *= uint64(i)
		}
	}
		return factorial
}