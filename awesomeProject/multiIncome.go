/*
Copy of income but with added multiple returns
 */
package main

import (
	"fmt"
	"strings"
)

func main(){
	var yearlyincome float64
	var stateTax float64
	var federalTax float64
	yearlyincome, stateTax, federalTax = Income()
	fmt.Printf("Yearly income after tax (based on federal/state average): $%.2f",yearlyincome)
	fmt.Printf("State Income Tax: $%.2f",stateTax)
	fmt.Printf("Federal Income Tax: $%.2f",federalTax)
}

func Income() (float64, float64, float64) {
	fmt.Println("Enter the number of hours you work a week:")
	var hours float64
	_, _ = fmt.Scanf("%f", &hours)
	fmt.Println("Enter your hourly wage:")
	var wage float64
	_, _ = fmt.Scanf("%f", &wage)
	fmt.Println("Enter two letter state abbreviation:")
	var state string
	_, _ = fmt.Scanf("%s", &state)
	state = strings.ToUpper(state)

	fmt.Printf("Hours worked a week: %.2f\n",hours)
	fmt.Printf("Hourly Wage: %.2f\n",wage)

	var week float64
	var year float64
	week = hours * wage
	year = week * 52.1429

	fmt.Printf("Yearly Salary(Before Tax): %.2f\n",wage)

	var federal = year * 0.114
	var stateTax float64
	if state == "MO"{
		stateTax = year * 0.037
	}else if state == "CA"{
		stateTax = year *0.0665
	}else{
		stateTax = year * 0.056
		fmt.Println("No state given default of 5.6% applied.")
	}
	var tax = federal + stateTax

	year = year - tax
	return year, stateTax, federal
}