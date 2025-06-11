package concurrency

import (
	"fmt"
	"time"
)

func BufferedChannels() {
	/*
		   All the channels we discussed in the previous tutorial were basically unbuffered.
		   It is possible to create a channel with a buffer. Sends to a buffered channel are blocked only when the buffer is full.
		   Similarly receives from a buffered channel are blocked only when the buffer is empty.

		   Buffered channels can be created by passing an additional capacity parameter to the make function which
		   specifies the size of the buffer.         	   ch := make(chan type, capacity)
		   capacity in the above syntax should be greater than 0 for a channel to have a buffer.
		   The capacity for an unbuffered channel is 0 by default and hence we omitted the capacity parameter while creating channels

		   Deadlock
		   we write 3 strings to a buffered channel of capacity 2. When the control reaches the third write,
		   the write is blocked since the channel has exceeded its capacity.
		   Now some Goroutine must read from the channel in order for the write to proceed, but in this case,
		   there is no concurrent routine reading from this channel. Hence there will be a deadlock and the
		   program will panic at run time with the following message,



		   Closing buffered channels
		    . There is one more subtlety to be considered when closing buffered channels.
		    . It’s possible to read data from a already closed buffered channel. The channel will return the data that is
		      already written to the channel and once all the data has been read, it will return the zero value of the channel.
		    . Even though the channel is closed, we can read the values already written to the channel.


		  Length vs Capacity
		    . The capacity of a buffered channel is the number of values that the channel can hold.
			This is the value we specify when creating the buffered channel using the make function.
		    . The length of the buffered channel is the number of elements currently queued in it.











	*/

	ch := make(chan string, 2)
	ch <- "naveen"
	ch <- "paul"
	//ch <- "steve" // deadlock   wirte more capacity
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//fmt.Println(<-ch)  // deadlock   read more capacity

	// unbufferedChan := make(chan int)
	// unbufferedChan <- 1
	// fmt.Println(<-unbufferedChan)

	intChan := make(chan int, 2)

	go ChanWriter(intChan, 15)
	time.Sleep(300 * time.Millisecond)
	for v := range intChan {
		fmt.Printf("Read value %v from ch\n", v)
		time.Sleep(300 * time.Millisecond)
	}

}

func ChanWriter(ch chan int, num int) {
	for v := range num {
		val := v + 1
		ch <- val
		fmt.Println("successfully wrote", val, "to ch")
	}
	close(ch)
}

func ClosedChan() {
	ch := make(chan int, 5)
	ch <- 5
	ch <- 6
	close(ch)

	flag := true
	if flag {
		// ch <- 1  // Panic : send on closed channel
		n, open := <-ch
		fmt.Printf("Received: %d, open: %t\n", n, open)
		n, open = <-ch
		fmt.Printf("Received: %d, open: %t\n", n, open)
		n, open = <-ch //fatal error: all goroutines are asleep - deadlock! where Ch is not closed
		fmt.Printf("Received: %d, open: %t\n", n, open)
	} else {
		for n := range ch {
			fmt.Println("Received:", n)
		}
	}

}

func ChanCapLen() {
	ch := make(chan int, 3)
	ch <- 100
	ch <- 200
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))

	chVal := <-ch
	fmt.Println("read from ch : ", chVal)
	fmt.Println("capacity is", cap(ch))
	fmt.Println("length is", len(ch))

}
