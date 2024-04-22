//

package main

//
import (
	"fmt"
	"math"
)

func main() {
	// you can still put a nr with no decimal place into and flaot.
	//but internally it will ber stored and threated differently
	var investmentAmount float64 = 1000
	var expectedReturnRate = 5.5
	var years float64 = 10

	var futureVal = (investmentAmount) * math.Pow(1+expectedReturnRate/100, years)
	//fmt.Print("Test")
	//return futureVal
	fmt.Println(futureVal)

}
