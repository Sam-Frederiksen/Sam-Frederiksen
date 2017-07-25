package main

import "fmt"

func zero(z *string) {
	*z = "sam"
}

func main() {
	name := "bill"
	zero(&name)
	fmt.Println("Hello", name)

	for i := 1; i < 11; i++ {
		fmt.Println(i)
	}
}
