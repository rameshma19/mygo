package main

import (
	"fmt"
	"io"
	"os"
)

//New update
func main() {
	fn := os.Args[1]
	f, e := os.Open(fn)

	if e != nil {
		fmt.Println(e)
		os.Exit(0)
	}

	io.Copy(os.Stdout, f)

	// b, e := ioutil.ReadFile(fn)

	// if e != nil {
	// 	fmt.Println(string(b))
	// } else {
	// 	fmt.Println(e)
	// }
}
