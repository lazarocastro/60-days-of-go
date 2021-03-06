package main

import (
	"errors"
	"fmt"
	"log"
	"time"
)

func main() {

	// timeout using select
	ch := make(chan string)
	go func() {
		time.Sleep(time.Second * 2)
		ch <- "never happen"
	}()

	// wait for ch returns or 1 second
	select {
	case res := <-ch:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout occurs")
	}

	// three channels
	// one for numbers, one for errors and one to indicate that computation are done
	done := make(chan bool)
	numbers := make(chan int)
	err := make(chan error)

	// goroutine that uses our three channels
	// until bug the gourotine runs without problem
	go func(bug int) {
		for i := 0; i < 10; i++ {
			if i == bug {
				err <- errors.New("Some error")
			}
			numbers <- i
		}
		// only executed if you changed the bug number
		done <- true
	}(7)

	for {
		select {
		// receive number and print
		case n := <-numbers:
			fmt.Println(n)
		// if an error occurs, stop and log the error
		case e := <-err:
			log.Fatal(e)
		// when done, returns to avoid deadlock on loop
		case <-done:
			return
		}
	}
}
