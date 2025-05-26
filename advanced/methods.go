package advanced

import (
	"fmt"
	"math"
)

type Customer struct {
	name     string
	balance  int
	currency string
}

type Rectangle struct {
	length int
	width  int
}

type Circle struct {
	radius float64
}

func (r Rectangle) Area() int {
	rsult := (r.length * r.width)
	fmt.Printf("Area Method result: %d\n", rsult)
	return rsult
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}
func Area(r Rectangle) {
	fmt.Printf("Area Function result: %d\n", (r.length * r.width))
}

func perimeter(r *Rectangle) {
	fmt.Println("perimeter function output:", 2*(r.length+r.width))

}

func (r *Rectangle) perimeter() {
	fmt.Println("perimeter method output:", 2*(r.length+r.width))
}
func (c Customer) Info() {
	fmt.Printf("Customer info of %s is %d%s\n\n", c.name, c.balance, c.currency)
}

func (e Employee) Info() {
	fmt.Printf("Employee info of %s is %s %d %f %v\n\n", e.firstName, e.lastName, e.age, e.salary, e.address)

}

func displaySalary(e Customer) {
	fmt.Printf("Salary of %s is %d%s\n\n", e.name, e.balance, e.currency)
}

type myInt int

func (a myInt) add(b myInt) myInt {
	return a + b
}

func Methods() {

	fmt.Println("Methods...")

	/*
		A method is just a function with a special receiver type between the func keyword and the method name.
		The receiver can either be a struct type or non-struct type.
		The syntax of a method declaration is provided below.

		func (t Type) methodName(parameter list) {
			//ToDo
		}

		The above snippet creates a method named methodName with receiver type Type.
		t is called as the receiver and it can be accessed within the method.

		So why do we have methods when we can write the same program using functions.
		There are a couple of reasons for this. Let’s look at them one by one.

		.Methods vs Functions
		 The above program can be rewritten using only functions and without methods.
		 So why do we have methods when we can write the same program using functions.
		 There are a couple of reasons for this. Let’s look at them one by one.

		 1- Go is not a pure object-oriented programming language and it does not support classes. Hence methods on types are a way to achieve behaviour similar to classes.
			Methods allow a logical grouping of behaviour related to a type similar to classes.
			In the above sample program, all behaviours related to the Employee type can be grouped by
			creating methods using Employee receiver type. For example, we can add methods
			like calculatePension, calculateLeaves and so on.

		2- Methods with the same name can be defined on different types whereas functions with the same names are not allowed.
		   Let’s assume that we have a Square and Circle structure.
		   It’s possible to define a method named Area on both Square and Circle.


		Pointer Receivers vs Value Receivers
			So far we have seen methods only with value receivers. It is possible to create methods with pointer receivers.
			The difference between value and pointer receiver is, changes made inside a method with a pointer receiver
			is visible to the caller whereas this is not the case in value receiver. Let’s understand this with the help of a program.

		When to use pointer receiver and when to use value receiver
			Generally, pointer receivers can be used when changes made to the receiver inside the method should be visible to the caller.
			Pointers receivers can also be used in places where it’s expensive to copy a data structure.
			Consider a struct that has many fields. Using this struct as a value receiver in a method will need
			the entire struct to be copied which will be expensive. In this case, if a pointer receiver is used,
			the struct will not be copied and only a pointer to it will be used in the method.
			In all other situations, value receivers can be used.

		Methods of anonymous struct fields
			Methods belonging to anonymous fields of a struct can be called as if they belong to the structure where
			the anonymous field is defined.

		Value receivers in methods vs Value arguments in functions
			. This topic trips most go newbies. I will try to make it as clear as possible 😀.
			. When a function has a value argument, it will accept only a value argument.
			. When a method has a value receiver, it will accept both pointer and value receivers.
			. Let’s understand this by means of an example.
			Now comes the tricky part => p.area() which accepts only a value receiver using the pointer receiver p.
			This is perfectly valid. The reason is that the line p.area(), for convenience will be interpreted by Go as (*p).area()
			since area has a value receiver.

		Pointer receivers in methods vs Pointer arguments in functions
			. Similar to value arguments, functions with pointer arguments will accept only pointers whereas methods with
			  pointer receivers will accept both pointer and value receiver.
			. We try to call the perimeter function with a value argument r.  This is not allowed since a function with a pointer argument
			   will not accept value arguments.
			. We call the pointer receiver method perimeter with a value receiver r. This is allowed and the line of code r.perimeter()
			  will be interpreted by the language as (&r).perimeter() for convenience.

		. Methods with non-struct receivers
		  So far we have defined methods only on struct types.
		  It is also possible to define methods on non-struct types, but there is a catch. To define a method on a type,
		  the definition of the receiver type and the definition of the method should be present in the same package.
		  So far, all the structs and the methods on structs we defined were all located in the same main package and hence they worked.

		 If you are trying to add a method on the built-in type int. This is not allowed since the definition of the method add and
		 the definition of type int is not in the same package. This program will throw compilation error cannot define new methods
		 on non-local type int
		 The way to get this working is to create a type alias for the built-in type int and then create a method with this type
		 alias as the receiver.











	*/

	cust1 := Customer{
		name: "alaa", balance: 2_950, currency: "Eg",
	}
	cust1.Info()
	displaySalary(cust1)

	e1 := Employee{
		firstName: "Ahmed",
		lastName:  "Ali",
		salary:    13_700,
		age:       29,
		address:   Address{city: "Assuit", state: "Un-kown"},
	}
	e1.Info()

	emp2 := Employee{
		firstName: "Ali",
		salary:    1000,
		age:       35,
		address:   Address{city: "A", state: "b"},
	}

	fmt.Println(emp2)
	emp2.ChangeAddress()
	emp2.ChangeAge()    // can call struct pointer receiver without (& operator)
	(&emp2).ChangeAge() // also right

	fmt.Println(emp2)

	p := person{
		firstname: "Ahmed",
		lastname:  "Bihery",
		address:   address{city: "asyut", state: "Salam"},
	}

	a := address{city: "cairo", state: "shrouk"}
	a.getFullAddress()
	p.getFullAddress()

	p2 := person2{
		firstname: "Ahmed",
		lastname:  "Bihery",
		add:       address{city: "asyut", state: "Salam"},
	}

	p2.add.getFullAddress()
	a.getFullAddress()

	r := Rectangle{
		width:  10,
		length: 20,
	}

	Area(r)
	r.Area()

	r1 := &r
	r1.Area()
	//area(r1) //cannot use r1 (variable of type *Rectangle) as Rectangle value in argument to

	//perimeter(r) //cannot use r (variable of struct type Rectangle) as *Rectangle value in argument to
	perimeter(&r)
	perimeter(r1)

	r.perimeter()
	r1.perimeter()

	num1 := myInt(6)
	num2 := myInt(10)
	sum := num1.add(num2)
	fmt.Println(sum)

	n1 := int(3)
	fmt.Println(n1)
}

func (e *Employee) ChangeAge() {
	e.age++
}

func (e Employee) ChangeAddress() {
	e.address.city = ""
	e.address.state = ""
}

type address struct {
	city  string
	state string
}

type person struct {
	firstname string
	lastname  string
	address
}

type person2 struct {
	firstname string
	lastname  string
	add       address
}

func (a address) getFullAddress() {
	fmt.Printf("Full address is %s, %s\n", a.city, a.state)
}
