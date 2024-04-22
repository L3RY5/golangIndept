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
	fmt.Println("Earnings before tax: ", ebt)

	//earnings after taxes (Profit)
	profit := ebt - (ebt * (taxRate) / 100)
	fmt.Println("Earnings after taxes: ", profit)

	//ratio
	ration := ebt / profit
	fmt.Println("the Profit lost ratio: ", ration)
	fmt.Println("-----------------------------------------------------------")
	fmt.Printf("print with format \n")
	//Print in a specific format
	fmt.Printf("Earnings before tax: %v\n", ebt)
	fmt.Printf("Earnings after taxes is  %v, and  the Profit lost ratio is %v\n", profit, ration)
	fmt.Printf("Earnings after taxes(.0) is  %.1f, and  the Profit lost ratio(.00) is %.2f ", profit, ration)
}
