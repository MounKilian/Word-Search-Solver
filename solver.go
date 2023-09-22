package Correction

import "github.com/01-edu/z01"

func Solver(field [10][10]rune) {
	for _, i := range field {
		z01.PrintRune('[')
		word := Vertical(i)
		for _, j := range word {
			z01.PrintRune(j)
		}
		z01.PrintRune(']')
	}
}

/// Mots Trouvés =
