package main

// import required modules
import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	var roll float32
	rand.Seed(time.Now().UnixNano())
	roll =(rand.Float32())*100
	fmt.Println(roll)
}
