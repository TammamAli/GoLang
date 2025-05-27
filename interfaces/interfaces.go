package interfaces

import "fmt"

type Describer interface {
	Describe()
}

type Person struct {
	name string
	age  int
}

type Address struct {
	state   string
	country string
}

func (p Person) Describe() { //implemented using value receiver
	fmt.Printf("%s is %d years old\n", p.name, p.age)
}

func (a *Address) Describe() { //implemented using pointer receiver
	fmt.Printf("%s is in %s\n", a.state, a.country)
}

func StartInterface() {

	var d1 Describer

	p1 := Person{name: "Amr", age: 32}
	p1.Describe()

	d1 = p1
	d1.Describe()

	p2 := Person{name: "alaa", age: 25}
	p2.Describe()

	d1 = &p2
	d1.Describe()

	ad := Address{country: "Egypt", state: "Cairo"}
	ad.Describe()

	// d1 = ad //cannot use ad (variable of struct type Address) as Describerr value in assignment: Address does not implement Describerr (method Describe has pointer receiver)
	d1.Describe()

	//Zero Type Interface
	var d4 Describer2

	d4.Describe()
	if d4 == nil {
		fmt.Printf("\n\nd1 is nil and has type %T value %v\n\n", d4, d4)
	}
}

type Describer2 interface {
	Describe()
}
