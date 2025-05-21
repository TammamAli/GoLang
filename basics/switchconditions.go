package basics

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Switch() {
	/*
		What is a switch statement?
		A switch is a conditional statement that evaluates an expression and compares it against a list of possible matches and
		executes the corresponding block of code. It can be considered as an idiomatic way of replacing complex if else clauses.

		Duplicate cases are not allowed: Duplicate cases with the same constant value are not allowed.
		Multiple expressions in case : It is possible to include multiple expressions in a case by separating them with comma.

		Expressionless switch
		The expression in a switch is optional and it can be omitted. If the expression is omitted, the switch is considered
		to be switch true and each of the case expression is evaluated for truth and the corresponding block of code is executed.
		This type of switch can be considered as an alternative to multiple if else clauses.

		Fallthrough
		In Go, the control comes out of the switch statement immediately after a case is executed.
		A fallthrough statement is used to transfer control to the first statement of the case that is present immediately
		after the case which has been executed.
		Fallthrough will happen even though the case evaluates to false.
		So be sure that you understand what you are doing when using fallthrough.

		Switch and case expressions need not be only constants. They can be evaluated at runtime too.
		cannot fallthrough final case in switchcompiler.

		Breaking switch => The break statement can be used to terminate a switch early before it completes.

		Breaking the outer for loop
		  When the switch case is inside a for loop, there might be a need to terminate the for loop early.
		  This can be done by labeling the for loop and breaking the for loop using that label inside the switch statement.
		Please note that if the break statement is used without the label, the switch statement will only be broken and
		the loop will continue running. So labeling the loop and using it in the break statement inside the switch
		is necessary to break the outer for loop.

		There is one more type of switch called type switch. We will look into this when we learn about interfaces.





	*/

	finger := 10
	fmt.Printf("Finger %d is ", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger numbber")
	}

	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Printf("%s is a vowel", letter)
	default:
		fmt.Printf("%s is not a vowel", letter)
	}

	hour := 15 // hour in 24 hour format
	// Using switch to determine the work shift
	switch {
	case hour >= 6 && hour < 12:
		fmt.Println("It's the morning shift.")
	case hour >= 12 && hour < 17:
		fmt.Println("It's the afternoon shift.")
	case hour >= 17 && hour < 21:
		fmt.Println("It's the evening shift.")
	case (hour >= 21 && hour <= 24) || (hour >= 0 && hour < 6):
		fmt.Println("It's the night shift.")
	default:
		fmt.Println("Invalid hour.")
	}

	var input string
	fmt.Print("Enter another input: ")
	fmt.Scanln(&input)
	switch num, _ := strconv.Atoi(input); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200\n", num)
		//fallthrough
	}

forloop:
	for {
		switch rnd := rand.Intn(333); {
		case (rnd)%2 == 0:
			fmt.Printf("Generated even number %d\n", rnd)
			time.Sleep(300 * time.Millisecond)
			break forloop
		}
	}
}
