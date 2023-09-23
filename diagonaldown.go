package Correction

import (
	"bufio"
	"fmt"
	"os"
)

func DiagonalDown(diagonal []rune) string {
	word := ""
	lineStr := ""
	for _, i := range diagonal {
		lineStr += string(i)
	}
	readFile, err := os.Open("words.txt")

	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		diagonalWord := fileScanner.Text()
		if len(diagonalWord) == len(lineStr) && (lineStr[0:len(diagonalWord)] == diagonalWord) {
			word += lineStr[0:len(diagonalWord)]
		} else {
			for i := 0; i < len(lineStr)-len(diagonalWord)+1; i++ {
				for j := 0; j < len(diagonalWord); j++ {
					if lineStr[i] == diagonalWord[j] {
						if lineStr[i:i+len(diagonalWord)] == diagonalWord {
							word += lineStr[i : i+len(diagonalWord)]
						}
					}
				}
			}
		}
	}
	return word
}