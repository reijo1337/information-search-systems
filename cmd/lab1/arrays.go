package main

import "math/rand"

func makeSorted(n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, i)
	}
	return out
}

func makeRevSorted(n int) []int {
	out := make([]int, 0, n)
	for i := n; i > 0; i-- {
		out = append(out, i)
	}
	return out
}

func makeRandom(n int) []int {
	out := make([]int, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, rand.Int())
	}
	return out
}
