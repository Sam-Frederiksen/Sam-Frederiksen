package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", CodeServer)
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func CodeServer(w http.ResponseWriter, req *http.Request) {
	stream, err := ioutil.ReadFile("Addressbook.txt")
	if err != nil {
		log.Fatal(err)
	}
	readString := string(stream)
	if err != nil {
		log.Fatal(err)
	}
	//io.WriteString(w,"<body><p>")
	io.WriteString(w, readString)
	//io.WriteString(w,"</p></body>")
}
