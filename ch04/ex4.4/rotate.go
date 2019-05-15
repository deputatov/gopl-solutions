package rotate

func rotate(q []int) {
	first := q[0]
	copy(q, q[1:])
	q[len(q)-1] = first
}
