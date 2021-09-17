package piscine

func Atoi(s string) int {
	PosOrNeg := 1
	returnedValue := 0
	var newString string
	for _, char := range s {
		ascii := int(char)
		if ascii == 43 || ascii == 45 || (ascii > 47 && ascii < 58) {
			newString += string(ascii)
		} else {
			return 0
		}
	}
	for index, char := range newString {
		ascii := int(char)
		if ascii == 45 && index == 0 {
			PosOrNeg = 0
		} else if ascii == 43 && index == 0 {
			PosOrNeg = 1
		} else if ascii > 47 && ascii < 58 {
			returnedValue = returnedValue*10 + (ascii - 48)
		} else {
			return 0
		}
	}
	if PosOrNeg == 0 {
		returnedValue = returnedValue * -1
	}
	return returnedValue
}
