package piscine

func ActiveBits(n int) int {
	if n < 2 {
		return int(n)
	}
	return (int(n) % 2) + ActiveBits(n/2)
}
