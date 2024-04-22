//

package main

//

import (
	"fmt"
	"math"
)

// Varibales
const inflationRate = 2.5

func main() {

	//userinput for the variables, here we dont need to var it first above.
	//we create and initialize directly with the value of the input
	investmentAmount := getUserInput("Investment Amount")
	expectedReturnRate := getUserInput("Expected return Rate")
	years := getUserInput("years")

	//this function returns 2 values so we are going to store them in individual vars to  use them afterwards
	returnedFutureVal, returnedFutureValReal := calculateFutureValues(investmentAmount, expectedReturnRate, years)

	fmt.Printf("%.1f\n", returnedFutureVal)
	fmt.Printf("%.1f\n", returnedFutureValReal)
}

// function with input argument, het the diffrent values needed
func getUserInput(question string) float64 {
	var userInput float64
	fmt.Println(question)
	fmt.Scan(&userInput)
	return userInput
}

// if we already let the function know the variable we want to return(The return variable Names)
// then we only need to assign the value to the variable not initialize it
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureVal float64, futureValReal float64) {
	futureVal = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureValReal = futureVal / math.Pow(1+inflationRate/100, years)

	return futureVal, futureValReal
}
