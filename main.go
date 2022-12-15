package main

import (
	"bufio"
	"calc"
	// "convertNum"
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
func arabOrRom(aStr, bStr string) (a, b int) {
	a, errA := strconv.Atoi(aStr)
	b, errB := strconv.Atoi(bStr)
	if errA == nil && errB == nil {
		return a, b
	}
	return 100, 100
}
func main() {
	mathExpression := inputUser()
	aStr := mathExpression[0]
	bStr := mathExpression[2]
	a, b := arabOrRom(aStr, bStr)
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
	fmt.Println(calc.Calc(a, b, sign))
}
