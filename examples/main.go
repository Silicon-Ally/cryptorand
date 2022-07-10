package main

import (
	"fmt"

	"github.com/Silicon-Ally/cryptorand"
)

func main() {
	r := cryptorand.New()
	// Will print a number [0,9]
	fmt.Println(r.Intn(10))
}
