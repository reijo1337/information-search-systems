package main

import "fmt"

func main() {
	t, p := "000010001010001", "0001"
	fmt.Println(nativeStringMatcher(t, p))
	fmt.Println(rabinKarpMathcer(t, p))
}
