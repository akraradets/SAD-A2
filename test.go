package main

import (
	"github.com/SAD-A2/singleton"
	"fmt"
)

func main() {
	s := singleton.New()
	
	s2 := singleton.New()
	
	fmt.Println("This is ", s)
	fmt.Println("This is ", s2)
	// This is that
}