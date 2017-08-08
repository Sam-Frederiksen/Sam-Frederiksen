package main

import "fmt"

func main() {
	var (
		i int
		j int
		d int
		k int
	)
	for {
		if k==1 {
			break
		}
		j++
		d = 0
		for i = 1; i < 21; i++ {
			if j%i == 0 {
				d++
			}
			if d > 19 {
				fmt.Print(j)
				k=1
			}

		}
	}
}
