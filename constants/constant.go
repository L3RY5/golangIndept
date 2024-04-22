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
	var investmentAmount float64 = 1000
	expectedReturnRate := 5.5
	years := 10.0
	// in some scenarios this could change multiple times depending  on the type of problem the app need to solve
	investmentAmount = 2000
	futureVal := (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	futureValReal := futureVal / math.Pow(1+inflationRate/100, years)
	fmt.Println(futureVal)
	fmt.Println(futureValReal)
}
