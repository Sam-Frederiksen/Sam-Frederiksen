package main

import (
	"fmt"
//	"io/ioutil"
//	"log"
//	"net/http"
	"net/http"
	"log"
	"io/ioutil"
)

func main() {
	var a string
	fmt.Print("Enter Url include http:// you want to see code for ")
	fmt.Scan(&a)
	res, err := http.Get(a)
	if err != nil {
		log.Fatalln(err)
	}

	bs, _ := ioutil.ReadAll(res.Body)
	str := string(bs)
	fmt.Println(str)
	fmt.Print("Press Enter to Quit")
	fmt.Scan(&a)
}
