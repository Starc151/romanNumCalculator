package main

import (
	"bufio"
	"calc"
	cN "convertNum"
	"errors"
	"fmt"
	"os"
	"quit"
	"strconv"
	"strings"
)

func getInput() []string {
	signs := [4]string{"+", "-", "*", "/"}
	sign := ""
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите мат.выражение: ")
	in.Scan()
	text := in.Text()
	text = strings.ReplaceAll(text, " ", "")
	for _, v := range signs {
		if strings.Contains(text, v) {
			sign = v
			text = strings.Replace(text, sign, " "+sign+" ", 1)
		}
	}
	mathExpression := strings.Split(text, " ")
	if len(mathExpression) != 3 {
		quit.Quit(errors.New("Строка не является мат.операцией"))
	}
	return mathExpression
}
func main() {
	mathExpression := getInput()
	aStr := mathExpression[0]
	bStr := mathExpression[2]
	if strings.Contains(aStr, ".") || strings.Contains(bStr, ".") {
		quit.Quit(errors.New("Калькулятор умеет работать только с целыми числами"))
	}
	a, aType := cN.ArabicOrRom(mathExpression[0])
	b, bType := cN.ArabicOrRom(mathExpression[2])
	if aType != bType {
		quit.Quit(errors.New("Используются разные системы счисления"))
	}
	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		quit.Quit(errors.New("Одно из чисел не в допустимом диапазоне"))
	}
	sign := mathExpression[1]
	res := ""
	resTmp := calc.Calc(a, b, sign)
	res = strconv.Itoa(resTmp)
	if aType == "rom" && bType == "rom" {
		res = cN.ToRoman(resTmp)
	}
	fmt.Println(res)
}
