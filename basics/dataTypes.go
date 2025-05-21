package basics

import "fmt"

func DataTypes() {

	/*
		Data type is an important concept in programming. Data type specifies the size and type of variable values.
		Go is statically typed, meaning that once a variable type is defined, it can only store data of that type.

		Go has three basic data types:
			bool: represents a boolean value and is either true or false
			Numeric: represents integer types, floating point values, and complex types
			string: represents a string value

	*/

	/*
						Go Integer Data Types
						Integer data types are used to store a whole number without decimals, has two categories:
						    Signed integers - can store both positive and negative values
							Unsigned integers - can only store non-negative values
						Tip: The default type for integer is int. If you do not specify a type, the type will be int.
		        ------------------------------------------------------------------------------------------------------
				Type	Size	Range
				------------------------------------------------------------------------------------------------------
				int	  Depends on platform: 	32 bits in 32 bit systems and
											64 bit in 64 bit systems	-2147483648 to 2147483647 in 32 bit systems and
											-9223372036854775808 to 9223372036854775807 in 64 bit systems
				int8	8 bits/1  byte	-128 to 127
				int16	16 bits/2 byte	-32768 to 32767
				int32	32 bits/4 byte	-2147483648 to 2147483647
				int64	64 bits/8 byte	-9223372036854775808 to 9223372036854775807
				------------------------------------------------------------------------------------------------------
				uint	Depends on platform: 32 bits in 32 bit systems and
											 64 bit in 64 bit systems	0 to 4294967295 in 32 bit systems and
											 0 to 18446744073709551615 in 64 bit systems
				uint8	8 bits/1 byte	0 to 255
				uint16	16 bits/2 byte	0 to 65535
				uint32	32 bits/4 byte	0 to 4294967295
				uint64	64 bits/8 byte	0 to 18446744073709551615
	*/

	var varInt int = -2147483648
	var varInt8 int8 = -127
	var varInt16 int16 = -32767
	var varInt32 int32 = -2147483647
	var varInt64 int64 = -9223372036854775807

	fmt.Printf("%v\n", varInt)
	fmt.Printf("%v\n", varInt8)
	fmt.Printf("%v\n", varInt16)
	fmt.Printf("%v\n", varInt32)
	fmt.Printf("%d\n", varInt64)

	var varUInt uint = 2147483648
	var varUInt8 uint8 = 127
	var varUInt16 uint16 = 32767
	var varUInt32 uint32 = 4294967295
	var varUInt64 uint64 = 18446744073709551615

	fmt.Printf("%v\n", varUInt)
	fmt.Printf("%v\n", varUInt8)
	fmt.Printf("%v\n", varUInt16)
	fmt.Printf("%v\n", varUInt32)
	fmt.Printf("%d\n", varUInt64)

	/*-----------------------------------------------------------------------
		Go Float Data Types
		The float data types are used to store positive and negative numbers with a decimal point, like 35.3, -2.34, or 3597.34987.
	-----------------------------------------------------------------------
		Type	Size	Range
		float32	32 bits	-3.4e+38 to 3.4e+38.
		float64	64 bits	-1.7e+308 to +1.7e+308.
		--Tip: The default type for float is float64. If you do not specify a type, the type will be float64.


	-----------------------------------------------------------------------

	*/

	flt := 5.556325 // float64 by default
	fmt.Printf("%.2f", flt)

	//String Data Type :  is used to store a sequence of characters (text). String values must be surrounded by double quotes:
}
