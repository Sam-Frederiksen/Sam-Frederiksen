package main

import "fmt"

func main() {
	var (
		i int = 3
		c int = 2
		b int
		f int
	)
	for {
		i++
		for b = 2; b < i; b++ {
			if i%b == 0 {
				//fmt.Println(i, "Not Prime")
				f = 1
				break
			}
		}
		if f == 0 {
			fmt.Println(i, "Prime")
			c=c+1
		}
		f = 0
		if c == 10001 {
			fmt.Println(i)
			break
		}
	}
}
