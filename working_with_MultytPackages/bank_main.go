package main

import (
	"fmt"

	"example.com/multipackages/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {

	//instead of hardcoding here we will get it from file
	var accountBalance, err = fileops.ExtractFloatFromFile(accountBalanceFile)

	// now we call the err again and attach the error meassage so that it return the error
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err) //the right error message failed to ...
		fmt.Println("------------------------------------------------------")
		//panic("Error To great, cant continue, Shuting down application")
	}

	fmt.Println("Welcome to the Go Bank")

	for {

		greetMenu()

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
			accountBalance += depositeAmount                             // Long version ==> accountBalance = accountBalance + depositeAmount
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile) //write it to the file
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile) //write it to the file
			fmt.Printf("Balance updated! New amount : %.2f\n", accountBalance)

		default:
			fmt.Println("Goodbye !! ")
			fmt.Println("Thank you for using our bank! ")
			return
		}

	}

}
