package main

import "fmt"

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("abc" + appendTest("aa"))
}

func appendTest(tes string) string {
	return "aa"
}
