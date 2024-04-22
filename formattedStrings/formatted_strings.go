//

package main

//

import (
	"fmt"
)

func main() {

	// the different variables we will need
	var revenue float64
	var expenses float64
	var taxRate float64

	fmt.Print("what is your yearly revenue?: ")
	fmt.Scan(&revenue)

	fmt.Print("what is your yearly expenses?: ")
	fmt.Scan(&expenses)

	fmt.Print("what is your tax rate?: ")
	fmt.Scan(&taxRate)

	//calculate the Earnings before tax (EBT)
	ebt := revenue - expenses

	//earnings after taxes (Profit)
	profit := ebt - (ebt * (taxRate) / 100)

	//ratio
	ration := ebt / profit

	//save the formatted strings in a variable
	formattedEBT := fmt.Sprintf("Earnings before tax: %v\n", ebt)
	formattedProfitNRation := fmt.Sprintf("Earnings after taxes(.0) is  %.1f, and  the Profit lost ratio(.00) is %.2f ", profit, ration)

	print(formattedEBT, formattedProfitNRation)
	fmt.Println(" ")

	// by adding backtick we can create multi-line strings
	print(`This is the first line of this string
	This is the second line of this string
	This is the third line of the string`)
}
