package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	//version 1, with ever variable we can create a custom mesage ect, but then the if err needs to be done for all inputrs
	revenue, err := getUserInput("Revenue: ")
	if err != nil {
		fmt.Println("Revenu ERROR")
		fmt.Println(err) //the right error message failed to ...
		fmt.Println("------------------------------------------------------")
		return //will stop the program after an error
	}

	//version 2, dynamic, only 1 "if err". we do this by given Err == diffrent, names and using OR statement in if
	expenses, err1 := getUserInput("Expenses: ")
	taxRate, err2 := getUserInput("Tax Rate: ")
	if err1 != nil || err2 != nil {
		fmt.Println(err) //the message 'INPUT ...' will not show
		fmt.Println("------------------------------------------------------")
		//panic("INPUT value error, application forced to stop")
		return
	}

	ebt, profit, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Printf("%.1f\n", ebt)
	fmt.Printf("%.1f\n", profit)
	fmt.Printf("%.3f\n", ratio)
	//store it into a file
	storeResults(ebt, profit, ratio)
}

// functions
func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64) {
	ebt := revenue - expenses
	profit := ebt * (1 - taxRate/100)
	ratio := ebt / profit
	return ebt, profit, ratio
}

func getUserInput(infoText string) (float64, error) {
	var userInput float64
	fmt.Print(infoText)
	fmt.Scan(&userInput)
	if userInput <= 0 {
		return 0, errors.New("ERROR : Input must be a positive number")
	}

	/*_, err := fmt.Scan(&userInput)
	    if err != nil {

		}*/

	return userInput, nil
}

func storeResults(ebt, profit, ratio float64) {
	results := fmt.Sprintf("EBT : %.1f\nPROFIT : %.1f\nRATIO : %.3f\n", ebt, profit, ratio)
	os.WriteFile("result.txt", []byte(results), 0644)
}
