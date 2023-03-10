package main

import "fmt"

func loanApplication(balance float32, amount float32) float32 {
	return balance - amount
}

func deposit(balance float32, amount float32) float32 {
	return balance + amount
}

func main() {
	var name string = "ainur"
	age := 24
	const minBalance = 100000
	var balance float32 = 150000000

	fmt.Println("Hello", name)
	fmt.Println("My age is", age)

	fmt.Println("Min balance is Rp.", minBalance)
	fmt.Printf(`Your balance is Rp.%f\n`, balance)
}
