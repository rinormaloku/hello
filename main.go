package main

import "time"

func main() {
	for {
		println("Hello 3")
		time.Sleep(500 * time.Millisecond)
	}

}
