package Correction

import (
	"github.com/01-edu/z01"
)

func Solver(field [10][10]rune) {
	/// Mots Trouvés =
	intro := "Mots Trouvés ="
	for _, o := range intro {
		z01.PrintRune(o)
	}
	z01.PrintRune('\n')
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
	/// Cherche tout les mots verticalement
	resultVertical := []string{}
	for x := 0; x < 10; x++ {
		var column []rune
		for y := 0; y < 10; y++ {
			column = append(column, field[y][x])
		}
		word := Vertical(column)
		if word != "" {
			resultVertical = append(resultVertical, word)
		}
	}
	/// On print tout les mots qu'on a trouvé verticalement
	z01.PrintRune('[')
	for _, j := range resultVertical {
		for _, k := range j {
			z01.PrintRune(k)
		}
		if j != resultVertical[len(resultVertical)-1] {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
	/// Cherche tout les mots en diagonal vers le bas
	resultDiagonalDown := []string{}
	for y := 9; y >= 0; y-- {
		var diagonal []rune
		for x := y; x < 10; x++ {
			diagonal = append(diagonal, field[y+(x-y)][x-y])
		}
		word := DiagonalDown(diagonal)
		if word != "" {
			resultDiagonalDown = append(resultDiagonalDown, word)
		}
	}
	for x := 9; x > 0; x-- {
		var diagonal []rune
		for y := x; y < 10; y++ {
			diagonal = append(diagonal, field[y-x][x+(y-x)])
		}
		word := DiagonalDown(diagonal)
		if word != "" {
			resultDiagonalDown = append(resultDiagonalDown, word)
		}
	}
	/// On print tout les mots qu'on a trouvé en diagonal vers le bas
	z01.PrintRune('[')
	for _, j := range resultDiagonalDown {
		for _, k := range j {
			z01.PrintRune(k)
		}
		if j != resultDiagonalDown[len(resultDiagonalDown)-1] {
			z01.PrintRune(',')
			z01.PrintRune(' ')
		}
	}
	z01.PrintRune(']')
}

/// Mots Trouvés =
