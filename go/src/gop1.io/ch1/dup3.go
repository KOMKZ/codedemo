package main

import (
	"os"
	"io/ioutil"
	"fmt"
	"strings"
)

func main()  {
	files := os.Args[1:]
	for _,filePath := range files{
		data,err := ioutil.ReadFile(filePath)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v \n", err)
		}
		for _, line := range strings.Split(string(data), "\n"){
			fmt.Println(line)
		}
	}
}