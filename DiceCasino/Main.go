package main
import "fmt"
import "time"
import "math/rand"
func main () {
	//Declare global varibles
	var chance float64 = 12.38
	var balance float64
	var minbet float64 = balance / 2500000000
	var bethigh bool = false
	var nextbet float64 = minbet
	var ld int64 = 0
	var nl int64 = 0
	var echance int64 = 0
	var minbalance float64 = balance * .65
	var win int64
	var roll float64
	//error check make sure next bet is possible
	if  nextbet < .00000001 {
		nextbet = .00000001
	}
	//reset seed
	rand.Seed(time.Now().UnixNano())
	//Print User settings
	fmt.Println("Betting at 12.38% - Roll Over =",bethigh)
	//Bets start here
	for {
		// Determine if bet is a win or loss
		roll = rand.Float64()
		if roll < chance {
			win = 1
			} else {
			win = 0
			}
		// Print Stats
		ld = ld + 1
		fmt.Println("Last Double Win")
		fmt.Println(ld)
		fmt.Println("rolls ago")
		fmt.Println("Number of Loss's in a Row")
		fmt.Println(nl)

		// If We already had more than 2 wins in a row
		if win == 1 && echance > 0 {
			nextbet = nextbet * 3.35
			echance = echance + 1
			nl = 0
			ld = 0
		}
		// If this is 2nd win in a row
		if win == 1 && echance == 0 {
			nextbet = nextbet * .5
			echance = echance + 1
			minbalance = balance * .65
			nl = 0
		}
		// If not win after more than 1 win
		if win == 0 && echance > 0 {
			nextbet = minbet
			echance = 0
			nl = nl + 1
		}
		// if first win after 1 or more loss's
		if win == 0 && echance == 0 {
		nextbet = nextbet * (1.235)
		nl = nl + 1
		echance = 0
		}
		//setup new minbet
		minbet = balance / 2500000000
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
			fmt.Println("Really Bad Luck Better Luck Next Time")
			break
		}
	}
}