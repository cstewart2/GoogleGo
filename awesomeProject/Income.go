/*
Salary and Tax calculator based on national and state averages(only programmed for MO and CA) that takes user input
 */
package main

import (
	"fmt"
	"strings"
)

func main(){
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
	var yearlyincome = YearlyIncome(hours,wage,state)
	fmt.Printf("Hours worked a week: %.2f\n",hours)
	fmt.Printf("Hourly Wage: %.2f\n",wage)
	fmt.Printf("Yearly income after tax (based on american average): $%.2f",yearlyincome)
	
}

func YearlyIncome(hours float64, wage float64, state string) float64 {
	var week float64
	var year float64 
	week = hours * wage
	year = week * 52.1429
	var tax = IncomeTax(year, state)
	year = year - tax
	return year
}

func IncomeTax(year float64, state string) float64{
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
	return tax
}
