package wordfreq

import (
	"bufio"
	"fmt"
	"os"
)

func wordfreq() map[string]int{
	counts := make(map[string]int)
	file, err := os.Open("text")
	if err != nil {
		fmt.Fprintf(os.Stderr, "file: %v\n", err)
	}
	countWords(file, counts)
	file.Close()
	return counts
}

func countWords(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
