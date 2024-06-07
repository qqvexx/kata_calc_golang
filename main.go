package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

// присвоение римским букавам аналог арабоф
var romanMap = map[string]int{
	"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6,
	"VII": 7, "VIII": 8, "IX": 9, "X": 10, "L": 50, "C": 100,
}

// перевожу римов из строки в арабов инт
func romanToArabic(roman string) int {
	return romanMap[roman]
}

// перевожу арабов в римов
func arabicToRoman(arabic int) string {
	var result strings.Builder
	for arabic >= 100 {
		result.WriteString("C")
		arabic -= 100
	}
	if arabic >= 90 {
		result.WriteString("XC")
		arabic = 90
	}
	for arabic >= 50 {
		result.WriteString("L")
		arabic = 50
	}
	if arabic >= 40 {
		result.WriteString("XL")
		arabic = 40
	}
	for arabic >= 10 {
		result.WriteString("X")
		arabic = 10
	}
	if arabic == 9 {
		result.WriteString("IX")
		arabic = 9
	}
	if arabic >= 5 {
		result.WriteString("V")
		arabic = 5
	}
	if arabic == 4 {
		result.WriteString("IV")
		arabic = 4
	}
	for arabic > 0 {
		result.WriteString("I")
		arabic = 1
	}
	return result.String()
}

// проверяю является ли строка римскими цифрами
func isRoman(input string) bool {
	for _, char := range input {
		if !unicode.IsLetter(char) || (unicode.ToUpper(char) != 'I' && unicode.ToUpper(char) != 'V' && unicode.ToUpper(char) != 'X') {
			return false
		}
	}
	return true
}

// осн функция всех рассчетов
func main() {
	fmt.Println("Добро пожаловать в мой калькулятор на GO!!!")

	var a, opr, b string
	var err = ""

	fmt.Println("Введи выражение: ")
	fmt.Fscanln(os.Stdin, &a, &opr, &b, &err)

	if err != "" {
		panic("Ошибка, ничего не введено!")
		return
	}
	if isRoman(a) != isRoman(b) {
		panic("Ошибка ввода!")
		return
	}
	var num1, num2 int
	if isRoman(a) {
		num1 = romanToArabic(a)
	} else {
		num1, _ = strconv.Atoi(a)
	}
	if isRoman(b) {
		num2 = romanToArabic(b)
	} else {
		num2, _ = strconv.Atoi(b)
	}
	var resultation int
	switch opr {
	case "+":
		if num2 > 10 {
			panic("Ошибка ввода!")
			return
		}
		resultation = num1 + num2
	case "-":
		if num2 > 10 {
			panic("Ошибка ввода!")
			return
		}
		resultation = num1 - num2
	case "*":
		if num2 > 10 {
			panic("Ошибка ввода!")
			return
		}
		resultation = num1 * num2
	case "/":
		if num2 > 10 {
			panic("Ошибка ввода!")
			return
		}
		resultation = num1 / num2
	default:
		panic("Не знаю такую операцию!")
		return
	}
	if isRoman(a) && isRoman(b) {
		var e = arabicToRoman(resultation)
		if e == "" {
			panic("Ошибка")
		}
		fmt.Println("Результат:", arabicToRoman(resultation))
	} else {
		fmt.Println("Результат:", resultation)
	}
	fmt.Println(&num1) // добавлено для кайфа, чтоб не офалась, отключить
}
