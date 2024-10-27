package main

import (
	"fmt"
)

// Function to reverse a string
func reverseString(s string) string {
	// Use a slice to reverse the string
	runes := []rune(s) // Convert string to a slice of runes to support unicode characters
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i] // Swap characters
	}
	return string(runes) // Convert the slice of runes back to a string
}

func main() {
	// Example usage of the function
	original := "halo dunia"
	reversed := reverseString(original)
	fmt.Println("String asli: ", original)
	fmt.Println("String terbalik: ", reversed)
}
