package main

import (
	"goHelper/data_structure"
	"goHelper/simpleinterest"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

var _ = simpleinterest.Calculate

func main() {

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
	data_structure.ArraysSlices2()
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
