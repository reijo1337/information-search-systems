package main

import "fmt"

func main() {
	t, p := "ababacaababaca", "ababaca"
	fmt.Println(t, p)
	fmt.Println(nativeStringMatcher(t, p))
	fmt.Println(rabinKarpMathcer(t, p))
	fmt.Println(finiteAutomationMatcher(t, p))
	fmt.Println(kmpMatcher(t, p))
	fmt.Println(bmMatcher(t, p))
	t, p = "00001010110001001", "0001"
	fmt.Println(t, p)
	fmt.Println(nativeStringMatcher(t, p))
	fmt.Println(rabinKarpMathcer(t, p))
	fmt.Println(finiteAutomationMatcher(t, p))
	fmt.Println(kmpMatcher(t, p))
	fmt.Println(bmMatcher(t, p))
}
