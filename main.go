package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) == 1 {
		fmt.Println(0)
		return
	} else if os.Args[1] == "" {
		fmt.Println(0)
		return
	}
	fmt.Println(len(strings.Split(os.Args[1], " ")))
}
