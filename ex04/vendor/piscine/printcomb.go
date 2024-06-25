package piscine

import "ft"

func PrintComb() {
	for l := '0'; l <= '7'; l++ {
		for m := l + 1; m <= '8'; m++ {
			for r := m + 1; r <= '9'; r++ {
				ft.PrintRune(l)
				ft.PrintRune(m)
				ft.PrintRune(r)
				if l != '7' {
					ft.PrintRune(',')
					ft.PrintRune(' ')
				}
			}
		}
	}
	ft.PrintRune('\n')
}
