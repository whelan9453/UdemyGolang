package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	res, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// bs := make([]byte, 1000)
	// res.Body.Read(bs)
	// fmt.Println(string(bs))

	// io.Copy(os.Stdout, res.Body)

	lw := logWriter{}
	io.Copy(lw, res.Body)
}

func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	fmt.Println("Just write this many bytes:", len(bs))
	return len(bs), nil
}
