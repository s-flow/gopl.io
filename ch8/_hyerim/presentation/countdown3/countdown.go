package main

import (
	"fmt"
	"os"
	"time"
)

func main() {

	abort := make(chan struct{})
	go func() {
		os.Stdin.Read(make([]byte, 1))
		abort <- struct{}{}
	}()

	fmt.Println("Commencing countdown. Press return to abort.")

	tick := time.Tick(5 * time.Second) // select를 이용한 다중화를 확인하기 위해 일부러 천천히 카운트다운함
	for countdown := 10; countdown > 0; countdown-- {
		fmt.Println(countdown)

		select {
		case <-tick:
			// none
		case <-abort:
			fmt.Println("Launch aborted!")
			return
		}
	}

	launch()
}

func launch() {
	fmt.Println("Lift off!")
}
