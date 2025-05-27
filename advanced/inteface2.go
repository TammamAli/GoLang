package advanced

import (
	"fmt"
	"reflect"
)

type Worker interface {
	Work()
}

type Persoon struct {
	name string
	age  int
}

func (p Persoon) Work() {
	fmt.Println(p.name, "is working")
}

func Describe(w Worker) {
	fmt.Printf("Interface type %T value %v\n", w, w)
}

func Explain() {

	//Worker interface has one method Work() and Person struct type implements that interface

	p := Persoon{
		name: "Ahmed", age: 45,
	}

	fmt.Printf("p: %v\n", p)
	w := p
	Describe(w)
	w.Work()

	var w2 Worker
	Describe(w2)
	//w2.Work() //nil dereference in dynamic method callnilnessnildere

	fmt.Printf("\n\n\n")

	s := "Hello world!"
	describe(s)

	i := 55
	describe(i)

	str := struct {
		name string
	}{
		name: "Tammam",
	}
	describe(str)

	var s1 interface{} = 56
	assert(s1)

	var s2 interface{} = 1
	assert(s2)

	findType(1)
	findType(true)
	findType(float32(1.5))
	findType("Any")
	findType2("Any")

	em := Employee{firstName: "Ahmed", lastName: "Ali", age: 35}
	findType2(em)
}

func describe(i interface{}) {
	if reflect.TypeOf(i).Kind() != reflect.TypeOf(int(0)).Kind() {
		fmt.Printf("Type is %T is not an int\n", i)
		return
	}
	fmt.Printf("Type is %T, value is %v\n", i, i)
}

func assert(i interface{}) {
	s := i.(int) //get the underlying int value from i
	fmt.Println(s)

}

func findType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Printf("I am an int and my value is %d\n", i.(int))
	case string:
		fmt.Printf("I am an int and my value is %s\n", i.(string))
	case float32:
		fmt.Printf("I am an int and my value is %.3f\n", i.(float32))
	case bool:
		fmt.Printf("I am an int and my value is %v\n", i.(bool))
	default:
		fmt.Printf("type is undefined\n")
	}
}

func findType2(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("I am an int and my value is %d\n", v)
	case string:
		fmt.Printf("I am an int and my value is %s\n", v)
	case float32:
		fmt.Printf("I am an int and my value is %.3f\n", v)
	case bool:
		fmt.Printf("I am an int and my value is %v\n", v)
	case Describer:
		v.Describe()
	default:
		fmt.Printf("unkhown type\n")
	}
}

type Describer interface {
	Describe()
}

func (e Employee) Describe() {
	fmt.Printf("%s %s is %d years old\n", e.firstName, e.lastName, e.age)
}
