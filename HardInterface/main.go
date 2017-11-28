package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	argsWithProg := os.Args
	// fmt.Println("argsWithProg", argsWithProg)
	byteSlice, err := ioutil.ReadFile(argsWithProg[1])

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Print(string(byteSlice[:]))
}
