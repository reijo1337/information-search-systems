package main

func bubbleSort(a []int) (int, int) {
	swap, comp := 0, 0
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			comp++
			if a[j] > a[j+1] {
				swap++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return swap, comp
}

func insertionSort(a []int) (int, int) {
	swap, comp := 0, 0
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		if !(i >= 0 && a[i] > key) {
			comp++
		}
		for ; i >= 0 && a[i] > key; i-- {
			comp++
			swap++
			a[i+1] = a[i]
		}
		a[i+1] = key
		swap++
	}
	return swap, comp
}

type tree struct {
	key   int
	left  *tree
	right *tree
}

func (t *tree) insert(key int) (comp int) {
	if t == nil {
		t = &tree{key: key}
		return
	}

	if key < t.key {
		comp = 1 + t.left.insert(key)
	} else {
		comp = 1 + t.right.insert(key)
	}
	return
}

func (t *tree) walk(a *[]int) {
	if t == nil {
		return
	}
	t.left.walk(a)
	*a = append(*a, t.key)
	t.right.walk(a)
}

func treeSort(a []int) (int, int) {
	swap, comp := len(a), 0
	var t *tree
	for _, key := range a {
		comp += t.insert(key)
	}
	a = a[:0]
	t.walk(&a)
	return swap, comp
}

func quickSort(a []int) (int, int) {
	return recursiveSort(a, 0, len(a)-1)
}

func recursiveSort(arr []int, start, end int) (swap int, comp int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	// Iterate sub array to find values less than pivot
	//   and move them to the beginning of the array
	//   keeping splitIndex denoting less-value array size
	for i := start; i < end; i++ {
		comp++
		if arr[i] < pivot {
			if splitIndex != i {
				swap++
				arr[splitIndex], arr[i] = arr[i], arr[splitIndex]
			}

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	swap++
	swap++

	s1, c1 := recursiveSort(arr, start, splitIndex-1)
	s2, c2 := recursiveSort(arr, splitIndex+1, end)
	return swap + s1 + s2, comp + c1 + c2
}
