package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://yahoo.com")
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	fmt.Println(resp)
}
