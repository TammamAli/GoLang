package basics

import (
	"fmt"
	"math"
	"unsafe"
)

var (
	AppVerson float32 = 1.0
	ConString         = "set connection string"
)

func Variables() {
	fmt.Println("Go Variables")

	/*
		   Variables are containers for storing data values.
		   Variable is the name given to a memory location to store a value of a specific type.
		   There are various syntaxes to declare variables in Go.

			In Go, there are different types of variables, for example:

			int- stores integers (whole numbers), such as 123 or -123
			float32- stores floating point numbers, with decimals, such as 19.99 or -19.99
			string - stores text, such as "Hello World". String values are surrounded by double quotes
			bool- stores values with two states: true or false
	*/

	//Declaring (Creating) Variables
	// 1- With the var keyword:  Use the var keyword, followed by variable name and type:
	//Syntax   var variablename type = value
	//Note: You always have to specify either type or value (or both).
	//declare without type, In this case, the type of the variable is inferred from the value

	var age int
	fmt.Println("the age is ", age)

	var number int = 10 //type is int
	fmt.Println(number)

	var str = "string variables" // The type of str will be inferred

	fmt.Println(str)

	isActive := true // shor hand declaration
	fmt.Println(isActive)

	//Variable Declaration Without Initial Value
	//In Go, all variables are initialized. So, if you declare a variable without an initial value,
	// its value will be set to the default value of its type:  "" for string , 0 for numbers false for bool

	var a string
	var b int
	var c bool

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	// var student1 string
	// student1 = "John"
	// fmt.Println(student1)

	/*
				Difference Between var and :=
		var
			Can be used inside and outside of functions
			variable declaration and value assignment can be done separately
		:=
			Can only be used inside functions
			Variable declaration and value assignment cannot be done separately (must be done in the same line)
	*/

	//Go Multiple Variable Declaration: it is possible to declare multiple variables on the same line.
	//Note: If you use the type keyword, it is only possible to declare one type of variable per line.
	//If the type keyword is not specified, you can declare different types of variables on the same line:
	///Go Variable Declaration in a Block

	var n1, n2, n3, n4 int = 1, 2, 3, 4
	fmt.Println(n1, n2, n3, n4)

	var str1, str2, str3 string = "a", "b", "C"
	fmt.Println(str1, str2, str3)

	var b1, b2 bool = true, false
	fmt.Println(b1, b2)

	var vNum, vStr, vBool = 10, "string ", true
	fmt.Println(vNum, vStr, vBool)

	c1, c2, c3 := 1, "c2", false

	fmt.Println(c1, c2, c3)

	var (
		a1, a2, a3 int     = 1, 2, 3
		d1, d2     float32 = 10.3, 20.5
		s1, s2     string  = "s1", "s2"
		bl1, bl2   bool    = true, true
	)

	fmt.Println(a1, a2, a3, d1, d2, s1, s2, bl1, bl2)
	/*
			 Go Variable Naming Rules
				A variable can have a short name (like x and y) or a more descriptive name (age, price, carname, etc.).

			Go variable naming rules:
			`	A variable name must start with a letter or an underscore character (_)
				A variable name cannot start with a digit
				A variable name can only contain alpha-numeric characters and underscores (a-z, A-Z, 0-9, and _ )
				Variable names are case-sensitive (age, Age and AGE are three different variables)
				There is no limit on the length of the variable name
				A variable name cannot contain spaces
				The variable name cannot be any Go keywords

			Multi-Word Variable Names
				Variable names with more than one word can be difficult to read.
				There are several techniques you can use to make them more readable:

				1- Camel Case  : MyVariableName = "John"
		        2- Pascal Case : MyVariableName = "John"
				3- Snake Case  : my_variable_name = "John"
	*/

	//Multiple variable declaration

	var v1, v2, v3 = 100, "200", false
	fmt.Printf("%#v , %#v, %#v", v1, v2, v3)

	var (
		g1 = "asdasd"
		g2 = 20
		g3 int
		g5 string
	)

	fmt.Println(g1, g2, g3, g5)

	r1, r2, r3 := 10, "", false
	fmt.Println(r1, r2, r3)

}

func Variables2() {
	a, b := 20, 30 // declare variables a and b
	fmt.Println("a is", a, "b is", b)
	b, c := 40, 50 // b is already declared but c is new
	fmt.Println("b is", b, "c is", c)
	b, c = 80, 90 // assign new values to already declared variables b and c
	fmt.Println("changed b is", b, "c is", c)

	k := 10.6
	k, d := 4, 20.5
	fmt.Println(k, d)
	fmt.Println(math.Min(k, d))

	var byt byte = 8
	fmt.Println(byt)

	var tn rune = 10000
	fmt.Println(tn)

	var intVar int = 20

	fmt.Printf("type of intVar is %T, size of a is %d bytes\n", intVar, unsafe.Sizeof(intVar)) //type and size of a

	c1 := complex(5, 7)
	fmt.Println(c1)
}

func DataConvert() {

	a := 20.3
	c := 10
	//sum := a + c  //invalid operation: a + c (mismatched types float64 and int)
	sum := a + float64(c)
	fmt.Println(a, c, sum)

	const n = "Sam"
	var name = n
	fmt.Printf("type %T value %v\n", name, name)

	var defaultName string = "Osama"
	fmt.Println(defaultName)

	type MyString string

	var customString MyString = MyString(defaultName)
	//var customString MyString = MyString(defaultName)
	fmt.Println(customString)

	var i = 5
	var f = 5.6
	var c2 = 5 + 6i
	fmt.Printf("i's type is %T, f's type is %T, c's type is %T\n", i, f, c2)
	fmt.Printf("i : %v, f: %v, c: %v\n", i, f, c2)
}

func Test() {
	const c = 5
	var intVar int = c
	var int32Var int32 = c
	var float64Var float64 = c
	var complex64Var complex64 = c
	fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	// var r = 4
	// intVar = r
	// int32Var = r
	// float64Var = r
	// complex64Var = r
	// fmt.Println("intVar", intVar, "\nint32Var", int32Var, "\nfloat64Var", float64Var, "\ncomplex64Var", complex64Var)

	const pi float64 = 3.14
	fmt.Println(pi)

}
