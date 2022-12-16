package score

func Score(i int) int {
	if i < 21 {
		return 0
	} else if i == 21 {
		return 1
	} else {
		return 2
	}
}
