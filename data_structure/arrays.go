package data_structure

import "fmt"

func Arrays() {
	// Arrays are used to store multiple values of the same type in a single variable,
	// instead of declaring separate variables for each value.

	/* Are two ways to declare an array:
	1- With the var keyword:
		var array_name = [length]datatype{values} // here length is defined
		or
		var array_name = [...]datatype{values} // here length is inferred
	2. With the := sign:
		array_name := [length]datatype{values} // here length is defined
		or
		array_name := [...]datatype{values} // here length is inferred

	Note: The length specifies the number of elements to store in the array.
	In Go, arrays have a fixed length. The length of the array is either defined by
	a number or is inferred (means that the compiler decides the length of the array, based on the number of values).

	You can access a specific array element by referring to the index number.
	indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.
	You can also change the value of a specific array element by referring to the index number.

	Array Initialization
		If an array or one of its elements has not been initialized in the code, it is assigned the default value of its type.
		Tip: The default value for int is 0, and the default value for string is "".
		It is possible to initialize only specific elements in an array.

	The len() function is used to find the length of an array:
	*/

	var arr1 = [3]int{}
	var arr2 = [3]int{1, 2, 3}
	var arr3 = [...]int{4, 5, 6, 7, 8}
	var arr4 = [...]int{}
	fmt.Printf("type is %T  & value %v\n", arr1, arr1)
	fmt.Printf("type is %T  & value %v\n", arr2, arr2)
	fmt.Printf("type is %T  & value %v\n", arr2, arr3)
	fmt.Printf("type is %T  & value %v\n", arr4, arr4)

	fmt.Println(arr1[0])
	fmt.Println(arr2[0])
	fmt.Println(arr3[0])
	//fmt.Println(arr4[0]) //invalid argument: index 0 out of bounds [0:0]

	arrStr := [3]string{}
	arrStr2 := [3]string{"a", "b", "C"}
	arrStr3 := [...]string{}
	fmt.Printf("type is %T  & value %v\n", arrStr, arrStr)
	fmt.Printf("type is %T  & value %v\n", arrStr2, arrStr2)
	fmt.Printf("type is %T  & value %v\n", arrStr3, arrStr3)

	arr5 := [5]int{1: 10, 2: 40}

	fmt.Println(arr5)
}
