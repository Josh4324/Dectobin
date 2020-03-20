package main

import "fmt"

func main() {
	var num int = 0
	for {
		fmt.Println("Dectobin - Number to Binary")
		fmt.Println("Type in 0 to quit the program ")

		fmt.Println("Please enter an integer ")
		fmt.Scan(&num)
		if num == 0 {
			println("Thank You for using Dectobin")
			break
		}
		fmt.Println("*************************************")
		fmt.Printf("%d in binary format is %b\n", num, num)
		fmt.Println("*************************************")
	}
}
