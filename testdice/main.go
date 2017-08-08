package main


import (
	"fmt"
	"net/http"
)


func main() {
fmt.Print("Welcome to Casino Twist")
http.ListenAndServe(:8080,nil)
}