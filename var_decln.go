/*
	FROM:
		Learning-Go-latest.pdf
	SHA1:
		0c67e186bd68154b986060583848ae64cef5652c
	SIZE:
		995036 B
*/

package main

func main() {

	/*first method, name it then assign a value*/
	var a int
	a = -15

	/*second method, take a variable name and assign it a value in the same statement with ":=" operator*/
	b := "a b c"	// go decides the data-type by itself

	/*for multiple declarations of variable(or const or import) following syntax may be used*/
	var (
		c int
		d float64
	)
	c, d = 3, 4.1	// Yup, tuple assignment works here

	_, __, g := 1, 2, 3	// NOTE: '_' is discarded variable but '__' is valid variable name

	const (
		ONE int32 = 1		// if '= 1' part is skipped and declared after ')' then its an error, strange -- cuz non-persistant syntax
		TWO string = "2"
	)
	//TWO, ONE = 2, 1
	const THREE, FOUR = 3, 4
}