package main

import "fmt"

func main() {
	var (
		a int
		b int
		c int
		d int
		f int
	)
	fmt.Println(a, b, c, d)
	for   {
		for a=1;a<500;a++{
			for b=1;b<500;b++{
				for c=1;c<500;c++{
					if a+b+c ==1000 && (a*a + b*b)==c*c{
						d=a*b*c
						f=1
					}
				}
			}
		}
		if f ==1 {
			fmt.Println(d)
			break
		}
	}
}
