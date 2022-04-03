package main

import "fmt"

func countDownUp(number1, number2 int) {
	var startNum, endNum int
	if number1 >= number2 {
		startNum = number1
		endNum = number2
	} else {
		startNum = number2
		endNum = number1
	}

	for i := startNum; i > endNum-1; i-- {
		fmt.Println(i)
	}
	for i := endNum; i < startNum+1; i++ {
		fmt.Println(i)
	}

	return
}

func main() {
	//Insert code here
	var number1, number2 int
	fmt.Println("Enter the first number:")
	fmt.Scanln(&number1)
	fmt.Println("Enter the second number:")
	fmt.Scanln(&number2)

	countDownUp(number1, number2)
}
