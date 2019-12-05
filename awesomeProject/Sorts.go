/*
Program showing off the built in sort library
 */
package main

import (
	"fmt"
	"sort"
)
func main() {
	num := []int{10,2040, 120, 10000, 56}
	fmt.Println(num)

	if sort.IntsAreSorted(num)==false{
		sort.Ints(num)
	}
	fmt.Println(num)
	fmt.Println(sort.SearchInts(num,2040))

	strings := []string{"Red","Blue","Yellow","Orange","Purple"}
	fmt.Println(strings)
	if sort.StringsAreSorted(strings)==false{
		sort.Strings(strings)
	}
	fmt.Println(strings)
	fmt.Println(sort.SearchStrings(strings,"Yellow"))

	decimals := []float64{0.5,30,5.2,50.8,64.7,11.01,11.2}
	fmt.Println(decimals)
	if sort.Float64sAreSorted(decimals)==false{
		sort.Float64s(decimals)
	}
	fmt.Println(decimals)
	fmt.Println(sort.SearchFloat64s(decimals,5.2))
}
