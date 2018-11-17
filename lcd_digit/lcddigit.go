package lcddigit

import (
	"strconv"
	"strings"
)

func getLcdDigit(num int) string {
	mapNumStr := map[int]string{
		0: " _ \n| |\n|_|",
		1: "   \n  |\n  |",
		2: " _ \n _|\n|_ ",
		3: " _ \n _|\n _|",
		4: "   \n|_|\n  |",
		5: " _ \n|_ \n _|",
		6: " _ \n|_ \n|_|",
		7: " _ \n  |\n  |",
		8: " _ \n|_|\n|_|",
		9: " _ \n|_|\n  |",
	}
	output := mapNumStr[num]
	return output
}

// GetLcdDigit is function for getting all LCD digit
func GetLcdDigit(nums int) string {
	strNums := strconv.Itoa(nums)
	rowSlice := make([]string, 3)
	for i, c := range strNums {
		char, _ := strconv.Atoi(string(c))
		wordSlice := strings.Split(getLcdDigit(char), "\n")
		for j := 0; j < 3; j++ {
			if i < len(strNums)-1 {
				rowSlice[j] = rowSlice[j] + wordSlice[j] + " "
			} else {
				rowSlice[j] = rowSlice[j] + wordSlice[j]
			}

		}
	}
	var finalText string
	for i := 0; i < len(rowSlice); i++ {
		if i < len(rowSlice)-1 {
			finalText = finalText + rowSlice[i] + "\n"
		} else {
			finalText = finalText + rowSlice[i]
		}

	}
	return finalText
}
