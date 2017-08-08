package main

import "fmt"

func main(){
	var (
		i int
		j int
		k int
	)
	for i=1;i<101;i++{
        j=j+i*i
		k=k+i
	}
	fmt.Print(k*k-j)
}
