package main

func bmMatcher(T, P string) (out []int) {
	n := len(T)
	m := len(P)
	lambda := computeLastOccurrenceFunction(P)
	gamma := computeGoodSuffixFunction(P)
	s := 0
	for s <= n-m {
		j := m - 1
		for j >= 0 && P[j] == T[s+j] {
			j--
		}
		if j == -1 { // Полное совпадение
			out = append(out, s)
			s += gamma[0]
		} else { // Работа эвристик
			plus := gamma[j] // эвристика безопасного суффикса
			if plus > j-lambda[T[s+j]] {
				plus = j - lambda[T[s+j]] // эвристика стоп символа
			}
			s += plus
		}
	}
	return
}

// эвристика стоп символа
func computeLastOccurrenceFunction(P string) map[byte]int {
	lambda := make(map[byte]int, len(alphabet))
	for i := range alphabet {
		a := alphabet[i]
		lambda[a] = -1
	}
	for j := range P {
		lambda[P[j]] = j
	}
	return lambda
}

// эвристика безопасного суффикса
func computeGoodSuffixFunction(P string) []int {
	pi := computePrefixFunction(P)
	pReverse := reverse(P)
	piReverse := computePrefixFunction(pReverse)

	m := len(P)
	gamma := make([]int, 0, m+1)
	for j := 0; j <= m; j++ {
		gamma = append(gamma, m-pi[m-1]-1)
	}
	for l := 0; l < m; l++ {
		j := m - piReverse[l] - 1
		if gamma[j] > l-piReverse[l] {
			gamma[j] = l - piReverse[l]
		}
	}

	return gamma
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
