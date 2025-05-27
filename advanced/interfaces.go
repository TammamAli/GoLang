package advanced

import (
	"fmt"
)

func Interfaces() {

	/*
		What is an interface?
		In Go, an interface is a set of method signatures.
		When a type provides definition for all the methods in the interface, it is said to implement the interface.
		Interface specifies what methods a type should have and the type decides how to implement these methods.

		For example, PaymentProcessor can be an interface with method signatures ProcessPayment() and GenerateReceipt().
		Any type which provides definitions for ProcessPayment() and GenerateReceipt()
		methods is said to implement the PaymentProcessor interface.

		This can include structs like CreditCardProcessor, PayPalProcessor or DirectDebitProcessor each of which implements the methods
		in a way specific to their payment system.

		interface defines the methods
		Type (struct) decides how to implement these method

		We add the method CalculateSalary() int to the receiver type Permanent. Now Permanent is said to implement the interface
		SalaryCalculator. This is quite different from other languages like Java where a class has to explicitly state that it
		implements an interface using the implements keyword. This is not needed in Go and Go interfaces are implemented implicitly
		if a type contains all the methods declared in the interface.

		The biggest advantage of this is that totalExpense can be extended to any new employee type without any code changes.
		Let’s say the company adds a new type of employee Freelancer with a different salary structure.
		This Freelancer can just be passed in the slice argument to totalExpense without even a single line of code change
		to the totalExpense function. This method will do what it’s supposed to do as Freelancer will also implement the
		SalaryCalculator interface :).

		Interface internal representation
		An interface can be thought of as being represented internally by a tuple (type, value).
		type is the underlying concrete type that implements the interface and value holds the value of the concrete type.

		Empty interface
		An interface that has zero methods is called an empty interface. It is represented as interface{}.
		Since the empty interface has zero methods, all types implement the empty interface.

		Type switch
		A type switch is used to compare the concrete type of an interface against multiple types specified in various case statements.
		It is similar to switch case. The only difference being the cases specify types and not values as in normal switch.

		The syntax for type switch is similar to Type assertion. In the syntax i.(T) for Type assertion,
		the type T should be replaced by the keyword type for type switch. Let’s see how this works in the program below.

		It is also possible to compare a type to an interface. If we have a type and if that type implements an interface,
		it is possible to compare this type with the interface it implements.


		Implementing interfaces using pointer receivers vs value receivers
		All the interfaces we discussed in part 1 were implemented using value receivers. It is also possible to implement
		interfaces using pointer receivers. There is a subtlety to be noted while implementing interfaces using pointer receivers.
		Lets understand that using the following program.


		As we have already learnt during our discussion about methods, methods with value receivers accept both pointer
		and value receivers. It is legal to call a value method on anything which is a value or whose value
		can be dereferenced.

		The reason is that it is legal to call a pointer-valued method on anything that is already a pointer or whose
		address can be taken. The concrete value stored in an interface is not addressable and hence it is not possible
		for the compiler to automatically take the address of a in line no. 45 and hence this code fails.


		Embedding interfaces
		Although Go does not support inheritance, embedding interfaces is possible.
		It is possible to create a new interface by embedding other interfaces.


		Any type is said to implement EmployeeOperations interface if it provides method definitions
		for the methods present in both SalaryCalculator and LeaveCalculator interfaces.


		Zero value of Interface
		The zero value of a interface is nil. A nil interface has both its underlying value and as well as concrete type as nil.






	*/

	fmt.Println("Interfaces...")

	pemp1 := Permanent{
		empId:    1,
		basicpay: 5000,
		pf:       30,
	}

	pemp2 := Permanent{
		empId:    1,
		basicpay: 3000,
		pf:       20,
	}

	emp3 := Contract{
		empId:    1,
		basicpay: 2500,
	}

	emps := []SalaryCalculator{pemp1, pemp2, emp3}
	totalExpense(emps)

	freelancer1 := Freelancer{
		empId:       4,
		ratePerHour: 70,
		totalHours:  120,
	}
	freelancer2 := Freelancer{
		empId:       5,
		ratePerHour: 100,
		totalHours:  100,
	}
	emps = []SalaryCalculator{pemp1, pemp2, emp3, freelancer1, freelancer2}
	totalExpense(emps)

}

type SalaryCalculator interface {
	CalculateSalary() int
}

type LeaveCalculator interface {
	CalculateLeavesLeft() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

type Freelancer struct {
	empId       int
	ratePerHour int
	totalHours  int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

// salary of contract employee is the basic pay alone
func (p Contract) CalculateSalary() int {
	return p.basicpay
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func totalExpense(cs []SalaryCalculator) {
	var expense int = 0
	for _, v := range cs {
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d\n\n", expense)
}
