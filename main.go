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
	"i": 1, "ii": 2, "iii": 3, "iv": 4, "v": 5, "vi": 6,
	"vii": 7, "viii": 8, "ix": 9, "x": 10, "l": 50, "c": 100,
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
		arabic -= 90
	}
	for arabic >= 50 {
		result.WriteString("L")
		arabic -= 50
	}
	if arabic >= 40 {
		result.WriteString("XL")
		arabic -= 40
	}
	for arabic >= 10 {
		result.WriteString("X")
		arabic -= 10
	}
	if arabic == 9 {
		result.WriteString("IX")
		arabic -= 9
	}
	if arabic >= 5 {
		result.WriteString("V")
		arabic -= 5
	}
	if arabic == 4 {
		result.WriteString("IV")
		arabic -= 4
	}
	for arabic > 0 {
		result.WriteString("I")
		arabic -= 1
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

	var x, opr, y, empt string

	fmt.Println("Введи выражение: ")
	fmt.Fscanln(os.Stdin, &x, &opr, &y, &empt)

	if empt != "" {
		panic("Ошибка, ничего не введено!")
		return
	}
	if isRoman(x) != isRoman(y) {
		panic("Ошибка, нужно вводить выражение в одной системе счисления")
		return
	}

	var num1, num2 int
	if isRoman(x) {
		num1 = romanToArabic(x)
	} else {
		num1, _ = strconv.Atoi(x)
	}
	if isRoman(y) {
		num2 = romanToArabic(y)
	} else {
		num2, _ = strconv.Atoi(y)
	}

	// здесь обработка ввода пользователя, если нет в мапе такого значения то выключаемся
	num1, ok := romanMap[x]
	if !ok {
		panic("Нужно вводить римские цифры правильно!")
	}
	num2, oke := romanMap[y]
	if !oke {
		panic("Нужно вводить римские цифры правильно!")
	}
	
	var resultation int

	switch opr {
	case "+":
		if num2 > 10 || num2 < 1 && num1 > 10 || num1 < 1 {
			panic("Ошибка, я работаю с числами от 1 до 10")
			return
		}
		resultation = num1 + num2
	case "-":
		if num2 > 10 || num2 < 1 && num1 > 10 || num1 < 1 {
			panic("Ошибка, я работаю с числами от 1 до 10")
			return
		}
		resultation = num1 - num2
	case "*":
		if num2 > 10 || num2 < 1 && num1 > 10 || num1 < 1 {
			panic("Ошибка, я работаю с числами от 1 до 10")
			return
		}
		resultation = num1 * num2
	case "/":
		if num2 > 10 || num2 < 1 && num1 > 10 || num1 < 1 {
			panic("Ошибка, я работаю с числами от 1 до 10")
			return
		}
		resultation = num1 / num2
	default:
		panic("Не знаю такую операцию!")
		return
	}

	if isRoman(x) && isRoman(y) {
		var negative = arabicToRoman(resultation)
		if negative == "" {
			panic("Ошибка, в римской системе чисел нет отрицательных значений")
		}
		fmt.Println("Твой результат:", arabicToRoman(resultation))
	} else {
		fmt.Println("Твой результат:", resultation)
	}
}
