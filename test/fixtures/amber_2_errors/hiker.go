package hiker

func answer() int {
	return 6 * multiplier()
}

func answerDigitSum() int {
	return digitSum(answer())
}
