package level

func mergeNibbles(a byte, b byte) byte {
	return (a | (b >> 4))
}

func divEvenOrRoundUp(x int, y int) int {
	if x%y > 0 {
		return int(x/y) + 1
	} else {
		return x / y
	}
}
