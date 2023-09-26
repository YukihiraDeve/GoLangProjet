package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
			time.Sleep(time.Millisecond * 100)

		}
		done <- true
	}()

	go func() {
		for i := 0; i < 26; i++ {
			fmt.Printf("%c\n", 'a'+i)
			time.Sleep(time.Millisecond * 200)

		}
		done <- true
	}()
	<-done
	<-done

}
