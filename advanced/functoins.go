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

			Only the last parameter of a function can be variadic.  The way variadic functions work is by converting
			the variable number of arguments to a slice of the type of the variadic parameter.

			Slice arguments vs Variadic arguments
			The following of the advantages of using variadic arguments instead of slices.
				1. There is no need to create a slice during each function call.
				2. We are creating an empty slice just to satisfy the signature of the find function.
				3. Variadic functions is more readable than the once with slices
				4. Append is a variadic function

			Passing a slice to a variadic function
				We are passing a slice to a function that expects a variable number of arguments.
				This will not work, will fail with compilation error.
				Well, it’s pretty straight forward.

			 According to the definition of a variadic function nums ...int means that it will
			 accept a variable number of arguments of type int. So is there a way to pass a slice to a
			 variadic function? The answer is yes.

			There is a syntactic sugar which can be used to pass a slice to a variadic function.
			You have to suffix the slice with ellipsis ... If that is done, the slice is directly
			passed to the function without a new slice being created.

			Gotcha : Just be sure you know what you are doing when you are modifying a slice inside a variadic function.
			if ... is used, the welcome slice itself will be passed as an argument without a new slice being created. Hence welcome will be passed to the change function as argument.



	*/

	testVariadic(1, 2)
	testVariadic(1, 2, 3, 4)
	testVariadic(1)

	find(1, 10, 20, 5, 51, 10)

	nums := []int{89, 90, 95}
	//find(89, nums)     // error
	find(1, nums...)

	names := []string{"ahmed", "ali", "amr"}
	fmt.Println(names)
	changeInVariadic(names...)
	fmt.Println(names)

}

func testVariadic(a int, b ...int) {
	fmt.Println(a, b)

}

func find(num int, nums ...int) {
	isFound := false

	for index, value := range nums {
		if num == value {
			fmt.Printf("This item is found at %v\n", index)
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Printf("This item is not found \n\n")
	}
}

func find2(num int, nums []int) {
	isFound := false

	for index, value := range nums {
		if num == value {
			fmt.Printf("This item is found at %v\n", index)
			isFound = true
			break
		}
	}
	if !isFound {
		fmt.Printf("This item is not found \n\n")
	}

}

func changeInVariadic(names ...string) {
	if len(names) > 0 {
		names[0] = "no-name"
	}

}
