package main

import (
	"fmt"
)

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#00ff00",
	}

	// var colors map[string]string

	// colors := make(map[string]string)

	colors["white"] = "#ffffff"
	// delete(colors, "white")
	// fmt.Printf("%+v", colors)

	printMap(colors)

}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key, value)
	}
}
