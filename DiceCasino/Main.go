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
	var maxbalance float64 = balance * 1000
	if  nextbet < .00000001 {
		nextbet = .00000001
	}
	for {
		rand.Seed(time.Now().UnixNano())
		ld = ld + 1
		fmt.Println("Last Double Win")
		fmt.Println(ld)
		fmt.Println("rolls ago")
		fmt.Println("Number of Loss's in a Row")
		fmt.Println(nl)
		if win and
		echance > 0
		then
		nextbet = nextbet * 3.35
		echance = echance + 1
		nl = 0
		ld = 0
		end
		if win and
		echance == 0
		then
		nextbet = nextbet * .5
		echance = echance + 1
		minbalance = balance * .65
		maxbalance = balance * 1000
		nl = 0
		end

		if !win and
		echance > 0
		then
		nextbet = minbet
		echance = 0
		nl = nl + 1
		end
		if !win and
		echance == 0
		then
		nextbet = nextbet * (1.235)
		nl = nl + 1
		echance = 0
		end
		if ld <= 50000 then
		minbet = balance / 2500000000
		end

		if ld > 50000 then
		minbet = balance / 9000000000
		end

		if minbet < .00000001 then
		minbet = .00000001
		end
		if nextbet < minbet then
		nextbet = minbet
		end
		if balance < minbalance then
		fmt.Println("Really Bad Luck Better Luck Next Time")
		stop()
		end
		if balance > maxbalance then
		fmt.Println("Target Achieved Well Done")
		stop()
		end

	}
}