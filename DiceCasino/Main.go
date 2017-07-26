package main
import "fmt"
import "time"
import "math/rand"
func main () {
	//Declair global varibles
	var chance float64 = 12.38
	var balance float64
	var minbet float64 = balance / 2500000000
	var bethigh bool = false
	var nextbet float64 = minbet
	var ld int64 = 0
	var nl int64 = 0
	var echance int64 = 0
	var minbalance float64 = balance * .65
	//error check make sure next bet is possible
	if  nextbet < .00000001 {
		nextbet = .00000001
	}
	//Bets start here
	for {
		rand.Seed(time.Now().UnixNano())
		ld = ld + 1
		fmt.Println("Last Double Win")
		fmt.Println(ld)
		fmt.Println("rolls ago")
		fmt.Println("Number of Loss's in a Row")
		fmt.Println(nl)
		if win | echance > 0 {
			nextbet = nextbet * 3.35
			echance = echance + 1
			nl = 0
			ld = 0
		}
		if win | echance == 0 {
			nextbet = nextbet * .5
			echance = echance + 1
			minbalance = balance * .65
			nl = 0
		}
		if !win | echance > 0 {
			nextbet = minbet
			echance = 0
			nl = nl + 1
		}
		if !win | echance == 0 {
		nextbet = nextbet * (1.235)
		nl = nl + 1
		echance = 0
		}
		if ld <= 50000 {
		minbet = balance / 2500000000
		}
		if ld > 50000 {
		minbet = balance / 9000000000
		}
		if minbet < .00000001 {
			minbet = .00000001
		}
		if nextbet < minbet {
			nextbet = minbet
		}
		if balance < minbalance {
			fmt.Println("Really Bad Luck Better Luck Next Time")
			break
		}
	}
}