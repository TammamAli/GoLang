package advanced

import "fmt"

func Packages() {
	/*
		So far we have seen Go programs that have only one file with a main function and a couple of other functions.
		In real-world scenarios, this approach of writing all source code in a single file is not scalable.
		It becomes impossible to reuse and maintain code written this way. This is where packages are helpful.

		Packages are used to organize Go source code for better reusability and readability.
		Packages are a collection of Go sources files that reside in the same directory.
		Packages provide code compartmentalization and hence it becomes easy to maintain Go projects.

		Go Modules are needed to create custom packages.
		Go Module is nothing but a collection of Go packages.
		Why do we need Go modules to create a custom package?
		The answer is the import path for the custom package we create is derived from the name of the go module.
		In addition to this, all the other third-party packages(such as source code from github) along with their versions
		which our application uses will be managed by the go.mod file. This go.mod file is created when we create a new module.


		1- Creating a Go module
		2- Create the simple interest custom package:
		  All Go files that belong to a package should be placed in separate folders of their own.
		  It is a convention in Go to name this folder with the same name as the package.

		The main package and the main function
		Every executable Go application must contain the main function.
		This function is the entry point for execution. The main function should reside in the main package.


		A bit more on go build
			Go tools like go build work in the context of the current directory.
			Let’s understand what that means. Till now we have been running go build from the directory ~/Documents/learnpackage/.
			If we try to run it from any other directory, it will fail.
			Let’s understand the reason behind this error.
			go build takes an optional package name as a parameter(in our case the package name is learnpackage)
			and it tries to compile the main function if the package exists in the current directory from which it
			is run or in the parent directory or it’s parent directory and so on.

			Go build has the ability to recursively search the parent directory for a go.mod file.

		Exported Names
			We capitalized the function Calculate in the Simple interest package.
			This has a special meaning in Go. Any variable or function which starts with a capital letter are exported names in go.
			Only exported functions and variables can be accessed from other packages.

		init function
			1- Each package in Go can contain an init function.
			2- The init function must not have any return type and it must not have any parameters.
			3- The init function cannot be called explicitly in our source code.
			4- It will be called automatically when the package is initialized. The init function has the following syntax.

			The init function can be used to perform initialization tasks and can also be used to verify the correctness of the
			program before the execution starts.

			The order of initialization of a package is as follows
				- Package level variables are initialised first
				- init function is called next.
				- A package can have multiple init functions (either in a single file or distributed across multiple files)
				 and they are called in the order in which they are presented to the compiler.
				- If a package imports other packages, the imported packages are initialized first.
				- A package will be initialized only once even if it is imported from multiple packages.

		p, r and t variables are moved to package level from the main function level.
		An init function has been added. The init function prints a log and terminates the program execution if either the principal,
		rate of interest or time duration is less than zero using log.Fatal function.
		The order of initialisation of the is as follows,

		The imported packages are initialized first.
		Hence simpleinterest package is initialized first and it's init method is called.
		Package level variables p, r and t are initialized next.
		init function is called in main package main function is called at last.

		Use of blank identifier
		It is illegal in Go to import a package and not to use it anywhere in the code.
		The compiler will complain if you do so. The reason for this is to avoid bloating
		of unused packages which will significantly increase the compilation time



	*/

	//init()  //undefined: initcompilerUndeclaredName

}

func init() {
	fmt.Println("Call automatically init function")
}

func init() {
	fmt.Println("second init function")
}
