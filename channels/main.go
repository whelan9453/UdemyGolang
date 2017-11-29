package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	watchList := []string{
		"https://google.com",
		"https://facebook.com",
		"https://golang.org",
		"https://stackoverflow.com",
	}

	c := make(chan string)

	for _, link := range watchList {
		go checkLinkStatus(link, c)
	}

	// for i := 0; i < len(watchList); i++ {
	// 	fmt.Println(<-c)
	// }

	// for {
	// 	go checkLinkStatus(<-c, c)
	// }

	// for link := range c {
	// 	go checkLinkStatus(link, c)
	// }

	//GOTCHA: One never simply reference a value in different go rutines
	// for link := range c {
	// 	go func() {
	// 		time.Sleep(time.Second * 5)
	// 		checkLinkStatus(link, c)
	// 	}()
	// }

	for link := range c {
		go func(l string) {
			time.Sleep(time.Second * 5)
			checkLinkStatus(l, c)
		}(link)
	}
}

func checkLinkStatus(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is okay!")
	c <- link
}
