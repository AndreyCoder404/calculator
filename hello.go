package main

import (
	"fmt"
)

var romanToArabicMap = map[rune]int{
	'I': 1, 'V': 5, 'X': 10, 'L': 50, 'C': 100, 'D': 500, 'M': 1000,
}

func romanToInt(s string) int {
	var total int
	n := len(s)
	for i := 0; i < n; i++ {
		value := romanToArabicMap[rune(s[i])]
		if i+1 < n && value < romanToArabicMap[rune(s[i+1])] {
			total -= value
		} else {
			total += value
		}
	}
	return total
}

func intToRoman(num int) string {
	if num <= 0 {
		panic("Результат работы меньше единицы для римских чисел")
	}
	val := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	syms := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	roman := ""
	for i := 0; i < len(val); i++ {
		for num >= val[i] {
			num -= val[i]
			roman += syms[i]
		}
	}
	return roman
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Деление на ноль")
		}
		return a / b
	default:
		panic("Неизвестный оператор")
	}
}

func main() {
	for {
		var operand1, operator, operand2 string

		fmt.Print("Ввести: ")
		fmt.Scanf("%s %s %s", &operand1, &operator, &operand2)

		var result int

		if isRoman(operand1) && isRoman(operand2) {
			a := romanToInt(operand1)
			b := romanToInt(operand2)
			result = calculate(a, b, operator)
			fmt.Printf("Output: %s\n", intToRoman(result))
		} else {
			var a, b int
			_, err1 := fmt.Sscanf(operand1, "%d", &a)
			_, err2 := fmt.Sscanf(operand2, "%d", &b)

			if err1 == nil && err2 == nil {
				result = calculate(a, b, operator)
				fmt.Printf("Output: %d\n", result)
			} else {
				fmt.Println("Ошибка: Используются одновременно разные системы счисления или некорректный ввод")
			}
		}
	}
}

func isRoman(s string) bool {
	for _, ch := range s {
		if _, ok := romanToArabicMap[ch]; !ok {
			return false
		}
	}
	return true
}
