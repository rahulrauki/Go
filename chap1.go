package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func exer1() string {
	s, sep := "", ""
	for _, arg := range os.Args {
		s += sep + arg
		sep = " "
	}
	return s
	// fmt.Println(s)
}

func exer2() {
	for idx, arg := range os.Args {
		fmt.Println(idx, arg)
	}
}

func exer3() string {
	final := strings.Join(os.Args, " ")
	return final
	// fmt.Println(strings.Join(os.Args, " "))
}

func exer3BenchmarkTime() {
	// String addition
	start := time.Now()
	var str1, str2 string
	for i := 0; i < 1000000; i++ {
		str1 = exer1()
	}
	additionTIme := time.Since(start).Seconds()
	// String join
	start = time.Now()
	for i := 0; i < 1000000; i++ {
		str2 = exer3()
	}
	joinTime := time.Since(start).Seconds()

	fmt.Println("Addition time", additionTIme)
	fmt.Println("Join time", joinTime)
	fmt.Println(str1 + str2)

}

func main() {
	// Function calls
	exer1()
	exer2()
	exer3BenchmarkTime()
}
