package piscine

import "ft"

func PrintComb2() {
	var l, r int32
	for l=0;l <= 98; l++ {
		for r= l+1; r <= 99; r++ {
			ft.PrintRune('0' + l / 10)
			ft.PrintRune('0' + l % 10)
			ft.PrintRune(' ')
			ft.PrintRune('0' + r / 10)
			ft.PrintRune('0' + r % 10)
			if l != 98 {
				ft.PrintRune(',')
				ft.PrintRune(' ')
			}
		}
	}
	ft.PrintRune('\n')
}
