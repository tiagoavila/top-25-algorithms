package euclideanfindgcd

func FindGCD(a, b int) int {
	if b == 0 {
		return a
	}

	return FindGCD(b, a%b)
}
