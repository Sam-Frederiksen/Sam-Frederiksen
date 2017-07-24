package main

import "fmt"

func zero(z *string) {
	fmt.Println(z)
	*z = "sam"
}

func main() {
	name := "bill"
	zero(&name)
	fmt.Println("Hello", name)
}
