package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	// argsWithProg := os.Args
	// fmt.Println("argsWithProg", argsWithProg)
	// byteSlice, err := ioutil.ReadFile(argsWithProg[1])

	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	os.Exit(1)
	// }

	// fmt.Print(string(byteSlice[:]))

	//Sample code from the instructor
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	io.Copy(os.Stdout, f)
}
