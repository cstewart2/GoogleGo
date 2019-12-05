/*
This program shows off Go's multiple return feature by calculating and returning both the area/circumference of a circle.
 */
package main

import "fmt"

func main(){
	var area float64
	var circumference float64
	area, circumference = Circle()

	fmt.Printf("Area of Circle is : %.2f",area)
	fmt.Printf("Circumference of Circle is : %.2f",circumference)
}
func Circle() (float64,float64) {
	var radius float64
	const PI float64 = 3.14
	fmt.Print("Enter radius of the Circle : ")
	fmt.Scanln(&radius)

	var area float64
	var circumference float64
	area = PI * radius * radius
	circumference = 2 * PI * radius

	return area,circumference
}
