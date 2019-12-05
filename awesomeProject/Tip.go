/*
Basic tip calculator that takes user input
 */
package main

import (
	"fmt"
	"os"
)

func main(){
	fmt.Println("Enter the amount of your bill:")
	var bill float64
	_, _ = fmt.Scanf("%f", &bill)
	fmt.Println("Enter the percentage you would like to tip:")
	var percent float64
	_, _ = fmt.Scanf("%f", &percent)
	var decimal = percent/100
	var tip = (bill*decimal)
	var total = tip + bill
	fmt.Printf("Subtotal: $%.2f\n",bill)
	fmt.Printf("Tip:: $%.2f\n",tip)
	fmt.Printf("Total: $%.2f",total)
	os.Exit(0)
}
