package main

import "time"

func getCurrentTimeAsBluePrint() [6][7]bool {
	now := time.Now()

	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()

	h1 := int8(hour / 10)
	h2 := int8(hour % 10)
	m1 := int8(minute / 10)
	m2 := int8(minute % 10)
	s1 := int8(second / 10)
	s2 := int8(second % 10)

	return [6][7]bool{
		getDigitBluePrint(h1),
		getDigitBluePrint(h2),
		getDigitBluePrint(m1),
		getDigitBluePrint(m2),
		getDigitBluePrint(s1),
		getDigitBluePrint(s2),
	}
}

func getDigitBluePrint(digit int8) [7]bool {
	bluePrintMap := [10][7]bool{
		{true, true, true, false, true, true, true},     // 0
		{false, false, true, false, false, true, false}, // 1
		{true, false, true, true, true, false, true},    // 2
		{true, false, true, true, false, true, true},    // 3
		{false, true, true, true, false, true, false},   // 4
		{true, true, false, true, false, true, true},    // 5
		{true, true, false, true, true, true, true},     // 6
		{true, false, true, false, false, true, false},  // 7
		{true, true, true, true, true, true, true},      // 8
		{true, true, true, true, false, true, true},     // 9
	}

	return bluePrintMap[digit]
}
