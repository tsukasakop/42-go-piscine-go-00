package piscine

import "ft"

func IsIncreasingN(i, n int) bool {
	prev := 10
	for i > 0 {
		if i % 10 >= prev {
			break
		}
		prev = i % 10
		i /= 10
		n--
	}
	return n == 0
}

func PrintPosNbr(n int) {
	if n / 10 > 0 {
		PrintPosNbr(n/10)
	}
	ft.PrintRune('0' + int32(n % 10))
}

func PrintCombN(n int) {
	lim := 1
	isFirst := true
	for i := 0; i < n; i++ {
		lim *= 10
	}
	for i :=0; i<lim; i++ {
		if !IsIncreasingN(i, n) {
			continue
		}
		if isFirst {
			isFirst = false
		} else {
			ft.PrintRune(',')
			ft.PrintRune(' ')
		}
		PrintPosNbr(i)
	}
	ft.PrintRune('\n')
}
