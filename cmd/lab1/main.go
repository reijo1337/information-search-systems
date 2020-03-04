package main

import "fmt"

type testCase struct {
	desc  string
	array []int
}

func main() {
	testCases := []testCase{
		testCase{
			desc:  "sorted 10",
			array: makeSorted(10),
		},
		testCase{
			desc:  "sorted 20",
			array: makeSorted(20),
		},
		testCase{
			desc:  "sorted 40",
			array: makeSorted(40),
		},
		testCase{
			desc:  "revers sorted 10",
			array: makeRevSorted(10),
		},
		testCase{
			desc:  "revers sorted 20",
			array: makeRevSorted(20),
		},
		testCase{
			desc:  "revers sorted 40",
			array: makeRevSorted(40),
		},
		testCase{
			desc:  "unsorted 10",
			array: makeRandom(10),
		},
		testCase{
			desc:  "unsorted 20",
			array: makeRandom(20),
		},
		testCase{
			desc:  "unsorted 40",
			array: makeRandom(40),
		},
	}
	for _, tC := range testCases {
		runTest(tC)
	}
}

func runTest(tc testCase) {
	fmt.Println()
	fmt.Println(tc.desc)
	bubble := make([]int, len(tc.array))
	copy(bubble, tc.array)
	insertion := make([]int, len(tc.array))
	copy(insertion, tc.array)
	tree := make([]int, len(tc.array))
	copy(tree, tc.array)
	quick := make([]int, len(tc.array))
	copy(quick, tc.array)

	s, c := bubbleSort(bubble)
	fmt.Printf("bubble\n\t\tswap: %d\tcomprasions: %d\n", s, c)
	s, c = insertionSort(insertion)
	fmt.Printf("insertion\n\t\tswap: %d\tcomprasions: %d\n", s, c)
	s, c = treeSort(tree)
	fmt.Printf("tree\n\t\tswap: %d\tcomprasions: %d\n", s, c)
	s, c = quickSort(quick)
	fmt.Printf("quick\n\t\tswap: %d\tcomprasions: %d\n", s, c)
}
