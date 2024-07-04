package main

import "fmt"

func main() {
	// printNumber is normal function that we store in a variable
	callbackFunc := printNumber

	// we passed callbackFunc as input for process
	// because printNumber satisfy or meet the data type of second input
	// that is func(int)
	process(2, callbackFunc)

	// this code is also functioning like normal
	// the difference is only the input
	// the first process second input is printNumber
	// and this function second input is a anonymous function
	process(2, func(a int) {
		fmt.Println("Calling person with ticket number:", a)
	})
}

func process(num int, callback func(int)) {
	callback(num)
}

// printNumber is just a normal function
func printNumber(i int) {
	fmt.Println("The number is", i)
}
