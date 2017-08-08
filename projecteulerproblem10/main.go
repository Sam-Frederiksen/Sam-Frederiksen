package main

import "fmt"

func main() {
	var (
		i   int
		j   int
		t   int
		sum int
	)
	for i = 1; i < 2000000; i++ {
		t = 0
		for j = 2; j < i; j++ {
			if i%j == 0 {
				t = 1
				break
			}
		}
		if t != 1 {
			sum = sum + i
			fmt.Println(i)
		}
	}
	fmt.Println(sum)
}
