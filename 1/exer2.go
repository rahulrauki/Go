package main

import (
	"fmt"
	"os"
)

func PrintInNewLineTest() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}
}
