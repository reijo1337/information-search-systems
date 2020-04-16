package main

func kmpMatcher(T, P string) (out []int) {
	n := len(T)
	m := len(P)
	pi := computePrefixFunction(P)
	q := -1

	for i := 0; i < n; i++ {
		for q > -1 && P[q+1] != T[i] {
			q = pi[q]
		}

		if P[q+1] == T[i] {
			q++
		}

		if q == m-1 {
			out = append(out, i-m+1)
			q = pi[q]
		}
	}

	return
}

func computePrefixFunction(P string) []int {
	m := len(P)
	pi := make([]int, m)
	k := -1
	pi[0] = -1
	for q := 1; q < m; q++ {
		for k > -1 && P[k] != P[q] {
			k = pi[k]
		}
		if P[k+1] == P[q] {
			k++
		}
		pi[q] = k
	}
	return pi
}
