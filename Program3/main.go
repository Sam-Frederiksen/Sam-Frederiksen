package main

import "fmt"

func main() {
	var name string
	fmt.Print("Please Enter Your Name: ")
	fmt.Scan(&name)
	fmt.Println("Hello", name)
	for i := 1; i < 11; i++ {
		fmt.Print(i)
	}
}
