package main

import (
	"fmt"
)

func main() {
	// Place your code here.
	length := readInt()

	if length <= 0 {
		length = 8
	}
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			printChessRow(length, "#", " ")
		} else {
			printChessRow(length, " ", "#")
		}
	}
}

func readInt() int {
	var length int
	fmt.Scan(&length)
	return length
}

func printChessRow(length int, firstSymbol string, secondSymbol string) {
	for i := 0; i < length; i++ {
		if i%2 == 0 {
			fmt.Print(firstSymbol)
		} else {
			fmt.Print(secondSymbol)
		}
	}
	fmt.Println()
}
