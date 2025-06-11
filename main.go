package main

import (
	"fmt"
	"goHelper/concurrency"
	"goHelper/simpleinterest"
	"log"
	"os"
	"time"
)

var p, r, t = 5000.0, 10.0, 1.0

var _ = simpleinterest.Calculate

func main() {

	// the first argument is always program name
	programName := os.Args[0]
	fmt.Println(programName)

	//fmt.Println("First app.")

	// goHelper.Variables()
	//
	//	goHelper.Variables2()

	// goHelper.DataConvert()
	//goHelper.Test()

	// basics.Constants()
	// advanced.Functions()

	// goHelper.Output()
	// goHelper.OutputFormat()
	// goHelper.DataTypes()

	// goHelper.Arrays()
	// goHelper.Slices()

	//goHelper.Conditions()

	//advanced.Packages()
	// basics.Loops()
	//basics.Switch()

	// data_structure.ArraysSlices1()
	// data_structure.ArraysSlices2()

	// advanced.VariadicFunction()

	// advanced.Maps()
	// advanced.Strings()

	// advanced.Pointers()

	// advanced.Structs()
	// advanced.Methods()
	// advanced.Interfaces()

	// advanced.Explain()

	// interfaces.StartInterface()
	// interfaces.MulitIntrface()
	//concurrency.Concurrency()

	// //concurrency.Concurrency()
	// go concurrency.Hello()
	// time.Sleep(1 * time.Second)
	// fmt.Println("main function")

	// go concurrency.Numbers()
	// go concurrency.Alphabets()
	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("main terminated")

	// concurrency.Channels()
	fmt.Printf("\n")

	// concurrency.CallChannelS1()
	//	concurrency.CallChannelS2()

	//This program will print the sum of the squares and cubes of the individual digits of a number.

	number := 589
	// fmt.Printf("%v\n", concurrency.CalcSquares(number))
	// fmt.Printf("%v\n", concurrency.CalCubes(number))

	sqrch := make(chan int)
	cubech := make(chan int)
	go concurrency.Chan_CalcSquares(number, sqrch)
	go concurrency.Chan_CalCubes(number, cubech)
	squares, cubes := <-sqrch, <-cubech
	fmt.Println("Final output", squares+cubes)

	// ch := make(chan int)
	// ch <- 6

	concurrency.CallSendOnlyChan()

	//chanelExample()

	for i := range 10 {
		fmt.Printf("%v ", i)
	}

	fmt.Printf("\n\n")

	// concurrency.Receiver()
	concurrency.ReceiverForRange()

	//number := 589
	concurrency.TestCall(589)

	fmt.Printf("\n\n")

	concurrency.BufferedChannels()

	concurrency.ClosedChan()
	concurrency.ChanCapLen()

}

func Chnl(ch chan int) {
	ch <- 1
}

// producer creates and sends values to the channel
func producer(ch chan<- int) {
	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}
	close(ch) // Close the channel when done sending
}

// consumer receives values from a receive-only channel
func consumer(ch <-chan int) {
	for {
		val, ok := <-ch
		if !ok {
			fmt.Println("Channel closed!")
			return
		}
		fmt.Println("Received:", val)
	}
}

func chanelExample() {
	ch := make(chan int)

	// Start producer (send-only)
	go producer(ch)

	// Start consumer (receive-only)
	go consumer(ch)

	// Wait for goroutines to finish
	time.Sleep(6 * time.Second)
	fmt.Println("Done!")
}

