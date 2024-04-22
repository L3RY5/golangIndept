package main

import "fmt"

func main() {

	var accountBalance = 1000.0
	fmt.Println("Welcome to the Go Bank")
	//until i i 2 run this loop
	//for i := 0; i < 2; i++ {
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
			accountBalance += depositeAmount // Long version ==> accountBalance = accountBalance + depositeAmount
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
			fmt.Printf("Balance updated! New amount : %.2f\n", accountBalance)

		default:
			fmt.Println("Goodbye !! ")
			fmt.Println("Thank you for using our bank! ")
			return
		}

	}

}
