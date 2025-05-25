package advanced

import (
	"fmt"
	"reflect"
)

func Pointers() {

	/*
		. What is a pointer?
		. Declaring pointers
		. Zero value of a pointer
		. Creating pointers using the new function
		. Dereferencing a pointer
		. Passing pointer to a function
		. Returning pointer from a function
		. Do not pass a pointer to an array as an argument to a function. Use slice instead.
		. Go does not support pointer arithmetic



		. What is a pointer: A pointer is a variable that stores the memory address of another variable.
		. Declaring pointers: *T is the type of the pointer variable which points to a value of type T.
		. Zero value of a pointer The zero value of a pointer is nil.
		. Creating pointers using the new function
			Go also provides a handy function new to create pointers.
			The new function takes a type as an argument and returns a pointer to a newly
			allocated zero value of the type passed as argument
		. Dereferencing a pointer
		    It means accessing the value of the variable to which the pointer points. *a is the syntax to deference a.

		. Returning pointer from a function
			It is perfectly legal for a function to return a pointer of a local variable.
			The Go compiler is intelligent enough and it will allocate this variable on the heap.

		Do not pass a pointer to an array as an argument to a function. Use slice instead.
		Let’s assume that we want to make some modifications to an array inside the function and the changes made to
		that array inside the function should be visible to the caller. One way of doing this is to pass a pointer
		to an array as an argument to the function. Although this way of passing a pointer to an array as an argument
		to a function and making modification to it works, it is not the idiomatic way of achieving this in Go. We have slices for this.

		Go does not support pointer arithmetic
		Go does not support pointer arithmetic which is present in other languages like C and C++.







	*/

	b := 255
	var a *int = &b // only declare with var , not short hand declaration
	fmt.Printf("Type of a is %T\n", a)
	fmt.Printf("Kind of a is %v \n", reflect.ValueOf(a).Kind())
	fmt.Println("address of b is", a)
	fmt.Println("value of pointer ", *a)

	str1 := "Golang"
	var ptr_str1 *string = &str1
	fmt.Println(str1)
	fmt.Println(ptr_str1)
	*ptr_str1 = "new string"
	fmt.Println()
	str2 := "C-sharp"
	ptr_str1 = &str2
	fmt.Println(str1)
	fmt.Println(ptr_str1)

	var ptr_int *int // zero value pointer
	if ptr_int == nil {
		fmt.Println(ptr_int)
	} else {
		ptr_int = a
		fmt.Println()
		fmt.Println(a)
		fmt.Println(ptr_int)
	}

	ptr_int2 := new(int) // declare as ptr
	fmt.Println(ptr_int2)
	fmt.Println(*ptr_int2)
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *ptr_int2, ptr_int2, ptr_int2)
	fmt.Println()

	fmt.Println("address of b is ", a)
	fmt.Println("value of b is", b)
	fmt.Println("value of b is from ptr", *a)
	fmt.Println()

	b--
	fmt.Println("value of b is", b)
	fmt.Println("value of b is from ptr", *a)
	fmt.Println()

	*a++
	fmt.Println("value of b is", b)
	fmt.Println("value of b is from ptr", *a)
	fmt.Println()

	value := 20
	fmt.Printf("value before call %v \n", value)
	chageValue(&value)
	fmt.Printf("value after call %v \n", value)

	ptr_returned := AnyThink()
	fmt.Println(ptr_returned)

	var arr [3]int = [3]int{10, 20, 40}
	fmt.Println(arr)
	PassArrayAsPtr(&arr)
	fmt.Println(arr)

	bArr := [...]int{109, 110, 111}
	p := &bArr
	//p++ //nvalid operation: p++ (non-numeric type *[3]int)
	fmt.Println(*p)

}

func chageValue(ptr *int) {
	*ptr++
}

func AnyThink() *int {
	localVar := 255
	fmt.Printf("%v\n", localVar)
	return &localVar

	/*
		The behaviour of this code is undefined in programming languages such as C and C++ as the variable localVar
		goes out of scope once the function AnyThink returns.
		But in the case of Go, the compiler does an escape analysis and allocates localVar on the heap
		as the address escapes the local scope.
	*/
}

func PassArrayAsPtr(arr *[3]int) {
	arr[0] = 80 //So (*arr)[0] in the above program can be replaced by arr[0].
	(*arr)[1] = 70
}
