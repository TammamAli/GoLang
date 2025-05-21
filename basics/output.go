package basics

import "fmt"

func Output() {

	//Go Output Functions
	/*
		Print()
		Println()
		Printf()
	*/

	//The Print() function prints its arguments with their default format.
	//If we want to print the arguments in new lines, we need to use \n.
	//It is also possible to only use one Print() for printing multiple variables.
	//If we want to add a space between string arguments, we need to use " ":
	//Print() inserts a space between the arguments if neither are strings:

	var i, j string = "Hello", "World"

	fmt.Print(i)
	fmt.Print(j)
	fmt.Print("\n")

	fmt.Print(i, "\n")
	fmt.Print(j, "\n")

	fmt.Print(i, "\n", j)
	fmt.Print("\n")

	fmt.Print(i, " ", j)
	fmt.Print("\n")

	var a, b int = 20, 30
	fmt.Print(a, b, "\n")

	//The Println() function is similar to Print() with the difference that
	// a whitespace is added between the arguments, and a newline is added at the end:

	//The Printf() Function: first formats its argument based on the given formatting verb and then prints them.
	/*
		Here we will use two formatting verbs:
			%v is used to print the value of the arguments
			%T is used to print the type of the arguments
	*/

	fmt.Printf("i has value: %v and type: %T\n", i, i)
	fmt.Printf("a has value: %v and type: %T", a, a)
}

func OutputFormat() {
	fmt.Println("Output Format ")

	//Go offers several formatting verbs that can be used with the Printf() function.

	/*
		General Formatting Verbs : The following verbs can be used with all data types:
		-------------------------------------------------
			Verb	Description
		-------------------------------------------------
			%v		Prints the value in the default format
			%#v		Prints the value in Go-syntax format
			%T		Prints the type of the value
			%%		Prints the % sign
		-------------------------------------------------
	*/

	fmt.Printf("Uning Printf\n")
	vInt := 10
	vFloat := 2.5
	vString := "Word"
	vBool := true

	fmt.Printf("vInt is %T , equal %v\n", vInt, vInt)
	fmt.Printf("vFloat is %T , equal  %v%%\n", vFloat, vFloat)
	fmt.Printf("vString is %T , equal  %v\n", vString, vString)
	fmt.Printf("vBool is %T , equal  %v\n", vBool, vBool)

	fmt.Printf("\n\n")

	fmt.Printf("vInt is %T , equal %#v\n", vInt, vInt)
	fmt.Printf("vFloat is %T , equal  %#v\n", vFloat, vFloat)
	fmt.Printf("vString is %T , equal  %#v\n", vString, vString)
	fmt.Printf("vBool is %T , equal  %#v\n", vBool, vBool)

	/*
		Integer Formatting Verbs
		--------------------------------------------------------------------------
		Verb	Description
		--------------------------------------------------------------------------
		%b		Base 2
		%d		Base 10
		%+d		Base 10 and always show sign
		%o		Base 8
		%O		Base 8, with leading 0o
		%x		Base 16, lowercase
		%X		Base 16, uppercase
		%#x		Base 16, with leading 0x
		%4d		Pad with spaces (width 4, right justified)
		%-4d	Pad with spaces (width 4, left justified)
		%04d	Pad with zeroes (width 4
		--------------------------------------------------------------------------
	*/

	fmt.Printf("\n\nInteger Formatting Verbs \n")
	var i = 15
	fmt.Printf("%v %%b into Base 2: %b\n", i, i)
	fmt.Printf("%v %%d into Base 10: %d\n", i, i)
	fmt.Printf("%v %%+d into Base 10 with sign: %+d\n", i, i)
	fmt.Printf("%v %%o into Base 8: %o\n", i, i)
	fmt.Printf("%v %%O into Base 8: %O with leading 0o\n", i, i)
	fmt.Printf("%v %%x into Base 16 lowercase: %x\n", i, i)
	fmt.Printf("%v %%X into Base 16 uppercase: %X\n", i, i)
	fmt.Printf("%v %%0x into Base 16 with leading 0x: %#x\n", i, i)
	fmt.Printf("%v %%4d iPad with spaces (width 4, right justified): %4d\n", i, i)
	fmt.Printf("%v %%-4d iPad with spaces (width 4, left justified): %-4d\n", i, i)
	fmt.Printf("%v %%04d Pad with zeroes (width 4: %04d\n", i, i)
	fmt.Printf("\n\n")
	/*
		String Formatting Verbs
		--------------------------------------------------------------------------
		Verb	Description
		--------------------------------------------------------------------------
		%s		Prints the value as plain string
		%q		Prints the value as a double-quoted string
		%8s		Prints the value as plain string (width 8, right justified)
		%-8s	Prints the value as plain string (width 8, left justified)
		%x		Prints the value as hex dump of byte values
		% x		Prints the value as hex dump with spaces
		--------------------------------------------------------------------------
	*/

	var txt = "Hello"

	fmt.Printf("%s %%s as plain string \n", txt)
	fmt.Printf("%q %%q as double-quoted string \n", txt)
	fmt.Printf("%#v %%#v as double-quoted string \n", txt)
	fmt.Printf("%8s %%8s as plain string (width 8, right justified) \n", txt)
	fmt.Printf("%-8s %%-8s as plain string (left 8, right justified) \n", txt)
	fmt.Printf("%x %%s as hex dump of byte values \n", txt)
	fmt.Printf("% x %%s as hex dump with spaces\n", txt)
	fmt.Printf("\n\n")

	//Boolean Formatting Verbs
	//%t	Value of the boolean operator in true or false format (same as using %v)

	var j = true
	fmt.Printf("%t\n", j)

	j = false
	fmt.Printf("%t\n", j)
	fmt.Printf("\n\n")
	/*
			Float Formatting Verbs
		    --------------------------------------------------------------------------
			Verb	Description
			--------------------------------------------------------------------------
			%e		Scientific notation with 'e' as exponent
			%f		Decimal point, no exponent
			%.2f	Default width, precision 2
			%6.2f	Width 6, precision 2
			%g		Exponent as needed, only necessary digits
			--------------------------------------------------------------------------
	*/

	var flt = 3.141

	fmt.Printf("%e\n", flt)
	fmt.Printf("%f\n", flt)
	fmt.Printf("%.2f\n", flt)
	fmt.Printf("%6.2f\n", flt)
	fmt.Printf("%g\n", flt)
}
