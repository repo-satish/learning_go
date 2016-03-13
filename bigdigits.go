package main

import (
	"os"
	"fmt"
	"log"
	"path/filepath"
)

func main() {

	// var longWeekend = []string{"Friday", "Saturday", "Sunday", "Monday"}
	// var lowPrimes = []int{2, 3, 5, 7, 11, 13, 17, 19}
	var bigDigits = [][]string{
		{"   000   ", " 0     0 ", "0       0", "0       0", "0       0", " 0     0 ", "   000   "},
		{"    1    ", "  1 1    ", "    1    ", "    1    ", "    1    ", "    1    ", "  11111  "},
		{"   222   ", "  2   2  ", "       2 ", "       2 ", "      2  ", "     2   ", "   22222 "},
		{"   333   ", "  3   3  ", "       3 ", "     33  ", "       3 ", "  3   3  ", "   333   "},
		{"    4    ", "   44    ", "  4  4   ", " 4   4   ", " 4444444 ", "     4   ", "     4   "},
		{"  55555  ", "  5      ", "  555    ", "     5   ", "      5  ", "     5   ", "  55     "},
		{"     6   ", "    6    ", "   6     ", "  6 6    ", " 6    6  ", " 6     6 ", "  66666  "},
		{"  77777  ", "      7  ", "      7  ", "     7   ", "    7    ", "    7    ", "    7    "},
		{"   888   ", "  8   8  ", " 8     8 ", "   888   ", " 8     8 ", "  8   8  ", "   888   "},
		{"    9999 ", "   9   9 ", "  9    9 ", "   99999 ", "       9 ", "       9 ", "      9  "},
	}

	if len(os.Args) == 1 {
		fmt.Printf("Usage %s <whole-number>\n", filepath.Base(os.Args[0]))
		os.Exit(1)
	}

	stringOfDigits := os.Args[1]
	for row := range bigDigits[0] {
		line := ""
		for column := range stringOfDigits {
			/*the follwoing line is HOT HOT HOT----
				remember C, char to int mapping, but why is it ok in Go,
				since rune should've been used???
			*/
			digit := /*rune(*/stringOfDigits[column]/*)*/ - /*rune(*/'0'/*)*/
			if 0 <= digit && digit <= 9 {
				line+= bigDigits[digit][row] + " "
			} else {
				log.Fatal("Invalid whole-number")
			}
		}
		fmt.Println(line)
	}

}