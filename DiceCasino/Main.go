package main

import "fmt"
import "time"
import "math/rand"

func main() {
	//Declare global varibles
	var trolls int64
	var bm float64 = 1.4
	var risk float64 = .98 //minbalance Percentage
	var risk2 float64 = 100000000
	var ld int64 = 0
	var nl int64 = 0
	var win int64
	var roll float64
	var chance float64 = 12.38
	var cm1 float64
	var cm2 float64
	var chancem float64
	cm1 = 100 / chance
	cm2 = cm1 * .01
	chancem = cm1 - cm2 //chance multiplier based on house 1 percent edge
	var balance float64
	var minbet float64 = balance / risk2
	var nextbet float64 = minbet
	var echance int64 = 0
	var minbalance float64 = balance * risk
	//error check make sure next bet is possible
	if nextbet < .00000001 {
		nextbet = .00000001
	}
	//reset seed
	rand.Seed(time.Now().UnixNano())
	//Get User to Add balance
	fmt.Print("Your Balance")
	fmt.Scan(&balance)
	//Bets start here
	for {
		// Determine if bet is a win or loss
		roll = rand.Float64() * 100
		trolls++
		if roll < chance {
			win = 1
			balance = balance + nextbet*chancem
		} else {
			win = 0
			balance = balance - nextbet
		}
		// Print Stats
		ld++
		// If We already had more than 2 wins in a row
		if win == 1 && echance > 0 {
			nextbet = nextbet * 3.35
			echance++
			nl = 0
			ld = 0
		}
		// If this is 2nd win in a row
		if win == 1 && echance == 0 {
			nextbet = nextbet * .5
			echance++
			minbalance = balance * risk
			nl = 0
		}
		// If not win after more than 1 win
		if win == 0 && echance > 0 {
			nextbet = minbet
			echance = 0
			nl++
		}
		// Normal loss routine
		if win == 0 && echance == 0 {
			nextbet = nextbet * bm
			nl++
			echance = 0
		}
		//setup new minbet
		minbet = balance / risk2
		//errorcheck to see if minbet is possible
		if minbet < .00000001 {
			minbet = .00000001
		}
		//error check to see if nextbet is possible
		if nextbet < minbet {
			nextbet = minbet
		}
		//check to see if we hit min balance or not
		if balance < minbalance {
			fmt.Println("Total Number Rolls", trolls)
			fmt.Println("Balance", balance)
			fmt.Print("Number of Loss's in a Row ")
			fmt.Println(nl)
			time.Sleep(100 * time.Millisecond)
			nextbet = minbet
			minbalance = balance * risk
		}
	}
}
