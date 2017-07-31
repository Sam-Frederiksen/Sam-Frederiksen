package main

import "fmt"
import "time"
import "math/rand"

func main() {
	var (
		trolls  int64
		bm      float64
		risk    float64
		risk2   float64
		wtu1    float64
		wtu2    float64
		wtu3    float64
		wtu4    float64
		ld      int64
		nl      float64
		win     int64
		roll    float64
		chance  float64
		cm1     float64
		cm2     float64
		chancem float64
		balance float64
	)
	//reset seed
	rand.Seed(time.Now().UnixNano())
	//Get User to Add balance
	fmt.Print("Enter Your Balance :")
	fmt.Scan(&balance)
	fmt.Print("Increase on Lose Percent 1.405 (1 to 1000)")
	fmt.Scan(&bm)
	fmt.Print("Risk Pertcentage of Balance .95 (.99 to .01) :")
	fmt.Scan(&risk)
	fmt.Print("Minbet Size Reducer  1000000(1 to 1000000000) :")
	fmt.Scan(&risk2)
	fmt.Print("Win  adjuster  .5(0 to 100) :")
	fmt.Scan(&wtu1)
	fmt.Print("Multiple Win adjuster  3.05(0 to 100) :")
	fmt.Scan(&wtu2)
	fmt.Print("Fine Tune Multiple WIn 0(1 to 1000000000) :")
	fmt.Scan(&wtu3)
	fmt.Print("Fine Tune First Win  0(1 to 1000000000) :")
	fmt.Scan(&wtu4)
	fmt.Print("Roll Under 15(5.00 to 98) :")
	fmt.Scan(&chance)

	var (
		minbet     float64 = balance / risk2
		nextbet    float64 = minbet
		echance    float64 = 0
		minbalance float64 = balance * risk
	)
	// set up house edge
	cm1 = 100 / chance
	cm2 = cm1 * .01
	chancem = cm1 - cm2
	//error check make sure next bet is possible
	if nextbet < .00000001 {
		nextbet = .00000001
	}
	//Bets start here
	for {
		roll = rand.Float64() * 100
		trolls++
		// Determine if bet is a win or loss
		if roll < chance {
			win = 1
			balance = balance + nextbet*chancem - nextbet
			minbalance = balance * risk
			nl = 0
		} else {
			win = 0
			balance = balance - nextbet
			nl++

		}
		ld++
		// If We already had more than 2 wins in a row
		if win == 1 && echance > 0 {
			nextbet = nextbet * (wtu2 - (echance * wtu3))
			echance++
			ld = 0
		}
		// If this is 2nd win in a row
		if win == 1 && echance == 0 {
			nextbet = nextbet * wtu1
			echance++
		}
		// If not win after more than 1 win
		if win == 0 && echance > 0 {
			nextbet = minbet
			echance = 0
		}
		// Normal loss routine
		if win == 0 && echance == 0 {
			nextbet = nextbet * (bm + (nl * wtu4))
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
		if trolls % 100 ==0 {
		//Display Stats every 1000 rolls
		fmt.Println("Total Number Rolls", trolls)
		fmt.Println("Balance", balance)
		fmt.Print("Number of Loss's in a Row ")
		fmt.Println(nl)
			time.Sleep(100 * time.Millisecond)
			}
		if balance < minbalance {
			fmt.Print("You Just Hit Your Min Balance Level")
			fmt.Println("Total Number Rolls", trolls)
			fmt.Println("Balance ", balance, " Bets Reset to MinBet")
			fmt.Print("Number of Loss's in a Row ")
			fmt.Println(nl)
			time.Sleep(1000 * time.Millisecond)
			nextbet = minbet
			minbalance = balance * risk
		}
	}
}
