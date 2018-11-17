package lcddigit

import (
	"testing"
)

func TestPrintLcdDigit_1(t *testing.T) {
	r := GetLcdDigit(1)
	expect := "   \n  |\n  |"
	if r != expect {
		t.Errorf("expect \n%s\n got \n%s\n", expect, r)
	}
}

func TestPrintLcdDigit_69(t *testing.T) {
	r := GetLcdDigit(69)
	expect := " _   _ \n|_  |_|\n|_|   |"
	if r != expect {
		t.Errorf("expect \n%s\n got \n%s\n", expect, r)
	}
}

// func ExamplePrintLcdDigit() {
// 	r := GetLcdDigit(0)
// 	fmt.Println(r)
// Output:  _
// | |
// |_|
// }
