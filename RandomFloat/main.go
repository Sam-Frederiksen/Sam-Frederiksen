package main

// import required modules
import (
	"fmt"
	"math/rand"
	"time"
)

func r() (r float64) {
	r = (rand.Float64()) * 100
	return r

}

func main() {
	var (
		chance  float64 = .50
		count float64
		win     float64
		result  string
		balance float64 = 100
		tb float64 = balance*1.001
		minbet float64 = balance / 100000
		nextbet=minbet
		cm      float64 = .999 / chance
	)
	rand.Seed(time.Now().UnixNano())
	//start
	for {
		count++
		roll := r()
		minbet=balance/100000

		if roll < chance*100 {
			win = 1
			result = "Win"
			balance += cm*nextbet - nextbet
		} else {
			win = 0
			result = "lose"
			balance -= nextbet
		}
		if win ==1{
			nextbet=nextbet*1.02
		}
		if win==0{
			nextbet=nextbet* 1.02
		}
		if balance >tb && win ==0{
			tb = balance*1.0001
			nextbet = minbet

			// Print Results
			fmt.Println(count)
			fmt.Printf("%.2f", roll)
			fmt.Println(result, balance)
		}
	}
}
