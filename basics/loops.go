package basics

import "fmt"

func Loops() {

	/*
		- A loop is used to execute a block of code repeatedly until a condition is satisfied.
		- for is the only loop available in Go.
		- Go doesn’t have while or do while loops which are present in other languages like C.

		for initialisation; condition; post {
			todo
		}

		The initialisation statement will be executed only once. After the loop is initialised, the condition is checked.
		If the condition evaluates to true, the body of the loop inside the { } will be executed followed by the post statement.
		The post statement will be executed after each successful iteration of the loop.
		After the post statement is executed, the condition will be rechecked.
		If it’s true, the loop will continue executing, else the for loop terminates.
		All the three components namely initialisation, condition and post are optional in Go.
		The variables declared in a for loop are only available within the scope of the loop.
		Hence i cannot be accessed outside the body of the for loop.

		The break statement
			is used to terminate the for loop abruptly before it finishes its normal execution and
			move the control to the line of code just after the for loop.

		The continue statement
			is used to skip the current iteration of the for loop. All code present in a for loop after the continue statement
			will not be executed for the current iteration. The loop will move on to the next iteration.

		Nested for loops
			A for loop which has another for loop inside it is called a nested for loop.

		Labels
			Labels can be used to break the outer loop from inside the inner for loop.
			This is where labels come to our rescue. A label can be used to break from an outer loop.

		Multiple variable declarations
			It is possible to declare and operate on multiple variables in a for loop.

		infinite loop

		There is one more construct range which can be used in for loops for array manipulation.
	*/

	for {
		fmt.Println("Hi")
		break
	}

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nloop ended \n")

	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}

	fmt.Printf("\nloop ended\n ")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}

	fmt.Printf("\nloop ended\n")

	fmt.Println()

	n := 10
	for i := 1; i <= n; i++ {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
	for i := n - 1; i > 0; i-- {
		for j := 1; j <= n-i; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}

		fmt.Printf("\n\n")

	outer:
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
				if i == j {
					break outer
				}
			}
		}
	}

	fmt.Printf("The line after the end of looop\n\n")

	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present. This is similar to while loop.
		fmt.Printf("%d ", i)
		i += 2
	}

	fmt.Printf("\n\n\n")
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i+1, no+1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
}
