package main

import "fmt"
import "time"
import "math/rand"

func main() {
	var (
		//trolls   int64
		nrolls int64
		win      int64
		roll     float64
		chance   float64 = 40
		cm1      float64
		cm2      float64
		chancem  float64
		balance  float64 = 100
		minbet   float64 = balance / 10000000
		mytarget         = balance
		c1               = 0
		c2               = 0
		nw       float64 = 0
		nextbet  float64 = minbet
	)
	//reset seed
	rand.Seed(time.Now().UnixNano())
	// set up house edge
	cm1 = 100 / chance
	cm2 = cm1 * .001
	chancem = cm1 - cm2
	//error check make sure next bet is possible
	if nextbet < .00000001 {
		nextbet = .00000001
	}
	for {
		roll = rand.Float64() * 100
		//trolls++
		nrolls++
		if balance < 0 {
			fmt.Println("out of money")
			break
		}
		if roll < chance {
			win = 1
			balance = balance + nextbet*chancem - nextbet
		} else {
			win = 0
			balance = balance - nextbet
		}
		minbet = balance / 600
		if win == 1 {
			nextbet = nextbet * 1.1
			nw=0
		}
		if win == 0 {
			nextbet = nextbet * 1.1
			nw=nw+1
		}

		if nw > 3{
			nextbet = minbet
			nw = 0
			//trolls = 0
			c1 = 0
			c2 = 0
		}
		if balance > mytarget*(1.1) {
			c1 = c1 + 1
		}
		if win == 0 && c1 > 1 {
			nextbet = minbet
			mytarget = balance
			c1 = 0
			//trolls = 0
			nw = 0
			fmt.Println(nrolls)
			fmt.Println(balance)
			time.Sleep(10 * time.Millisecond)
		}

		if balance < mytarget*(.9) {
			c2 = c2 + 1
		}
		if win == 0 && c2 > 1 {
			nextbet = minbet
			mytarget = balance
			c2 = 0
			nw = 0
			//trolls = 0
			fmt.Println(nrolls)
			fmt.Println(balance)
			time.Sleep(10 * time.Millisecond)
		}
	}
}
