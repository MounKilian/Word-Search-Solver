package Correction

import (
	"github.com/01-edu/z01"
)

func Solver(field [10][10]rune) {
	/// Print du field
	z01.PrintRune('\n')
	for _, i := range field {
		z01.PrintRune('[')
		for _, j := range i {
			z01.PrintRune(j)
			z01.PrintRune(' ')
		}
		z01.PrintRune(']')
		z01.PrintRune('\n')
	}
	z01.PrintRune('\n')

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
		if len(word) != 0 {
			for _, i := range word {
				resultHorizontal = append(resultHorizontal, string(i))
			}
		}
	}
	/// On print tout les mots qu'on a trouvé horizontalement
	z01.PrintRune('[')
	if len(resultHorizontal) == 0 {
		z01.PrintRune(']')
	} else {
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

	/// Cherche tout les mots verticalement
	resultVertical := []string{}
	for x := 0; x < 10; x++ {
		var column []rune
		for y := 0; y < 10; y++ {
			column = append(column, field[y][x])
		}
		word := Vertical(column)
		if len(word) != 0 {
			for _, i := range word {
				resultVertical = append(resultVertical, string(i))
			}
		}
	}
	/// On print tout les mots qu'on a trouvé verticalement
	z01.PrintRune('[')
	if len(resultVertical) == 0 {
		z01.PrintRune(']')
	} else {
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
	}

	/// Cherche tout les mots en diagonal vers le bas
	resultDiagonalDown := []string{}
	for y := 9; y >= 0; y-- {
		var diagonal []rune
		for x := y; x < 10; x++ {
			diagonal = append(diagonal, field[y+(x-y)][x-y])
		}
		word := DiagonalDown(diagonal)
		if len(word) != 0 {
			for _, i := range word {
				resultDiagonalDown = append(resultDiagonalDown, string(i))
			}
		}
	}
	for x := 9; x > 0; x-- {
		var diagonal []rune
		for y := x; y < 10; y++ {
			diagonal = append(diagonal, field[y-x][x+(y-x)])
		}
		word := DiagonalDown(diagonal)
		if len(word) != 0 {
			for _, res := range word {
				resultDiagonalDown = append(resultDiagonalDown, string(res))
			}
		}
	}
	/// On print tout les mots qu'on a trouvé en diagonal vers le bas
	z01.PrintRune('[')
	if len(resultDiagonalDown) == 0 {
		z01.PrintRune(']')
	} else {
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

	/// Cherche tout les mots en diagonal vers le haut
	resultDiagonalUp := []string{}
	for y := 0; y < 10; y++ {
		var diagonal []rune
		for x := 0; x <= y; x++ {
			diagonal = append(diagonal, field[y-x][x])
		}
		word := diagonalUp(diagonal)
		if len(word) != 0 {
			resultDiagonalUp = append(resultDiagonalUp, word...)
		}
	}
	for x := 1; x < 10; x++ {
		var diagonal []rune
		for y := 9; y >= x; y-- {
			diagonal = append(diagonal, field[y][x+(9-y)])
		}
		word := diagonalUp(diagonal)
		if len(word) != 0 {
			resultDiagonalUp = append(resultDiagonalUp, word...)
		}
	}
	/// On print tout les mots qu'on a trouvé en diagonal vers le haut
	z01.PrintRune('[')
	if len(resultDiagonalUp) == 0 {
		z01.PrintRune(']')
	} else {
		for _, j := range resultDiagonalUp {
			for _, k := range j {
				z01.PrintRune(k)
			}
			if j != resultDiagonalUp[len(resultDiagonalUp)-1] {
				z01.PrintRune(',')
				z01.PrintRune(' ')
			}
		}
		z01.PrintRune(']')
	}
	z01.PrintRune('\n')
}
