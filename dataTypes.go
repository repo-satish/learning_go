/*
	FROM:
		Learning-Go-latest.pdf
	SHA1:
		0c67e186bd68154b986060583848ae64cef5652c
	SIZE:
		995036 B
*/

package main

import "fmt"

func main() {

	const FOUR, FIVE = 4, 5
	var (
		a int	//generic type
		b int32	//specific 32bit int type
	)

	a = 10
	// b = a + a	//-->> ERROR: illegal mixing, ".\dataTypes.go:15: cannot use a + a (type int) as type int32 in assignment"
	b = b + FOUR
	b = b + 5





	const (
		ZERO = iota	// iota being called for the first time, hence it is 0
		ONE = iota	// second time iota is being called -therefore- its value now becomes 1
		TWO		// *******     ← Implicitly TWO = iota; let Go repeat the use of '= iota'
	)
	// fmt.Printf("ZERO is %d ONE is %d TWO is %d !!!", ZERO, ONE, TWO)





	var s string
	s = "Hello World"
	// s[0] = 'B'	// ERROR: strings are immutable in Go, to do that
	chars := []rune(s)	// ***** RUNE is alias for int32. It is an UTF-8 encoded code point, e.g. when iterating over string in C we use char datatype but go is unicode lover so we use int32 or rune to iterate over unicode strings thereby ensuring 6-byte long japanese chars don't get read as 1-byte garbage values
	chars[0] = 'B'
	s2 := string(chars)
	fmt.Printf("%s", s2)

	var s3 string
	/*s3 = "Starting part"
		+ "Ending part"		ERROR - because go automatically inserts ';' at end of every line so the '+ "Ending part"' becomes invalid syntax, to overcome this*/
	s3 = "Starting part" +	// use + at end of line so go won't insert ';' at end
		"Ending part"
	s3 = `
		Starting
		Ending
	`// OR use backticks to make RAW strings literals -- WARNING------  s3 contains newline char """Unlike interpreted string literals the value of a raw string literal is composed of the uninterpreted characters between the quotes"""





	var (
		small complex64	// 32bit real and complex parts
		big complex128
	)
	small = 5 + 5i
	fmt.Printf("Complex number displayed as: %v", small)

}