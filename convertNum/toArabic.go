package convertNum

import "strings"

func ToArabic(num string) int {
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
	arabic := 0
	for _, conversion := range nums {
		for i := len(num); i > 0; i-- {
			if strings.HasPrefix(num, conversion.rom) {
				arabic += conversion.arabic
				num = strings.Replace(num, conversion.rom, "", 1)
				continue
			}
		}
	}
	return arabic
}
