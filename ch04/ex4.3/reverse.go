package reverse

func reverse(s *[3]int) {
	j := len(s) - 1
	for i := 0; i < len(s)/2; i++ {
		s[i], s[j-i] = s[j-i], s[i]
	}
}
