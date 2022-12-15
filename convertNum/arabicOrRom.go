package convertNum

import "strconv"

func ArabicOrRom(str string) (num int, typeNum string) {
	num, err := strconv.Atoi(str)
	if err == nil {
		typeNum = "arabic"
		return num, typeNum
	} else {
		num = ToArabic(str)
		typeNum = "rom"
		// В римской СС нет 0
		if num == 0 {
			typeNum = "err"
		}
		return num, typeNum
	}
}
