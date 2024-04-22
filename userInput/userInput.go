//

package main

//
import (
	"fmt"
	"math"
)

func main() {
	// constant, the inflation rate should not change
	const inflationRate = 2.5 // of "const inflationRate float64 = 2"
	//we can change this  and remove the initial value due rto the user input
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("years: ")
	fmt.Scan(&years)

	fmt.Print("Expected return Rate: ")
	fmt.Scan(&expectedReturnRate)

	futureVal := (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	futureValReal := futureVal / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureVal)
	fmt.Println(futureValReal)
}
