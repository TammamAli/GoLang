package helpers

import (
	"fmt"
	"time"
)

func StartTopic(word string) {

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
