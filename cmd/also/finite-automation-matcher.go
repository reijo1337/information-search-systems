package main

const alphabet = "1234567890-=~!@#$%^&*()_+|QWERTYUIOP[]qwertyuiop{}ASDFGHJKL;'asdfghjkl:ZXCVBNM,./zxcvbnm<>? "

// const alphabet = "abc"

// Построение конечного автомата. Как в книжке не вышло, посмотрел в другом месте
func computeTF(P string) []map[byte]int {
	m := len(P)
	out := make([]map[byte]int, m+1)
	for i := range out {
		out[i] = make(map[byte]int)
	}

	for state := 0; state <= m; state++ {
		for i := range alphabet {
			x := alphabet[i]
			out[state][x] = getNextState(P, state, x)
		}
	}

	return out
}

func getNextState(P string, state int, x byte) int {
	m := len(P)

	// Если x совпадает с символом P[state], то значит нам нужно перейти на следующее состояние (очевидно)
	if state < m && x == P[state] {
		return state + 1
	}

	// ns - nextState
	// ищем максимально возможный ns такой, чтоб P[:ns] был суфиксом P[:state-1]x
	for ns := state; ns > 0; ns-- {
		if P[ns-1] == x {
			i := 0
			for ; i < ns-1; i++ {
				if P[i] != P[state-ns+1+i] {
					break
				}
			}
			if i == ns-1 {
				return ns
			}
		}
	}

	return 0
}

func finiteAutomationMatcher(T, P string) (out []int) {
	o := computeTF(P)

	// for i := range o {
	// 	fmt.Printf("%d:\t", i)
	// 	for j := range alphabet {
	// 		fmt.Printf("%d\t", o[i][alphabet[j]])
	// 	}
	// 	fmt.Println()
	// }

	n := len(T)
	q := 0
	for i := 1; i <= n; i++ {
		q = o[q][T[i-1]]
		if q == len(P) {
			out = append(out, i-len(P))
		}
	}
	return
}
