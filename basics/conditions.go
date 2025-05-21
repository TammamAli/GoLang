package basics

import "fmt"

func Conditions() {

	/*
		if statement has a condition and it executes a block of code if that condition evaluates to true.
		It executes an alternate else block if the condition evaluates to false.
		Unlike in other languages like C, the braces { } are not optional and they are mandatory even
		if there is only one line of code between the braces{ }.
		The condition is evaluated for the truth from the top to bottom.
		In general, whichever if or else if’s condition evaluates to true, it’s corresponding code block is executed.
		If none of the conditions are true then else block is executed.




		Conditional statements are used to perform different actions based on different conditions.
		A condition can be either true or false.
		Go supports the usual comparison operators from mathematics:

		Go has the following conditional statements:
			Use if to specify a block of code to be executed, if a specified condition is true
			Use else to specify a block of code to be executed, if the same condition is false
			Use else if to specify a new condition to test, if the first condition is false
			Use switch to specify many alternative blocks of code to be executed
	*/

	fmt.Println("Go Conditions")

	time := 20
	if time < 18 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	temperature := 14
	if temperature > 15 {
		fmt.Println("It is warm out there")
	} else {
		fmt.Println("It is cold out there")
	}

	i := 5
	if i < 10 {
		fmt.Println(i)
	}

	// 	  if (temperature > 15) {
	//     fmt.Println("It is warm out there.")
	//   } // this raises an error
	//   else {
	//     fmt.Println("It is cold out there.")
	//   }

	if time < 10 {
		fmt.Println("Good morning.")
	} else if time < 20 {
		fmt.Println("Good day.")
	} else {
		fmt.Println("Good evening.")
	}

	/*
		The switch Statement: Use the switch statement to select one of many code blocks to be executed.
		The switch statement in Go is similar to the ones in C, C++, Java, JavaScript, and PHP.
		The difference is that it only runs the matched case so it does not need a break statement.
		All the case values should have the same type as the switch expression. Otherwise, the compiler will raise an error:

		The assignment-statement is first executed before the condition is evaluated.
		i.e. the scope of variable is limited to the if, else if and else blocks.
		If we try to access variable outside the if, else if or else blocks, the compiler will complain.

			The else statement should start in the same line after the closing curly brace } of the if statement.
		If not the compiler will complain because of the way Go inserts semicolons automatically.


		Single-Case switch Syntax
		Syntax
		switch expression {
			case x:
			// code block
			case y:
			// code block
			case z:
			...
			default:
			// code block
			}

		This is how it works:
			The expression is evaluated once
			The value of the switch expression is compared with the values of each case
			If there is a match, the associated block of code is executed
			The default keyword is optional. It specifies some code to run if there is no case match

		The Multi-case switch Statement
		It is possible to have multiple values for each case in the switch statement:

		switch expression {
		case x,y:
		// code block if expression is evaluated to x or y
		case v,w:
		// code block if expression is evaluated to v or w
		case z:
		...
		default:
		// code block if expression is not found in any cases
		}
	*/

	day := 4
	switch day {
	case 1:
		fmt.Println("Saturday")
	case 2:
		fmt.Println("Sunday")
	case 3:
		fmt.Println("Monday")
	case 4:
		fmt.Println("Tuesday")
	case 5:
		fmt.Println("Wedensday")
	case 6:
		fmt.Println("Thursday")
	case 7:
		fmt.Println("Friday")
	default:
		fmt.Println("Not a weekday")
	}

	switch day {
	case 1, 3, 5:
		fmt.Println("Odd weekdey")
	case 2, 4:
		fmt.Println("Even weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid day of day number")
	}

	if true {
		fmt.Println("Ssdfsdfs")
	}

	if ft := 20; ft < 100 {
		fmt.Println(ft)
	}

}

func GoPhilosophy() {

	// In Go’s philosophy, it is better to avoid unnecessary branches and indentation of code.
	// It is also considered better to return as early as possible.

	num := 10
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}

	// This is better than previous code, this avoids the unnecessary else code branch.
	// This is the way things are done in Go. Please keep this in mind whenever writing Go programs.
	if num%2 == 0 { //checks if number is even
		fmt.Println(num, "is even")
		return
	}
	fmt.Println(num, "is odd")

}
