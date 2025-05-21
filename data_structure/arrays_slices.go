package data_structure

import (
	"fmt"
	"math/rand"
	"reflect"
)

func ArraysSlices1() {

	/*
		Arrays : An array is a collection of elements that belong to the same type.
		Mixing values of different types, for example, an array that contains both strings and integers is not allowed in Go.

		Declaration : An array belongs to type [n]T. n denotes the number of elements in an array and T represents the type of each element.
		The number of elements n is also a part of the type(We will discuss this in more detail shortly.)

		All elements in an array are automatically assigned the zero value of the array type
		The index of an array starts from 0 and ends at length - 1
		It is not necessary that all elements in an array have to be assigned a value during short hand declaration.
		You can even ignore the length of the array in the declaration and replace it with ... and let the compiler find the length for you.
		The size of the array is a part of the type. Hence [5]int and [25]int are distinct types. Because of this, arrays cannot be resized.
		Don’t worry about this restriction since slices exist to overcome this.

		a := [3]int{5, 78, 8}
		var b [5]int
		b = a //not possible since [3]int and [5]int are distinct types

		Arrays are value types
		Arrays in Go are value types and not reference types. This means that when they are assigned to a new variable,
		a copy of the original array is assigned to the new variable.
		If changes are made to the new variable, it will not be reflected in the original array.
		Similarly when arrays are passed to functions as parameters, they are passed by value and the original array is unchanged.

		Length of an array is found by passing the array as parameter to the len function.

		The len built-in function returns the length of v, according to its type:
			Array: the number of elements in v.
			Pointer to array: the number of elements in *v (even if v is nil).
			Slice, or map: the number of elements in v; if v is nil, len(v) is zero.
			String: the number of bytes in v.
			Channel: the number of elements queued (unread) in the channel buffer; if v is nil, len(v) is zero.

		Iterating arrays using range The for loop can be used to iterate over elements of an array.

		Go provides a better and concise way to iterate over an array by using the range form of the for loop.
		range returns both the index and the value at that index.
		In case you want only the value and want to ignore the index, you can do this by replacing the index with the _ blank identifier.
		Similarly, the value can also be ignored.

		Multidimensional arrays,  The arrays we created so far are all single dimension. It is possible to create multidimensional arrays.
		Although arrays seem to be flexible enough, they come with the restriction that they are of fixed length.
		It is not possible to increase the length of an array.



	*/

	fmt.Println("Arrays_Slices")

	var a [3]int
	fmt.Println(a)

	a[0] = 12 // array index starts at 0
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	var b = [3]int{}
	fmt.Println(b)

	c := [3]int{}
	fmt.Println(c)

	a1 := [3]int{12, 78, 50} // short hand declaration to create array
	fmt.Println(a1)

	a2 := [...]int{1, 2, 3, 4}
	fmt.Printf("%v , %v\n\n", len(a2), a2)

	s1 := [...]string{"USA", "China", "India", "Germany", "France"}
	s2 := s1 // a copy of a is assigned to b
	s1[0] = "Singapore"
	fmt.Println("s1 is ", s1)
	fmt.Println("s2 is ", s2)

	numArr := [5]int{10, 20, 30, 40, 50}
	fmt.Printf("array before call %v\n", numArr)
	changeLocal(numArr)
	fmt.Printf("array after call %v\n", numArr)

	for i := 0; i < len(numArr); i++ {
		fmt.Printf("arr<%v> = %v\n", i, numArr[i])
	}
	fmt.Printf("\n\n")
	for index, value := range numArr {
		fmt.Printf("arr<%v> = %v\n", index, value)
	}

	for _, v := range numArr {
		fmt.Println(v)
	}

	arr := [3][2]int{
		{rand.Intn(100), rand.Intn(100)},
		{rand.Intn(100), rand.Intn(100)},
		{rand.Intn(100), rand.Intn(100)}, // this comma is necessary. The compiler will complain if you omit this comma
	}
	fmt.Println(arr)

	for rowIndx, rowArr := range arr {
		for colIndx, val := range rowArr {
			fmt.Printf("arr[%v, %v] = %v\n", rowIndx, colIndx, val)
		}

	}
}

