package main

import (
	"bufio"
	// "calc"
	cN "convertNum"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func inputUser() []string {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите мат.выражение: ")
	in.Scan()
	mathExpression := strings.Split(in.Text(), " ")
	return mathExpression
}
func arabOrRom(str string) (num int, typeNum string) {
	num, err := strconv.Atoi(str)
	if err == nil {
		typeNum = "arabic"
		return num, typeNum
	}
	if err != nil {
		num = cN.ToArabic(str)
		typeNum = "rom"
		return num, typeNum
	}
	return
}
func main() {
	fmt.Println(arabOrRom("G"))
	// mathExpression := inputUser()
	// if len(mathExpression) < 3 {
	// 	fmt.Println("Строка не является математической операцией!")
	// 	return
	// }
	// if len(mathExpression) > 3 {
	// 	fmt.Println("Формат мат.операции не удовлетворяет заданию")
	// 	return
	// }
	// aStr := mathExpression[0]
	// bStr := mathExpression[2]
	// a, b := arabOrRom(aStr, bStr)
	// sign := mathExpression[1]
	// if (a < 1 || a > 10) || (b < 1 || b > 10) {
	// 	fmt.Println("Одно из чисел не в допустимом диапазоне")
	// 	return
	// }
	// signs := [4]string{"+", "-", "*", "/"}
	// signsBool := false
	// for _, v := range signs {
	// 	if sign == v {
	// 		signsBool = true
	// 		break
	// 	}
	// }
	// if !signsBool {
	// 	fmt.Println("Нет такой мат.операции")
	// 	return
	// }
	// fmt.Println(calc.Calc(a, b, sign))
}
