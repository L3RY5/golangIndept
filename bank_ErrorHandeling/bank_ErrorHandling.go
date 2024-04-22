package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance2.txt"

// extract from a file
func extractBalanceFromFile() (float64, error) {

	data, err := os.ReadFile(accountBalanceFile)

	//nil ==> the abcent of a usefull value you could use.
	//so this "err" wil be nil(true) if we hav NO error
	// != nil ==> not nill(FALSE) IS ther is an ERR
	if err != nil {
		//use "errors" package to create a new error by calling new function and attach an error message to it
		return 1000, errors.New("failed to read file, using the default value of 1000")
	}

	balancedText := string(data)
	balance, err := strconv.ParseFloat(balancedText, 64)

	//nil ==> the abcent of a usefull value you could use.
	//so this "err" wil be nil(true) if we hav NO error
	// != nil ==> not nill(FALSE) IS ther is an ERR
	if err != nil {
		return 1000, errors.New("failed to parse stored balance")
	}
	// we still need to return a second value even if there is no error, so nil meaning "no error"
	return balance, nil

}

// function to Write balance to file
func writeBalanceToFile(balance float64) {

	balancedText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balancedText), 0644)
}

func main() {

	//instead of hardcoding here we will get it from file
	var accountBalance, err = extractBalanceFromFile()

	// now we call the err again and attach the error meassage so that it return the error
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err) //the right error message failed to ...
		fmt.Println("------------------------------------------------------")
		panic("Error To great, cant continue, Shuting down application")
	}

	fmt.Println("Welcome to the Go Bank")

	for {
		fmt.Println(" ")
		fmt.Println("------------------------------------------------------")
		fmt.Println("WHAT DO YOU WANT TO DO ?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw money")
		fmt.Println("4. Exit")

		var choice int
		fmt.Print("Your Choice : ")
		fmt.Scan(&choice)
		fmt.Printf("You chose : %v\n", choice)
		//condition ==> switch
		switch choice {
		case 1:
			fmt.Println("Your balance is : ", accountBalance)

		case 2:
			fmt.Print("Your Deposit : ")
			var depositeAmount float64
			fmt.Scan(&depositeAmount)
			//see if the deposited amont is smaller or equal to zero
			if depositeAmount <= 0 {
				fmt.Println("Invalid amount. must be greater than 0.")
				//return //empty return will let the code stop and not continue in this if part
				continue //it will skip the rest of THIS iteration and then restart the next iteration of the LOOP
			}
			accountBalance += depositeAmount   // Long version ==> accountBalance = accountBalance + depositeAmount
			writeBalanceToFile(accountBalance) //write it to the file
			fmt.Printf("Balance updated! New amount : %.2f\n", accountBalance)

		case 3:
			fmt.Print("Your withdraw : ")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			if withdrawAmount <= 0 {
				fmt.Println("Invalid amount. must be greater than 0.")
				continue //it will skip the rest of THIS iteration and then restart the next iteration of the LOOP
			}

			if withdrawAmount > accountBalance {
				fmt.Printf("Invalid amount. amount must be smaller or equal to %.2f.", accountBalance)
				//return ==> empty return will let the code stop and not continue in this if part
				continue //it will skip the rest of THIS iteration and then restart the next iteration of the LOOP
			}
			accountBalance -= withdrawAmount
			writeBalanceToFile(accountBalance) //write it to the file
			fmt.Printf("Balance updated! New amount : %.2f\n", accountBalance)

		default:
			fmt.Println("Goodbye !! ")
			fmt.Println("Thank you for using our bank! ")
			return
		}

	}

}
