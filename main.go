package main

import (
	"bufio"
	"calc"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func convertToRoman(num int) {
	nums := []struct {
		rom  string
		arab int
	}{
		// Больше 100 числане нужны по условию задачи (10*10=100)
		// {"M", 1000},
		// {"CM", 900},
		// {"D", 500},
		// {"CD", 400},
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

	res := ""
	for _, conversion := range nums {
		for num >= conversion.arab {
			res += conversion.rom
			num -= conversion.arab
		}
	}
	fmt.Println(res)
}

func inputUser() []string {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите мат.выражение: ")
	in.Scan()
	mathExpression := strings.Split(in.Text(), " ")
	return mathExpression
}

func main() {
	res := 0
	mathExpression := inputUser()
	a, _ := strconv.Atoi(mathExpression[0])
	b, _ := strconv.Atoi(mathExpression[2])
	sign := mathExpression[1]
	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		fmt.Println("Одно из чисел не в допустимом диапазоне")
		return
	}
	signs := [4]string{"+", "-", "*", "/"}
	signsBool := false
	for _, v := range signs {
		if sign == v {
			signsBool = true
			break
		}
	}
	if !signsBool {
		fmt.Println("Нет такой мат.операции")
		return
	}
	res = calc.Calc(a, b, sign)
	fmt.Println(res)
}
