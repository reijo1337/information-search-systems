package main

func bubble(a []int) (int, int) {
	swap, comp := 0, 0
	for i := 0; i < len(a)-1; i++ {
		for j := 0; j < len(a)-i-1; j++ {
			if a[j] > a[j+1] {
				swap++
				comp++
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}
	return swap, comp
}

func insertion(a []int) (int, int) {
	swap, comp := 0, 0
	for j := 1; j < len(a); j++ {
		key := a[j]
		i := j - 1
		for ; i >= 0 && a[i] > key; i-- {
			swap++
			comp++
			a[i+1] = a[i]
		}
		a[i+1] = key
		swap++
	}
	return swap, comp
}
