package main

// import required modules
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var (
		chance float64 = .05
		ee float64=0
		roll    float64
		count   float64
		win     float64
		result  string
		balance float64 = 100
		//tb float64 = balance*1.001
		minbet  float64 = balance / 10000
		nextbet         = minbet
		cm      float64 = .99 / chance
	)
	rand.Seed(time.Now().UnixNano())
	//start
	for {
		count++
		roll = (rand.Float64()) * 100
		minbet = balance / 10000
		cm = .99 / chance

		if roll < chance*100 {
			win = 1
			result = "Win"
			balance = balance + cm*nextbet - nextbet
		} else {
			win = 0
			result = "lose"
			balance = balance - nextbet
		}
		if win == 1 {
			nextbet = minbet
			chance = .05
			ee=1
		}
		if win == 0 && ee<cm*.75{
			ee=ee+1
			}
		if win ==0 && ee>=cm*.75{
			ee=1
			chance=chance+.05
			nextbet=nextbet*2.5
		}
		if nextbet < minbet {
			nextbet = minbet
		}
		if chance >.5 {
			chance =.5
			ee=1
			nextbet=minbet
		}
		if balance < 0 {
			break
		}
		// Print Results
		time.Sleep(1000000000 )
		fmt.Println(count,chance,cm,ee)
		//fmt.Println(nextbet)
		fmt.Printf("%.2f", roll)
		//fmt.Println(ee)
		fmt.Println(result, balance)
	}

}
