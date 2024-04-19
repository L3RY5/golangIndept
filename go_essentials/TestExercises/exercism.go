package main

import (
	"fmt"
	"strings"
)

func main() {

	WelcomeMsg := WelcomeMessage("sedric")
	fmt.Println(WelcomeMsg)

	borderedText := AddBorder("Welcome!", 10)
	fmt.Println(borderedText)

	CleanedMessage := CleanupMessage(borderedText)
	fmt.Println(CleanedMessage)
}

// /func
func WelcomeMessage(customer string) string {
	//panic("Please implement the WelcomeMessage() function")
	welcomeMsg := "Welcome to the Tech Palace, " + strings.ToUpper(customer)
	return welcomeMsg
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	//panic("Please implement the AddBorder() function")
	stareBorders := strings.Repeat("*", numStarsPerLine)
	textWithoutBorder := welcomeMsg
	//textWitBorder := fmt.Printf("%s\n%s\n%s",stareBorders,textWithoutBorder,stareBorders)
	return stareBorders + "\n" + textWithoutBorder + "\n" + stareBorders
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	return strings.Trim(strings.ReplaceAll(oldMsg, "*", ""), " \n")
}
