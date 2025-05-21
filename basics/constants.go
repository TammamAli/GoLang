package basics

import (
	"fmt"
	"math"
)

func Constants() {
	/*
			Go Constants

			Constants are generally used to represent values that do not change throughout the life time of an application.
			Constants, as the name indicate, cannot be reassigned again to any other value.
			The value of a constant should be known at compile time.
			Hence it cannot be assigned to a value returned by a function call since the function call takes place at run time.
			String Constants, Typed and Untyped Constants
			Go is a strongly typed language. All variables require an explicit type.
			Go is a strongly typed language. Mixing types during the assignment is not allowed.
			Untyped constants have a default type associated with them and they supply it
			  if and only if a line of code demands it
			Numeric constants are free to be mixed and matched in expressions and a type is needed only when they are
			assigned to variables or used in any place in code which demands a type.

			If a variable should have a fixed value that cannot be changed, you can use the const keyword.
			The const keyword declares the variable as "constant", which means that it is unchangeable and read-only.
		    Note: The value of a constant must be assigned when you declare it.
			Syntax
				const CONSTNAME type = value

			Constant Rules
				Constant names follow the same naming rules as variables
				Constant names are usually written in uppercase letters (for easy identification and differentiation from variables)
				Constants can be declared both inside and outside of a function

			Constant Types : There are two types of constants:
				Typed constants  : Typed constants are declared with a defined type:
				Untyped constants : Untyped constants are declared without a type:

			Multiple Constants Declaration : Multiple constants can be grouped together into a block for readability:

	*/
	const PI = 3.14
	const RATE int = 3

	fmt.Printf("Typed constants RATE %v\n", RATE)
	fmt.Printf("Untyped constants PI %v\n", PI)

	const untypedInt = 30
	const untypedFloat = 12.5
	const untypedString = "untyped constant"

	const typedInt int = 10
	const typedFloat float32 = 5.02
	const typedString string = "typed string"

	fmt.Println(untypedFloat, untypedInt, untypedString, typedFloat, typedInt, typedString)

	var sqrt = math.Sqrt(9)
	fmt.Println(sqrt)

	const n = "Sam"
	var name = n
	fmt.Printf("type %T value %v", name, name)

	//Boolean constants have two untyped constants true and false.

	const c = 5
	var intVar int = c
	var int32Var int32 = c
	var float64Var float64 = c
	var complex64Var complex64 = c
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	/*
		In the program above, the value of c is 5 and the syntax of c is generic.
		It can represent a float, integer or even a complex number with no imaginary part.
		The default type of these kinds of constants can be thought of as being generated based on the context where they are used.
		var intVar int = c requires c to be int so it becomes an int constant.
		var complex64Var complex64 = c requires c to be a complex number and hence it becomes a complex constant.
		Pretty neat :).

	*/

	//Numeric Expressions

	var c1 = 9.5 / 8
	fmt.Println(c1)

	c2 := 10
	c3 := 20.5
	//c4 := c3 / c2  //invalid operation: c3 / c2 (mismatched types float64 and int)compiler
	c4 := c3 / float64(c2)
	fmt.Println(c2, c3, c4)

}
