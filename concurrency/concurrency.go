package concurrency

import (
	"fmt"
	"time"
)

func Concurrency() {
	Display("concurrency")
}

func Display(word string) {

	time.Sleep(time.Millisecond * 300)
	fmt.Printf(word)
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 300)
		fmt.Print(".")
	}
	time.Sleep(time.Millisecond * 300)
	fmt.Printf("!\n")
	time.Sleep(time.Millisecond * 300)
}
