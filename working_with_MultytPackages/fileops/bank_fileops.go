package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

// extract from a file
func ExtractFloatFromFile(floatFileName string) (float64, error) {

	data, err := os.ReadFile(floatFileName)

	//nil ==> the abcent of a usefull value you could use.
	//so this "err" wil be nil(true) if we hav NO error
	// != nil ==> not nill(FALSE) IS there is an ERR
	if err != nil {
		//use "errors" package to create a new error by calling new function and attach an error message to it
		return 1000, errors.New("failed to read file, using the default value of 1000")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)

	//nil ==> the abcent of a usefull value you could use.
	//so this "err" wil be nil(true) if we hav NO error
	// != nil ==> not nill(FALSE) IS ther is an ERR
	if err != nil {
		return 1000, errors.New("failed to parse stored float")
	}
	// we still need to return a second value even if there is no error, so nil meaning "no error"
	return value, nil

}

// function to Write balance to file
func WriteFloatToFile(value float64, fileName string) {

	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}
