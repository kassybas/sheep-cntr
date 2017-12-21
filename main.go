package main

import "fmt"
import "time"

var inChan chan string

func countMe(s string, limit int) {
	for i := 0; i < limit; i++ {
		fmt.Printf("%d %s\n", i+1, s)
		time.Sleep(time.Millisecond)
	}
}

func main() {
	inChan = make(chan string)
	go startMe()
	inChan <- "sheep"
	inChan <- "wolf"
	inChan <- "beef"
	time.Sleep(10 * time.Second)
}

func startMe() {
	done := false
	for done != true {
		select {
		case s := <-inChan:
			countMe(s, 1000)
		default:
			done = true
			fmt.Println("all good")
		}
	}
	fmt.Println("We are done")
	inChan = make(chan string)
}
