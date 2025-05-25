package advanced

import (
	"fmt"
	"reflect"
)

func Maps() {
	/*
		What is a map?
		A map is a inbuilt data type in Go which is used to store key-value pairs.
		A practical use case for a map is for storing the currency codes and the corresponding currency names.
		A map will be a perfect fit for the above use case use case.
		The currency code can be the key and the currency name can be the value.
		Maps are similar to dictionaries in other languages such as Python.

		How to create a map?
		A map can be created by passing the data type of key and value to the make function. The following is the syntax to create a new map.
			make(map[type of key]type of value)

		As you might have recognized from the above output, the order of the retrieval of values from a map is not guaranteed
		to be the same as the order in which the elements were added to the map.

		It is also possible to initialize a map during the declaration itself.

		It’s not necessary that only string types should be keys.
		All comparable types such as boolean, integer, float, complex, string, … can also be keys.
		Even user-defined types such as structs can be keys.

		nil map panics
		The zero value of a map is nil. If you try to add elements to a nil map, a run-time panic will occur.
		Hence the map has to be initialized before adding elements.

		There will be no runtime error when we try to retrieve the value for a key that is not present in the map,
		it will returns empty string as the item value.

		Checking if a key exists
		The following syntax is used to find out whether a particular key is present in a map.
			value, ok := map[key]    => ok is bool
		ok in the above line of code will be true when the key is present and the value for the key is present in the variable value.
		If the key is not present, ok will be false and the zero value is returned for value.

		Iterate over all elements in a map
		The range form of the for loop is used to iterate over all elements of a map.

		One important fact to note is the order of the retrieval of values from a map when using for range is not guaranteed
		to be the same for each execution of the program. It is also not the same as the order in which
		the elements were added to the map

		Deleting items from a map
		delete(map, key) is the syntax to delete key from a map. The delete function does not return any value.
		Even if we try to delete a key that is not present in the map, there will be no runtime error.

		Map of structs
		So far we have only been storing the currency name in the map.
		Wouldn’t it be nice if we are able to store the symbol of the currency too?
		This can be achieved by using a map of structs. The currency can be represented as a
		struct containing the fields currency name and currency symbol.
		This struct value can be stored in the map with a currency code as key .
		Let’s write a program to understand how this can be done.

		Length of the map: Length of the map can be determined using the len function.

		Maps are reference types:
			Similar to slices, maps are reference types.
			When a map is assigned to a new variable, they both point to the same underlying data structure.
			Hence changes made in one will reflect in the other.

			Similar is the case when maps are passed as parameters to functions.
			When any change is made to the map inside the function, it will be visible to the caller also.

		Maps equality
		   Maps can’t be compared using the == operator. The == can be only used to check if a map is nil.
		   One way to check whether two maps are equal is to compare each one’s individual elements one by one.
		   The other way is using reflection. I would encourage you to write a program for this and make it work :).












	*/

	currencyCode := make(map[string]string)
	currencyCode["USD"] = "US Dollar"
	currencyCode["GBP"] = "Pound Sterling"
	currencyCode["EUR"] = "Euro"
	currencyCode["INR"] = "Indian Rupee"
	fmt.Printf("%T\n", currencyCode)
	fmt.Printf("%v\n", reflect.TypeOf(currencyCode))
	fmt.Printf("%v\n", reflect.ValueOf(currencyCode).Kind())
	fmt.Printf("%v\n", currencyCode)

	fmt.Printf("\n\n")
	for indx, value := range currencyCode {
		fmt.Printf("currencyCode[%v] = %v \n", indx, value)

	}
	fmt.Printf("\n\n")
	words :=
		map[string]string{
			"a": "apple",
			"b": "book",
		}
	words["c"] = "car"

	fmt.Println(words)

	var mapInt map[int]int
	fmt.Println(mapInt)
	//mapInt is nil and we are trying to add a new key to a nil map. The program will panic with error
	//mapInt[1] = 10 //nil dereference in map
	fmt.Println(`words["d"] is:`, words["d"])

	key := "a1"

	if val, ok := words[key]; ok {
		fmt.Println(val)
	} else {
		fmt.Printf("the key %v is not found in this map\n", key)
	}

	fmt.Println(words)
	delete(words, "a22")
	fmt.Println(words)
	words["a"] = "apple 1"
	words["a"] = "apple 2"
	fmt.Println(words)

	curUSD := currency{
		name:   "US Dollar",
		symbol: "$",
	}
	curGBP := currency{
		name:   "Pound Sterling",
		symbol: "£",
	}
	curEUR := currency{
		name:   "Euro",
		symbol: "€",
	}

	currencies := map[string]currency{
		"USD": curUSD,
		"GBP": curGBP,
		"EUR": curEUR,
	}

	for cyCode, cyInfo := range currencies {
		fmt.Printf("Currency Code: %s, Name: %s, Symbol: %s\n", cyCode, cyInfo.name, cyInfo.symbol)
	}

	currencies2 := make(map[string]currency)
	currencies2["EG"] = currency{name: "Egypt", symbol: "£"}
	fmt.Println(currencies2)

	salaries := map[string]float32{}
	salaries["a"] = 125.25
	salaries["b"] = 225.25
	salaries["c"] = 325.25
	fmt.Println(salaries)

	newSalaries := salaries
	newSalaries["a"] = 225.25
	fmt.Println(newSalaries)
	fmt.Println(salaries)

	var mapTest map[string]int
	if mapTest != nil {
		mapTest = make(map[string]int)
		fmt.Println("equal zeor")
	}
	mapTest["a"] = 1

	map3 := map[string]int{"a": 1, "b": 2, "c": 3}
	map4 := map[string]int{"a": 1, "b": 2, "c": 3}
	fmt.Println("mapEqual", mapEqual(map3, map4))

}

type currency struct {
	name   string
	symbol string
}

func mapEqual(a, b map[string]int) bool {
	if len(a) != len(b) {
		return false
	}

	for k, v := range a {
		if bval, ok := b[k]; !ok || bval != v {
			return false
		}
	}
	return true
}
