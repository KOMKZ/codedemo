package main

import (
	"os"
	"fmt"
	"bufio"
)

func main()  {
	files := os.Args[1:]
	for _,filePath := range files {
		f, err:= os.Open(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err)
			continue
		}
		input := bufio.NewScanner(f)
		for input.Scan() {
			fmt.Println(input.Text())
		}
		f.Close()
	}
}