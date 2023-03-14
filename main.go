package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func loanApplication(balance float64, amount float64) float64 {
	return balance - amount
}

func deposit(balance float64, amount float64) float64 {
	return balance + amount
}

func menu() int {
	reader := bufio.NewReader(os.Stdin)
	var res int

	fmt.Println("Pilih Jenis transaksi")
	fmt.Print("1.) Tarik tunai\n2.) Setor tunai\n3.) Cek saldi\n4.) Batal")
	option, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	switch strings.TrimRight(option, "\n") {
	case "1":
		res = 1
	case "2":
		res = 2
	case "3":
		res = 3
	default:
		res = 4
	}

	return res
}

func main() {
	// age := 24
	const minBalance = 100000
	// var balance float64 = 150000000

	fmt.Print("Input username:")
	reader := bufio.NewReader(os.Stdin)

	name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	fmt.Print("PIN: ")
	pin, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}

	if strings.TrimRight(pin, "\n") == "123" {
		var selectedMenu = menu()
		fmt.Println("Halo ", name)
		fmt.Println(selectedMenu)

		// fmt.Println("My age is", age)

		// fmt.Println("Min balance is Rp.", minBalance)
		// fmt.Printf("Your balance is Rp.%f\n", balance)

		// var updatedBalance = loanApplication(balance, 30000.0)
		// fmt.Printf("balance after loan Rp.%f\n", updatedBalance)
	} else {
		fmt.Println(pin)
		fmt.Println("PIN anda salah")
	}

}
