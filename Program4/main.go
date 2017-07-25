// Betting Algorithium that hopefully doesn't lose in a dice gambling game
package main

import "fmt"
import "math/rand"
import "time"

func main() {
	//Declare Variables based on User Preferences
	var chance float32
	var risk float32
	var nextbet float32
	var balance float32
	var pb float32
	var minbalance float32
	var minbet float32 = .00000001
	var cm1 float32
	var cm2 float32
	var chancem float32
	var roll float32
	var bethigh bool
	var lc int32
	// Get User Preferences
	fmt.Print("Your Balance Amount :")
	fmt.Scan(&balance)
	fmt.Print("Pertcentage of Balance Lowest Limit :")
	fmt.Scan(&pb)
	fmt.Print("Chance between 5.00 and 98.00 :")
	fmt.Scan(&chance)
	fmt.Print("Risk between 1000 and 5000000000 :")
	fmt.Scan(&risk)
	fmt.Print("true or false :")
	fmt.Scan(&bethigh)

	// Maths Formulaes
	nextbet = balance / risk
	minbalance = balance * pb / 100
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

	for {
		fmt.Println("Number of loss's", lc)
		if balance-nextbet < minbalance {
			break
		}
		fmt.Println("Your Balance", balance)
		//Generator Random number Between 0.00 and 99.99
		roll = rand.Float32() * 100
		fmt.Println("Dice Roll ", roll)
		// Compare Roll vs bet to see if win or lose assign wc=1 to win
		var wc int32
		if bethigh == true {
			if roll > 100-chance {
				wc = 1
			}
		}
		if bethigh == false {
			if roll < chance {
				wc = 1
			}
		}
		// Let User Know if Win or Loss and Update Balance
		if wc == 1 {
			fmt.Println("You Win ")
			balance = balance + nextbet*chancem
			nextbet = balance / risk
			minbalance = balance * pb / 100
			lc = 0
		}
		if wc != 1 {
			fmt.Println("You Lose")
			balance = balance - nextbet
			nextbet = nextbet * (1 + chance/100)
			lc++
		}

	}
	fmt.Println("Next Bet will Put you under Min Balance")

}
