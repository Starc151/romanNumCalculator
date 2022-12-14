package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertToRoman(num int) {
	nums := map[string]int{
		// Больше 100 числане нужны по условию задачи (10*10=100)
		// "M":  1000,
		// "CM": 900,
		// "D":  500,
		// "CD": 400,
		"C":  100,
		"XC": 90,
		"L":  50,
		"XL": 40,
		"X":  10,
		"IX": 9,
		"V":  5,
		"IV": 4,
		"I":  1,
	}
	fmt.Println(nums)
	// res := ""
	// for _, conversion := range nums {
	//     for num >= conversion.value {
	//         res += conversion.digit
	//         num -= conversion.value
	//     }
	// }
	// return res

}

func inputUser() []string {
	in := bufio.NewScanner(os.Stdin)
	fmt.Print("Введите мат.выражение: ")
	in.Scan()
	mathExpression := strings.Split(in.Text(), " ")
	return mathExpression
}

func main() {
	convertToRoman(5)
}
