/*
	a more powerful hello world applications
	which shows off a little :P

	LESSON
		how to accept (*args style) command line argumanets
*/

package main

import (
	"os"		// order of import doesn't matter
	"fmt"
	"strings"
)

func main() {

	who:="World"
	if len(os.Args) > 1 && true {	// if statement doesn't require '()'
		who = strings.Join(os.Args[1:], " ")
	}
	fmt.Println("Hello", who)

}