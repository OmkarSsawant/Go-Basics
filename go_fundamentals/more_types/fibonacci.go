package moretypes

func fibonacci() func() int {
	prev, nxt := 0, 1
	return func() int {
		prev, nxt = nxt, prev+nxt
		return nxt - prev
	}
}
