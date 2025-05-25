package advanced

import (
	"fmt"
)

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    float32
	address   Address
}

type Address struct {
	city  string
	state string
}

type Manager struct {
	id        rune
	surName   string
	age       int
	netSalary float32
	JobTitle
}

type JobTitle struct {
	jobId       int
	jobDescribe string
}

type Person struct {
	string
	int
	float32
}

type image struct {
	data map[int]int
}

func Structs() {
	fmt.Println("Struct in Golang")

	/*
		A struct is a user-defined type that represents a collection of fields.
		It can be used in places where it makes sense to group the data into a single unit rather
		than having each of them as separate values.

		Creating named structs
		The order of the fields need not necessarily be the same as that of the order of the field names while declaring the struct type.

		Creating non-named structs
		is defined by omitting the field names. In this case, it is necessary to maintain the order of fields to be the same
		as specified in the struct declaration. Please refrain from using this syntax since it makes it difficult to figure
		out which value is for which field.

		Creating anonymous structs
		It is possible to declare structs without creating a new data type. These types of structs are called anonymous structs.
		an anonymous struct variable emp3 is defined. As we have already mentioned, this struct is called anonymous because it only
		creates a new struct variable emp3 and does not define any new struct type like named structs.

		Accessing individual fields of a struct
		The dot . operator is used to access the individual fields of a struct.

		Zero value of a struct
		When a struct is defined and it is not explicitly initialized with any value, the fields of the struct are
		assigned their zero values by default.
		It is also possible to specify values for some fields and ignore the rest. In this case,
		the ignored fields are assigned zero values.

		Pointers to a struct
		It is also possible to create pointers to a struct.

		Anonymous fields
		It is possible to create structs with fields that contain only a type without the field name.
		These kinds of fields are called anonymous fields. Even though anonymous fields do not have an explicit name,
		by default the name of an anonymous field is the name of its type.

		Nested structs
		It is possible that a struct contains a field which in turn is a struct. These kinds of structs are called nested structs.

		Promoted fields
		Fields that belong to an anonymous struct field in a struct are called promoted fields since they can be accessed
		as if they belong to the struct which holds the anonymous struct field. I can understand that this definition
		is quite complex so let’s dive right into some code to understand this :).

		Exported structs and fields
		If a struct type starts with a capital letter, then it is an exported type and it can be accessed from other packages.
		Similarly, if the fields of a struct start with caps, they can be accessed from other packages.


		Structs Equality
		Structs are value types and are comparable if each of their fields are comparable.
		Two struct variables are considered equal if their corresponding fields are equal.
		Struct variables are not comparable if they contain fields that are not comparable


	*/
	//empStruct()

	// personStruct()

	ad := Address{city: "Assuit", state: "Assuit"}
	emp1 := Employee{
		firstName: "Ali",
		lastName:  "Ahmed",
		age:       30,
		salary:    15_000,
		address:   ad,
	}
	fmt.Println(emp1)

	emp2 := Employee{
		firstName: "Ali",
		lastName:  "Ahmed",
		age:       30,
		salary:    15_000,
		address:   Address{city: "Cairo", state: "Cairo"},
	}
	fmt.Println(emp2)
	fmt.Println(emp2.address.city)

	mngr1 := Manager{
		id:        1,
		surName:   "Alaa Ahmed",
		age:       31,
		netSalary: 19_000.500,
		JobTitle:  JobTitle{jobId: 1, jobDescribe: "Sales-Manager"},
	}

	fmt.Println(mngr1)
	fmt.Printf("id:%v, surame:%v, age:%v, ", mngr1.id, mngr1.surName, mngr1.age)

	//jobId & jobDescribe are promoted fields
	fmt.Printf("netSalary:%v, jobId:%v, jobDescribe:%v", mngr1.netSalary, mngr1.jobId, mngr1.jobDescribe)
	fmt.Println()
	compareStruct()

}
func compareStruct() {
	type name struct {
		firstName string
		lastName  string
	}

	nme1 := name{firstName: "Ahmed", lastName: "ali"}
	nme2 := name{firstName: "Ahmed", lastName: "Ali"}

	if nme1 == nme2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}

	img1 := image{
		data: map[int]int{0: 155},
	}

	img2 := image{
		data: map[int]int{0: 155},
	}

	fmt.Println(img1)
	fmt.Println(img2)

	//invalid operation: image1 == image2 (struct containing map[int]int cannot be compared)
	/*if img1 == img2 {
		fmt.Println("name1 and name2 are equal")
	} else {
		fmt.Println("name1 and name2 are not equal")
	}
	*/

}

func PersonStruct() {
	prsn := Person{
		string:  "Any one",
		int:     25,
		float32: 25_000,
	}
	fmt.Println(prsn)
	fmt.Println(prsn.string)
	fmt.Println(prsn.int)
	fmt.Println(prsn.float32)

}
func EmpStruct() {
	//Creating named structs ,ct specifying field names
	named_struct := Employee{
		firstName: "Amr",
		lastName:  "Ali",
		age:       42,
		salary:    12_000.00,
	}
	fmt.Println(named_struct)

	var non_named_struct Employee = Employee{"Ahmed", "Ali", 40, 10_000, Address{"", ""}}
	fmt.Println(non_named_struct)

	manger := struct {
		id  int
		job string
	}{
		id:  10,
		job: "team-leader",
	}

	fmt.Printf("print anonymous struct %v\n", manger)

	emp1 := Employee{
		firstName: "Amr",
		lastName:  "Ahmed",
		age:       35,
		salary:    2500,
	}
	printEmp(emp1)

	// 	Zero value of a struct
	emp2 := Employee{}
	printEmp(emp2)

	emp3 := Employee{lastName: "Ali", age: 20}
	printEmp(emp3)

	emp8 := &Employee{
		firstName: "Sam",
		lastName:  "Anderson",
		age:       55,
		salary:    6000,
	}
	fmt.Println(emp8) // is different
	fmt.Println("First Name:", (*emp8).firstName)
	fmt.Println("Age:", (*emp8).age)

	//The Go language gives us the option to use emp8.firstName instead of the explicit dereference (*emp8).firstName to access the firstName field.
	fmt.Println("First Name:", emp8.firstName)
	fmt.Println("Age:", emp8.age)

	emp9 := Employee{
		firstName: "emp",
		lastName:  "9",
		age:       32,
		salary:    6_000,
	}

	emp10 := &emp9
	fmt.Println(emp9)
	fmt.Println(emp10)

	emp9.age = 50
	fmt.Println(emp9)
	fmt.Println(emp10)

}
func printEmp(emp Employee) {
	fmt.Printf("Employee is %v\n", emp)
	fmt.Printf("FirstName:%v, LastName: %v, Age: %v, Salary:%v\n", emp.firstName, emp.lastName, emp.age, emp.salary)
}
