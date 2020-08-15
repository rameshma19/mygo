package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	mySites := []string{"http://google.com",
		"http://yahoo.com",
		"http://facebook.com",
	}

	c := make(chan string)

	for _, link := range mySites {
		go checkSite(link, c)
	}

	for l := range c {
		go func(l2 string) {
			time.Sleep(5 * time.Second)
			checkSite(l2, c)
		}(l)
	}

	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
}

func checkSite(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is down!!")
		c <- site
		os.Exit(0)
	}
	fmt.Println(site, "is Up!!")
	c <- site
}
