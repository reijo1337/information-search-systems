package main

import "fmt"

func main() {
	t, p := "ababacaababaca", "ababaca"
	fmt.Println(nativeStringMatcher(t, p))
	fmt.Println(rabinKarpMathcer(t, p))
	fmt.Println(finiteAutomationMatcher(t, p))
	fmt.Println(kmpMatcher(t, p))
}
