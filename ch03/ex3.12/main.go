package main

import (
	"bytes"
	"fmt"
	"sort"
)

func isAnagram(firstWord, secondWord string) bool {

	b1 := []byte(firstWord)
	b2 := []byte(secondWord)

	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })

	return bytes.EqualFold(b1, b2)

}

func main() {

	var firstWord, secondWord string
	fmt.Print("Input two words : ")
	fmt.Scan(&firstWord, &secondWord)

	fmt.Println(isAnagram(firstWord, secondWord))

}
