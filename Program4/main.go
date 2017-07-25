package main

import "fmt"

func main() {
	var chance float32
	var risk float32
	var nextbet float32
	var balance float32
	fmt.Print("Your Balance Amount :")
	fmt.Scan(&balance)
	fmt.Print("Chance between 5.00 and 98.00 :")
	fmt.Scan(&chance)
	fmt.Print("Risk between 1000 and 5000000000 :")
	fmt.Scan(&risk)
	minbet := .00000001
	nextbet = (balance / risk)
	fmt.Println("Balance :",balance)
	fmt.Println("minbet = ", minbet)
	fmt.Println("Your next bet = ", nextbet)

}