func ArraysSlices2() {
	/*
													Slices
				A slice is a convenient, flexible and powerful wrapper on top of an array.
				Slices do not own any data on their own. They are just references to existing arrays.

				Creating a slice
				A slice with elements of type T is represented by []T

				modifying a slice
				A slice does not own any data of its own. It is just a representation of the underlying array.
				Any modifications done to the slice will be reflected in the underlying array.
				When we print the array after the for loop, we can see that the changes to the slice are reflected in the array.
				When a number of slices share the same underlying array, the changes that each one makes will be reflected in the array.
				numa[:] the start and end values are missing. The default values for start and end are 0 and len(numa) respectively.

				ength and capacity of a slice
				The length of the slice is the number of elements in the slice.
				The capacity of the slice is the number of elements in the underlying array starting from the index from which the slice is created.
				A slice can be re-sliced upto its capacity.

				creating a slice using make:
		        func make([]T, len, cap) []T can be used to create a slice by passing the type, length and capacity.
				The capacity parameter is optional and defaults to the length.
				The make function creates an array and returns a slice reference to it.

				Appending to a slice:
				As we already know arrays are restricted to fixed length and their length cannot be increased.
				Slices are dynamic and new elements can be appended to the slice using append function.
				The definition of append function is func append(s []T, x ...T) []T.
				x …T in the function definition means that the function accepts variable number of arguments for the parameter x.
				These type of functions are called variadic functions.

				One question might be bothering you though.
				If slices are backed by arrays and arrays themselves are of fixed length then how come a slice is of dynamic length.
				Well what happens under the hood is, when new elements are appended to the slice, a new array is created.
				The elements of the existing array are copied to this new array and a new slice reference for this new array is returned.
				The capacity of the new slice is now twice that of the old slice.  Pretty cool right :).

				The zero value of a slice type is nil. A nil slice has length and capacity 0.
				It is possible to append values to a nil slice using the append function.
				It is also possible to append one slice to another using the ... operator.

				Passing a slice to a function
				Slices can be thought of as being represented internally by a structure type.
				A slice contains the length, capacity and a pointer to the zeroth element of the array.
				When a slice is passed to a function, even though it’s passed by value, the pointer variable
				will refer to the same underlying array. Hence when a slice is passed to a function as parameter,
				changes made inside the function are visible outside the function too. Lets write a program to check this out.

				Multidimensional slices
				Similar to arrays, slices can have multiple dimensions.

				Memory Optimisation
				Slices hold a reference to the underlying array.
				As long as the slice is in memory, the array cannot be garbage collected.
				This might be of concern when it comes to memory management.
				Lets assume that we have a very large array and we are interested in processing only a small part of it.
				Henceforth we create a slice from that array and start processing the slice.
				The important thing to be noted here is that the array will still be in memory since the slice references it.

				One way to solve this problem is to use the copy function
					func copy(dst, src []T) int
				to make a copy of that slice. This way we can use the new slice and the original array can be garbage collected.







	*/

	arr := [6]int{80, 30, 48, 98, 12, 55}
	slice := arr[1:4]
	fmt.Printf("%v  %v \n", reflect.ValueOf(arr).Kind(), reflect.ValueOf(slice).Kind())

	c := []int{6, 7, 8} //creates an array and returns a slice reference
	fmt.Println(c)

	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]
	fmt.Printf("array type %v and value before\n %v \n", reflect.ValueOf(darr).Kind(), darr)
	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after ", "\n", darr)

	numa := [3]int{78, 79, 80}
	nums1 := numa[:] //creates a slice which contains all elements of the array
	nums2 := numa[:]
	fmt.Println("array before change 1", numa, nums1, nums2)
	nums1[0] = 100
	fmt.Println("array after modification to slice nums1", numa, nums1, nums2)
	nums2[1] = 101
	fmt.Println("array after modification to slice nums2", numa, nums1, nums2)

	callByRef(numa[:])
	fmt.Println("array after modification in local function call by ref", numa, nums1, nums2)

	callByValue(numa)
	fmt.Println("array after modification in local function call by value", numa, nums1, nums2)

	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of array %d capacity %d\n", len(fruitarray), cap(fruitarray))
	fmt.Printf("length of slice %d capacity %d\n", len(fruitslice), cap(fruitslice)) //length of fruitslice is 2 and ca

	fmt.Println(fruitslice)

	fruitslice = fruitslice[:cap(fruitslice)-1]
	fmt.Println("After re-slicing length is", len(fruitslice), "and capacity is", cap(fruitslice))
	fmt.Println(fruitslice)

	arr1 := [6]int{10, 20, 21, 22, 23, 24}
	fmt.Println(arr1)
	slic1 := arr1[4:5]
	fmt.Println(slic1)

	sliceByMake := make([]rune, 8, 20)
	fmt.Println(sliceByMake)

	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	cars = append(cars, "A")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6
	cars = append(cars, "B")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	cars = append(cars, "C")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	var names []string
	if len(names) == 0 { // names==nil
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// veggies := []string{"potatoes", "tomatoes", "brinjal"}
	// fruits := []string{"oranges", "apples"}
	// food := append(veggies, fruits...)
	// fmt.Println("food:", food)

	banksIds := []int{10, 20, 30, 40, 50}
	chequesIds := []int{1, 2, 3}

	ids := append(banksIds, chequesIds...)
	fmt.Println(ids)
	ids = banksIds[:]
	fmt.Println(ids)
	ids = append(ids, chequesIds[0], chequesIds[1], chequesIds[2])
	fmt.Println(ids)
	ids = banksIds[:]
	fmt.Println(ids)
	ids = append(ids, chequesIds[0:3]...)
	fmt.Println(ids)

	pls := [][]string{
		{"C", "C++"},
		{"JavaScript"},
		{"Go", "Rust"},
	}
	for _, v1 := range pls {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}

	countries()

}

func changeLocal(num [5]int) {
	num[0] = 11
	fmt.Println("inside function ", num)

}

// call by Ref
func callByRef(slice1 []int) {
	if len(slice1) > 2 {
		slice1[2] = 500
	}
}

func callByValue(arr [3]int) {
	if len(arr) > 2 {
		arr[2] = 5
	}
	fmt.Println("array in function ", arr)
}

type Slice struct {
	Length        int
	Capacity      int
	ZerothElement *byte
}

func countries() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	fmt.Printf("%v\n\n", countries)
	neededCountries := countries[:len(countries)-2]
	fmt.Printf("%v\n\n", neededCountries)
	countriesCpy := make([]string, len(neededCountries))
	fmt.Printf("%v\n\n", countriesCpy)
	copy(countriesCpy, neededCountries) //copies neededCountries to countriesCpy
	fmt.Printf("%v\n\n", countriesCpy)
	return countriesCpy
}
