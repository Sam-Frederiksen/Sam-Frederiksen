package main

import "fmt"

func main() {
	var (
		i int
		a int
		b int
		c int
		g int
		)
	for {
		if i > 500 {
			break
		}
		a++
		b = b + a
			g = 0
		for c = 1; c < b/2; c++ {
			if b%c == 0 {
				g++
						}
		}
		if g > 500 {
			fmt.Print(b)
			i = g
		}
	}
}
