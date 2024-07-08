package main

import (
	"errors"
	"fmt"
)

func main() {
	// var intNum int = 32767
	// intNum = intNum + 1
	// fmt.Println(intNum)

	// var floatNum float64 = 12345678.9
	// fmt.Println(floatNum)

	// var myString string = "Hello Sagar!"
	// fmt.Println(len(myString))

	// myVar := "New to code"
	// myNum := 100
	// fmt.Print(myVar)
	// fmt.Println(myNum)
	// fmt.Println("Hello Sagar !")
	myPrint()
	var myRandomString string = "Hello this is my random string "
	printMe(myRandomString)

	numberOne := 0
	numberaTwo := 90
	var result, reminder, err = intDivision(numberOne, numberaTwo)
	if err != nil {
		fmt.Printf(err.Error())
	} else if reminder == 0 {
		fmt.Printf("The result of the integer division is %v", result)
	} else {
		fmt.Printf("The result of the lanagau of the divison of the intiger is %v, and the remider is %v ", result, reminder)
	}
}

func myPrint() {
	fmt.Println("Hello i m inside a function call from main")
}

func printMe(printMeVar string) {
	fmt.Println(printMeVar)
}

func intDivision(numerator int, denominirator int) (int, int, error) {
	var err error
	if denominirator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominirator
	var reminder int = numerator % denominirator
	return result, reminder, err
}
