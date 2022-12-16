package convertNum

import (
	"strconv"
)

func ArabicOrRom(str string) (int, string) {
	num, err := strconv.Atoi(str)
	typeNum := ""
	if err == nil {
		typeNum = "arabic"
	} else {
		num = ToArabic(str)
		typeNum = "rom"
	}
	return num, typeNum
}
