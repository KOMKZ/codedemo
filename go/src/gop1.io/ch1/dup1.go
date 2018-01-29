package main

import (
	"bufio"
	"os"
	"fmt"
)

func main() {
	fmt.Println("enter:")
	container := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		line := input.Text()
		container[line]++
		fmt.Printf(line)
	}

	for line,count := range container {
		fmt.Printf("%s %d", line, count)
	}
}