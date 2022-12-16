package main

import (
	"bufio"
	"calc"
	cN "convertNum"
	"fmt"
	"os"
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
	return mathExpression
}

func main() {
	mathExpression := getInput()
	if len(mathExpression) < 3 {
		fmt.Println("Строка не является мат.операцией!")
		return
	}
	if len(mathExpression) > 3 {
		fmt.Println("Формат мат.операции не удовлетворяет заданию")
		return
	}
	aStr := mathExpression[0]
	bStr := mathExpression[2]
	if strings.Contains(aStr, ".") || strings.Contains(bStr, ".") {
		fmt.Println("Калькулятор умеет работать только с целыми числами.")
		return
	}
	a, aType := cN.ArabicOrRom(mathExpression[0])
	b, bType := cN.ArabicOrRom(mathExpression[2])
	if aType == "err" || bType == "err" {
		fmt.Println("Строка не является мат.операцией!")
		return
	}
	if aType != bType {
		fmt.Println("Используются одновременно разные системы счисления")
		return
	}
	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		fmt.Println("Одно из чисел не в допустимом диапазоне")
		return
	}
	sign := mathExpression[1]
	res := ""
	resTmp := calc.Calc(a, b, sign)
	res = strconv.Itoa(resTmp)
	if aType == "rom" && bType == "rom" {
		res = cN.ToRoman(resTmp)
	}
	if res == "negative" {
		fmt.Println("В римской СС нет 0 и отрицательных чисед")
		return
	}
	fmt.Println(res)
}
