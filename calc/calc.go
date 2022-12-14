package calc

func Calc(a, b int, sign string) int {
	switch sign {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	default:
		return a / b
	}
}
