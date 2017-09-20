package main

import (
	"fmt"
	"net/http"
)
type hotdog int

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
func  (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln( w,"Hello World")
}
