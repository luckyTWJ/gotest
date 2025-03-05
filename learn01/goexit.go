package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		defer fmt.Println("A.defer...\n")
		func() {
			defer fmt.Println("B.defer...\n")

			fmt.Printf("B...\n")
			//退出当前goroutine
			//runtime.Goexit()
		}()

		fmt.Printf("A...\n")
	}()

	go func(a int, b int) bool {
		fmt.Println("a+b=", a+b)
		return true
	}(10, 20)
	for {
		time.Sleep(1 * time.Second)
	}
}
