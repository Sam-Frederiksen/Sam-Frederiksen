package main

import "fmt"

func main() {
	var (
		i  int
		n  int
		c1 int
		c2 int
	)
	for i = 1000000; i > 1; i-- {
		n=i
		for {
			c1++
			if n == 1 {
				if c1>c2{
					c2=c1
				}
				fmt.Println(c2,i)
				c1=0
				break
			}
			if n<1{
				c1=0
				break
			}

			if n%2 == 0 {
				n = n / 2
			} else {
				n = 3*n + 1
			}
		}
	}
}
