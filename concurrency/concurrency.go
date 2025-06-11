package concurrency

import (
	"fmt"
	"goHelper/helpers"
	"time"
)

func Concurrency() {

	/*
		Concurrency

		Go is a concurrent language and not a parallel one.

		What is concurrency and how it is different from parallelism.
			. Concurrency is the capability to deal with lots of things at once.
			. Parallelism is doing lots of things at the same time. It might sound similar to concurrency but it’s actually different.
			. Concurrency is an inherent part of the Go programming language. Concurrency is handled in Go using Goroutines and channels.

		What are Goroutines?
			Goroutines are functions or methods that run concurrently with other functions or methods.
			Goroutines can be thought of as lightweight threads. The cost of creating a Goroutine is tiny when compared to a thread.
			Hence it’s common for Go applications to have thousands of Goroutines running concurrently.

		Advantages of Goroutines over threads.
			. Goroutines are extremely cheap when compared to threads. They are only a few kb in stack size and the stack can grow and shrink
			  according to the needs of the application whereas in the case of threads the stack size has to be specified and is fixed.
			. The Goroutines are multiplexed to a fewer number of OS threads. There might be only one thread in a program with thousands
			  of Goroutines. If any Goroutine in that thread blocks say waiting for user input, then another OS thread is created
			  and the remaining Goroutines are moved to the new OS thread. All these are taken care of by the runtime and we as
			  programmers are abstracted from these intricate details and are given a clean API to work with concurrency.
			. Goroutines communicate using channels. Channels by design prevent race conditions from happening when accessing
			  shared memory using Goroutines. Channels can be thought of as a pipe using which Goroutines communicate.

		How to start a Goroutine?
		Prefix the function or method call with the keyword go and you will have a new Goroutine running concurrently.
		The main function runs in its own Goroutine and it’s called the main Goroutine.


		--We need to understand the two main properties of goroutines to understand why this happens.
		  . When a new Goroutine is started, the goroutine call returns immediately. Unlike functions, the control does not wait for
		    the Goroutine to finish executing. The control returns immediately to the next line of code after the Goroutine call
			and any return values from the Goroutine are ignored.
		  . The main Goroutine should be running for any other Goroutines to run. If the main Goroutine terminates then
		     the program will be terminated and no other Goroutine will run.


		go hello()
		fmt.Println("main function")

		I guess now you will be able to understand why our Goroutine did not run. After the call to go hello() in line no.
		11, the control returned immediately to the next line of code without waiting for the hello goroutine to finish
		and printed main function. Then the main Goroutine terminated since there is no other code to execute and hence
		 the hello Goroutine did not get a chance to run.

		This way of using sleep in the main Goroutine to wait for other Goroutines to finish their execution is a hack we are
		using to understand how Goroutines work. Channels can be used to block the main Goroutine until all other Goroutines finish
		their execution. We will discuss channels in the next tutorial.




	*/

	helpers.StartTopic("concurrency")

}

func Hello() {
	time.Sleep(time.Second * 2)
	fmt.Println("Hello world goroutine")
}

func Numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(250 * time.Millisecond)
		fmt.Printf("%d ", i)
	}
}
func Alphabets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(400 * time.Millisecond)
		fmt.Printf("%c ", i)
	}
}
