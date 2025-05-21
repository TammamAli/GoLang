package data_structure

import (
	"fmt"
	"reflect"
)

func Slices() {

	/*
			lices are similar to arrays, but are more powerful and flexible.
			Like arrays, slices are also used to store multiple values of the same type in a single variable.
		    However, unlike arrays, the length of a slice can grow and shrink as you see fit.

			In Go, there are several ways to create a slice:
				Using the []datatype{values} format
				Create a slice from an array
				Using the make() function

			len() function - returns the length of the slice (the number of elements in the slice)
			cap() function - returns the capacity of the slice (the number of elements the slice can grow or shrink to)

			Create a Slice From an Array
				var myarray = [length]datatype{values} // An array
				myslice := myarray[start:end] // A slice made from the array

			Create a Slice With The make() Function
				The make() function can also be used to create a slice.

			Note: If the capacity parameter is not defined, it will be equal to length.
			You can access a specific slice element by referring to the index number.
			In Go, indexes start at 0. That means that [0] is the first element, [1] is the second element, etc.
			You can also change a specific slice element by referring to the index number.

			Append Elements To a Slice
				You can append elements to the end of a slice using the append()function:
				slice_name = append(slice_name, element1, element2, ...)

			Memory Efficiency
			 	When using slices, Go loads all the underlying elements into the memory.
				If the array is large and you need only a few elements, it is better to copy those elements
				using the copy() function.
				The copy() function creates a new underlying array with only the required elements for the slice.
				This will reduce the memory used for the program.

			The copy() function takes in two slices dest and src, and copies data from src to dest.
			It returns the number of elements copied.

	*/

	var arr1 = [...]int{1, 3, 4}
	var slice1 = []int{}

	fmt.Printf("The type of %v is %T \n", arr1, arr1)
	fmt.Printf("The type of %v is %T \n", slice1, slice1)

	fmt.Printf("The type of %v is %v \n", arr1, reflect.TypeOf(arr1))
	fmt.Printf("The type of %v is %v \n", slice1, reflect.TypeOf(slice1))

	fmt.Printf("The type of %v is %v \n", arr1, reflect.ValueOf(arr1).Kind())
	fmt.Printf("The type of %v is %v \n", slice1, reflect.ValueOf(slice1).Kind())

	/// common ways to create slices

	s1 := []int{} // declares an empty slice of 0 length and 0 capacity.
	fmt.Printf("%v\n", s1)

	var s2 = []string{"S"}
	fmt.Printf("%v\n", s2)

	a1 := [1]int{}
	fmt.Printf("%v\n", a1)

	var a2 = [2]string{"S"}
	fmt.Printf("%v\n", a2)

	myslice1 := []int{}
	fmt.Println(len(myslice1))
	fmt.Println(cap(myslice1))
	fmt.Println(myslice1)

	myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	fmt.Println(len(myslice2))
	fmt.Println(cap(myslice2))
	fmt.Println(myslice2)

	numbers := [18]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice2 := numbers[0:3] // A slice made from the array. The slice can grow to the end of the array.
	fmt.Println(slice2)
	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

	slice3 := numbers[5:7] // The slice can grow to the end of the array.
	fmt.Println(slice2)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 := make([]int, 5)
	fmt.Printf("%v , len: %v, cap: %v\n", slice4, len(slice4), cap(slice4))

	slice5 := make([]int, 4, 60)
	fmt.Printf("%v , len: %v, cap: %v\n", slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, 1, 3, 4)
	fmt.Printf("%v , len: %v, cap: %v\n", slice5, len(slice5), cap(slice5))

	slice5 = append(slice5, slice2...)
	fmt.Printf("%v , len: %v, cap: %v\n", slice5, len(slice5), cap(slice5))

	numbersArr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Printf("numbers = %v, length= %d, capacity= %d\n", numbersArr, len(numbersArr), cap(numbersArr))

	neededNumbers := numbersArr[:len(numbersArr)-10]
	fmt.Printf("numbers = %v, length= %d, capacity= %d\n", neededNumbers, len(neededNumbers), cap(neededNumbers))

	numbersCopy := make([]int, len(neededNumbers))
	fmt.Printf("numbers = %v, length= %d, capacity= %d\n", numbersCopy, len(numbersCopy), cap(numbersCopy))

	result := copy(numbersCopy, neededNumbers)
	fmt.Printf("numbers = %v, length= %d, capacity= %d, result = %d\n", numbersCopy, len(numbersCopy), cap(numbersCopy), result)

}
