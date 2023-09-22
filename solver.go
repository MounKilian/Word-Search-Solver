package Correction

import "github.com/01-edu/z01"

func Solver(field [10][10]rune) {
	/// Cherche tout les mots horizontalement
	resultHorizontal := []string{}
	for _, i := range field {
		word := Horizontal(i)
		if word != "" {
			resultHorizontal = append(resultHorizontal, word)
		}
	}
	/// On print tout les mots qu'on a trouvé horizontalement
	z01.PrintRune('[')
	for _, j := range resultHorizontal {
		for _, k := range j {
			z01.PrintRune(k)
		}
		if j != resultHorizontal[len(resultHorizontal)-1] {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
}

/// Mots Trouvés =
