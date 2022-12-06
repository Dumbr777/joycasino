package cards

import (
	"fmt"
)

func CardValue(who, s string, score int) int {

	cartmun := 0
	switch s {
	case "двойка":
		cartmun = 2
	case "тройка":
		cartmun = 3
	case "четверка":
		cartmun = 4
	case "пятерка":
		cartmun = 5
	case "шестерка":
		cartmun = 6
	case "семерка":
		cartmun = 7
	case "восьмерка":
		cartmun = 8
	case "девятка":
		cartmun = 9
	case "десятка":
		cartmun = 10
	case "валет":
		cartmun = 10
	case "дама":
		cartmun = 10
	case "король":
		cartmun = 10
	case "туз":
		if who == "Дилер" {
			if score < 11 {
				cartmun = 11
			} else {
				cartmun = 1
			}
		} else {
			var input int
			fmt.Printf("Выберете 1 или 11, Ваш счет: %v\n", score)
			fmt.Scanln(&input)
			if input == 1 {
				cartmun = 1
			} else if input == 11 {
				cartmun = 11
			}
		}
	}
	return cartmun
}
