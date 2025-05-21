package advanced

import "fmt"

func Functions() {

	/*
			A function is a block of code that performs a specific task.
			A function takes an input, performs some operations on the input and generates outputs.
			The parameters and return type are optional in a function.
		    If consecutive parameters are of the same type, we can avoid writing the type each time and it is enough to be written once at the end.
			Now that we have a function ready, let’s call it from somewhere in the code.
			Multiple return values: It is possible to return multiple values from a function.
			Named return values: It is possible to return named values from a function.
			If a return value is named, it can be considered as being declared as a variable in the first line of the function.
			In the function that have named return values, note that the return statement in the function does not explicitly return any value

			Blank Identifier : _ is known as the blank identifier in Go.

			   func functionname(parametername datatype) returntype {
			         //function body
				}
	*/

	var sm = sum(10, 20)
	fmt.Println(sm)

	width := 20
	length := 30
	area, perimeter := rectProps(float32(length), float32(width))
	fmt.Printf("Area %.02f Perimeter %.02f\n", area, perimeter)

	area, perimeter = rectProps2(float32(length), float32(width))
	fmt.Printf("Area %.02f Perimeter %.02f\n", area, perimeter)

	a, _ := rectProps(20, 20) // / perimeter is discarded
	fmt.Printf("Area %.02f \n", a)
}

func sum(n1, n2 int) int {
	return n1 + n2
}

func rectProps(length, width float32) (float32, float32) {
	area := length * width
	perimeter := (length + width) * 2
	return area, perimeter
}

func rectProps2(length, width float32) (area, perimeter float32) {
	area = length * width
	perimeter = (length + width) * 2
	return
}

func VariadicFunction() {
	/*
		What is a variadic function?
			Functions in general accept only a fixed number of arguments.
			A variadic function is a function that accepts a variable number of arguments.
			If the last parameter of a function definition is prefixed by ellipsis …, then
			the function can accept any number of arguments for that parameter.
			Only the last parameter of a function can be variadic. We will learn why this is the case in the next section of this tutorial.
	*/

}
