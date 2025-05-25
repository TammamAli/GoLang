package advanced

import (
	"fmt"
	"unicode/utf8"
)

func Strings() {
	/*
		Strings deserve a special mention in Go as they are different in implementation when compared to other languages.

		What is a String?
		A string is a slice of bytes in Go. Strings can be created by enclosing a set of characters inside double quotes " ".

		Strings in Go are Unicode compliant and are UTF-8 Encoded.  (very Important)

		Accessing individual bytes of a string
		Since a string is a slice of bytes, it’s possible to access each byte of a string. ,

		%s is the format specifier to print a string.
		%c format specifier is used to print the characters of the string in the printChars method.



		The reason is that the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1.
		We are trying to print characters assuming that each code point will be one byte long which is wrong.
		In UTF-8 encoding a code point can occupy more than 1 byte. So how do we solve this? This is where rune saves us.

		We are trying to print the characters of Señor and it outputs S e Ã ± o r which is wrong
		The reason is that the Unicode code point of ñ is U+00F1 and its UTF-8 encoding occupies 2 bytes c3 and b1.
		We are trying to print characters assuming that each code point will be one byte long which is wrong.
		In UTF-8 encoding a code point can occupy more than 1 byte. So how do we solve this? This is where rune saves us.


		Rune
		A rune is a builtin type in Go and it’s the alias of int32.
		Rune represents a Unicode code point in Go. It doesn’t matter how many bytes the code point occupies,
		it can be represented by a rune.

		Accessing individual runes using for range loop
		The above program is a perfect way to iterate over the individual runes of a string.
		But Go offers us a much easier way to do this using the for range loop.


		Creating a string from a slice of bytes

		String length
		The RuneCountInString(s string) (n int) function of the utf8 package can be used to find the length of the string.
		This method takes a string as an argument and returns the number of runes in it.

		len(s)
		is used to find the number of bytes in the string and it doesn’t return the string length.

		As we already discussed, some Unicode characters have code points that occupy more than 1 byte.

		The above output confirms that len(s) and RuneCountInString(s) return different values 😀.

		String comparison
		The == operator is used to compare two strings for equality. If both the strings are equal, then the result is true else it’s false.

		String concatenation
		There are multiple ways to perform string concatenation in Go. Let’s look at a couple of them.
			1. The most simple way to perform string concatenation is using the + operator.
			2/ using the Sprintf function of the fmt package.

		Strings are immutable
		Strings are immutable in Go. Once a string is created it’s not possible to change it.
		Strings are immutable in Go - once created, their contents cannot be changed
		Any valid unicode character within single quote is a rune

		To workaround this string immutability, strings are converted to a slice of runes.
		Then that slice is mutated with whatever changes are needed and converted back to a new string.



	*/

	word := "abcdefghijklmnopqrstuvwxyz"
	for _, v := range word {
		fmt.Printf("%v ", v)
	}
	fmt.Println()

	for _, v := range word {
		fmt.Printf("%v ", string(v))
	}

	name := "Hello World"
	fmt.Printf("String: %s\n", name)
	printBytes(name)
	fmt.Println()
	printChars(name)
	fmt.Printf("\n\n")
	name = "Señor"
	fmt.Printf("String: %s\n", name)
	printChars(name)
	fmt.Printf("\n")
	printBytes(name)
	fmt.Printf("\n")
	printCharsByRunes(name)
	fmt.Printf("\n")
	charsAndBytePosition(name)

	byteSclice := []byte{0x43, 0x61, 0x66, 0xC3, 0xA9}
	str := string(byteSclice)
	fmt.Println(str)

	runeSlice := []rune{0x0053, 0x0065, 0x00f1, 0x006f, 0x0072}
	str = string(runeSlice)
	fmt.Println(str)

	stringLength()

	str1 := "hi"
	str2 := "Hi"
	compareStrings(str1, str2)
	str3 := str1 + " " + str2
	fmt.Println(str3)
	str4 := fmt.Sprintf("%s %s", str1, str2)
	fmt.Println(str4)
}

func printBytes(s string) {
	fmt.Printf("Bytes: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%x ", s[i])
	}

}

func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}

func printCharsByRunes(s string) {
	fmt.Printf("Characters: ")
	runes := []rune(s) //convert

	for i := 0; i < len(runes); i++ {
		fmt.Printf("%c ", runes[i])
	}
}

func charsAndBytePosition(s string) {
	for index, rune := range s {
		fmt.Printf("%c starts at byte %d\n", rune, index)
	}
}

func stringLength() {
	word := "Señor"
	fmt.Printf("String: %s\n", word)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word))
	fmt.Printf("Number of bytes: %d\n", len(word))

	fmt.Printf("\n")
	word2 := "Pets"
	fmt.Printf("String: %s\n", word2)
	fmt.Printf("Length: %d\n", utf8.RuneCountInString(word2))
	fmt.Printf("Number of bytes: %d\n", len(word2))
}

func compareStrings(str1 string, str2 string) {
	if str1 == str2 {
		fmt.Printf("%s and %s are equal\n", str1, str2)
		return
	}
	fmt.Printf("%s and %s are not equal\n", str1, str2)
}

func mutate(s []rune) string {
	s[0] = 'a' //any valid unicode character within single quote is a rune
	//cannot assign to s[0] (neither addressable nor a map index expression)co

	return string(s)
}
