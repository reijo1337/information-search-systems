package main

// Просто ходим вдоль исходной строки от 0 до len(t)-len(p) и каждый раз сравниваем
func nativeStringMatcher(t, p string) (out []int) {
	n := len(t)
	m := len(p)
	for s := 0; s <= n-m; s++ {
		if p == t[s:s+m] {
			out = append(out, s)
		}
	}
	return
}
