package main

const d = uint64(16777619)

func rabinKarpMathcer(T, P string) (out []int) {
	n := len(T)
	m := len(P)
	h := uint64(1)
	for i := 0; i < m-1; i++ {
		h *= d
	}
	var p, t, tt uint64

	for i := 0; i < m; i++ {
		p = (d*p + uint64(P[i]))
		t = (d*t + uint64(T[i]))
		tt = (d*tt + uint64(T[i+1]))
	}

	for s := 0; s <= n-m; s++ {
		if p == t && P == T[s:s+m] {
			out = append(out, s)
		}
		if s < n-m {
			t = d*(t-(uint64(T[s]))*h) + uint64(T[s+m])
		}
	}

	return
}
