//

package main

//
import (
	"fmt"
	"math"
)

func main() {
	// 1 __bundel the variables with the same explicit or inverd type
	//var investmentAmount, years = 1000, 10 ==> here the dont have to be of the same type so var investmentAmount, years = 1000, "10" will work to
	// OR
	// 2__
	//var investmentAmount, years  float64= 1000, 10

	// OR
	// 3 __ the short Version, then you need to make sure its the type you need "" or . etc ==>
	//investmentAmount, years  := 1000.0, 10.0

	// OR
	// 4__  Bundle everything all together. tey are of the same type
	investmentAmount, years, expectedReturnRate := 1000.0, 10.0, 5.5

	var futureVal = (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	//fmt.Print("Test")
	//return futureVal
	fmt.Println(futureVal)

}