func DateTimeLayout() {
	currentTime := time.Now()
	fmt.Println("Current Time in String: ", currentTime.String())
	fmt.Println("MM-DD-YYYY : ", currentTime.Format("01-02-2006"))
	fmt.Println("YYYY-MM-DD : ", currentTime.Format("2006-01-02"))
	fmt.Println("YYYY.MM.DD : ", currentTime.Format("2006.01.02 15:04:05"))
	fmt.Println("YYYY#MM#DD {Special Character} : ", currentTime.Format("2006#01#02"))
	fmt.Println("YYYY-MM-DD hh:mm:ss : ", currentTime.Format("2006-01-02 15:04:05"))
	fmt.Println("Time with MicroSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000"))
	fmt.Println("Time with NanoSeconds: ", currentTime.Format("2006-01-02 15:04:05.000000000"))
	fmt.Println("ShortNum Month : ", currentTime.Format("2006-1-02"))
	fmt.Println("LongMonth : ", currentTime.Format("2006-January-02"))
	fmt.Println("ShortMonth : ", currentTime.Format("2006-Jan-02"))
	fmt.Println("ShortYear : ", currentTime.Format("06-Jan-02"))
	fmt.Println("LongWeekDay : ", currentTime.Format("2006-01-02 15:04:05 Monday"))
	fmt.Println("ShortWeek Day : ", currentTime.Format("2006-01-02 Mon"))
	fmt.Println("ShortDay : ", currentTime.Format("Mon 2006-01-2"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 PM"))
	fmt.Println("Short Hour Minute Second: ", currentTime.Format("2006-01-02 3:4:5 pm"))
}

func init() {

	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

/*
	Go also known as Golang is an open source, compiled and statically typed programming language developed by Google.
	The key people behind the creation of Go are Rob Pike, Ken Thompson and Robert Griesemer.

	Advantages of Go
		1- Simple syntax :
			The syntax is simple and concise and the language is not bloated with unnecessary features.
			This makes it easy to write code that is readable and maintainable.

		2- Easy to write concurrent programs
			Concurrency is an inherent part of the language. As a result, writing multithreaded programs is a piece of cake.
			This is achieved by Goroutines and channels which we will discuss in the upcoming tutorials.

		3. Compiled language
			Go is a compiled language. The source code is compiled to a native binary.
			This is missing in interpreted languages such as JavaScript used in nodejs.

		4.Fast compilation
			The Go compiler is pretty amazing and it has been designed to be fast right from the beginning.

		5. Static linking
			The Go compiler supports static linking. The entire Go project can be statically linked into one big
			fat binary and it can be deployed in cloud servers easily without worrying about dependencies.

		6.Go Tooling
			Tooling deserves a special mention in Go. Go comes bundled with powerful tools that help developers write
			better code. Few of the commonly used tools are,
				i.   gofmt - gofmt is used to automatically format go source code. It uses tabs for indentation and blanks for alignment.
				ii.  vet - vet analyses the go source code and reports possible suspicious code. Everything reported by vet is not a genuine problem but it has the capability to catch errors that are not reported by the compiler such as incorrect format specifiers when using Printf.
				iii. staticcheck - staticcheck is used to enforce styling issues in the code.

		7. Garbage collection
		Go uses garbage collection and hence memory management is pretty much taken care automatically and
		the developer doesn’t need to worry about managing memory. This also helps to write concurrent programs easily.


*/
//-------------------------------------------------------------------------------------------------------

/*  Go Syntax : A Go file consists of the following parts:
Package declaration
Import packages
Functions
Statements and expressions	*/

/*
	A short explanation of the hello world program
	package main
		-Every go file must start with the package name statement.
		-Packages are used to provide code compartmentalization and reusability.
		-The package name main is used here. The main function should always reside in the main package.
*/
// Note: In Go, any executable code belongs to the main package.
// Hitting the Enter key adds ";" to the end of the line implicitly (does not show up in the source code).
// The left curly bracket { cannot come at the start of a line.
// Error  unexpected semicolon or newline before {
//Tip It is up to you which you want to use. Normally, we use // for short comments, and /* */ for longer comments.

/*
The following are the changes made to main.go

p, r and t variables are moved to package level from the main function level.
An init function has been added. The init function prints a log and terminates the program execution if either the principal, rate of interest or time duration is less than zero using log.Fatal function.
The order of initialisation of the is as follows,

The imported packages are initialized first. Hence simpleinterest package is initialized first and it's init method is called.
Package level variables p, r and t are initialized next.
init function is called in main package
main function is called at last.
*/
