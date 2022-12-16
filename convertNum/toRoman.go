package convertNum

import (
	"errors"
	"quit"
)

func ToRoman(num int) string {
	if num <= 0 {
		quit.Quit(errors.New("В римской СС нет 0 и отрицательных чисед"))
	}
	nums := []struct {
		rom    string
		arabic int
	}{
		{"M", 1000},
		{"CM", 900},
		{"D", 500},
		{"CD", 400},
		{"C", 100},
		{"XC", 90},
		{"L", 50},
		{"XL", 40},
		{"X", 10},
		{"IX", 9},
		{"V", 5},
		{"IV", 4},
		{"I", 1},
	}

	rom := ""
	for _, conversion := range nums {
		for num >= conversion.arabic {
			rom += conversion.rom
			num -= conversion.arabic
		}
	}
	return rom
}
