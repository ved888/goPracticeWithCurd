package main

import "fmt"

func main() {
	fmt.Println("basic pattern=======")
	BasicPattern()
	fmt.Println("RightAngledTriangle=======")
	RightAngledTriangle()
	fmt.Println("InvertedTriangle==========")
	InvertedTriangle()
	fmt.Println("NumberPattern============")
	NumberPattern()
	fmt.Println("alphabet pattern============")
	AlphabetPattern()
	fmt.Println("PyramidPattern=============")
	PyramidPattern()
	fmt.Println("HollowSquarePattern=============")
	HollowSquarePattern()
}

// *****
// *****
// *****
// *****
// *****
func BasicPattern() {
	str := "*"

	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			fmt.Print(str)
		}
		fmt.Println() // Move to the next line after each row
	}
}

func RightAngledTriangle() {
	str := "*"
	for i := 0; i < 7; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print(str)
		}
		fmt.Println()
	}
}

func InvertedTriangle() {
	str := "*"

	for i := 7; i >= 0; i-- {
		for j := 0; j < i; j++ {
			fmt.Print(str)
		}
		fmt.Println()
	}
}

func NumberPattern() {
	for i := 1; i <= 5; i++ {
		for j := 1; j <= i; j++ {
			fmt.Print(j)
		}
		fmt.Println()
	}
}

// 'A' is a rune in Go, which represents the ASCII value of the letter A (65).
// Adding j to 'A' increments the ASCII value, printing the next letter.
// %c in fmt.Printf() prints the character corresponding to the ASCII value.
func AlphabetPattern() {
	for i := 0; i < 5; i++ {
		for j := 0; j <= i; j++ {
			fmt.Printf("%c", 'A'+j) // Convert byte to character
		}
		fmt.Println()
	}
}

//	   *
//	  ***
//	 *****
//	*******
//
// *********
func PyramidPattern() {
	for i := 0; i < 7; i++ {
		for j := 0; j < 7-i-1; j++ {
			fmt.Print(" ")
		}
		for k := 0; k < 2*i+1; k++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

//	 *
//	***
//
// *****
//
//	***
//	 *
func DiamondPattern() {

}

// QUESTIONS
// Print a spiral matrix pattern.
// Print a hollow pyramid pattern.
// Print a butterfly pattern:
// *    *
// **  **
// ******
// **  **
// *    *

func HollowSquarePattern() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i == 0 || i == 5-1 || j == 0 || j == 5-1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
