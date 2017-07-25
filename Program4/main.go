// Betting Algorithium that hopefully doesn't lose in a dice gambling game
package main

import "fmt"

func main() {
	//Declare Variables based on User Preferences
	var chance float32
	var risk float32
	var nextbet float32
	var balance float32
	var minbalance float32
	var minbet float32 = .00000001
	var cm1 float32
	var cm2 float32
	var chancem float32
	// Get User Preferences
	fmt.Print("Your Balance Amount :")
	fmt.Scan(&balance)
	fmt.Print("Lowest Balance You Will Go :")
	fmt.Scan(&minbalance)
	fmt.Print("Chance between 5.00 and 98.00 :")
	fmt.Scan(&chance)
	fmt.Print("Risk between 1000 and 5000000000 :")
	fmt.Scan(&risk)
	// Maths Formulaes
	nextbet = balance / risk
	cm1 = 100 / chance
	cm2 = cm1 * .01
	chancem = cm1 - cm2 //chance multiplier based on house 1 percent edge
	//Show User Current Bet Settings
	fmt.Println("Balance :", balance)
	fmt.Println("Min Balance you will go", minbalance)
	fmt.Println("minbet = ", minbet)
	fmt.Println("Your next bet = ", nextbet)
	fmt.Println("Chance and NextBet Multipler", chance, chancem)

	// Program Will Run Until we hit min balance

	for balance <= minbalance {
        fmt.Println("Your Balance",balance)

	}

}
