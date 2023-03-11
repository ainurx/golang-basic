package main

import (
	"fmt"
	"math"
)

func loanApplication(balance float64, amount float64) float64 {
	return balance - amount
}

func deposit(balance float64, amount float64) float64 {
	return balance + amount
}

func main() {
	var name string = "ainur"
	age := 24
	const minBalance = 100000
	var balance float64 = 150000000

	fmt.Println("Hello", name)
	fmt.Println("My age is", age)

	fmt.Println("Min balance is Rp.", minBalance)
	fmt.Printf("Your balance is Rp.%f\n", balance)

	var updatedBalance = loanApplication(balance, 30000.0)
	fmt.Printf("balance after loan Rp.%f\n", updatedBalance)
	fmt.Printf("balance after loan Rp.", math.Round(updatedBalance))
}
