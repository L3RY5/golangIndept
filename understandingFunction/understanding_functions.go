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
	outputText("Sedric", 31)

	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	//userinput for the variables
	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return Rate: ")
	fmt.Scan(&expectedReturnRate)

	//this function returns 2 values so we are going to store them in individual vars to  use them afterwards
	returnedFutureVal, returnedFutureValReal := calculateFutureValues(investmentAmount, expectedReturnRate, years)

}

// function with input argument
func outputText(name string, age int) {
	fmt.Printf("Hi my name is %v and I am %v", name, age)
}

// Version 1
// we can return the calculated value with return,
func calculateFutureValues2(investmentAmount, expectedReturnRate, years float64) (float64, float64) {
	futureVal := investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureValReal := futureVal / math.Pow(1+inflationRate/100, years)

	return futureVal, futureValReal
}

// Version 2
// if we already let the function know the variable we want to return(The return variable Names)
// then we only need to assign the value to the variable not initialize it
func calculateFutureValues(investmentAmount, expectedReturnRate, years float64) (futureVal float64, futureValReal float64) {
	futureVal = investmentAmount * math.Pow(1+expectedReturnRate/100, years)
	futureValReal = futureVal / math.Pow(1+inflationRate/100, years)

	return
	//This will work but for bigger applications and for a good overvieuw we should alsways use "return futureVal, futureValReal"
}
