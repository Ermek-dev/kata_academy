package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romanToArabic(roman string) int {
	romanValues := map[rune]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	var result int
	var prevValue int

	for _, char := range roman {
		value := romanValues[char]
		result += value
		if prevValue < value {
			result -= 2 * prevValue
		}
		prevValue = value
	}

	return result
}

func isRomanNumber(str string) bool {
	romanChars := "IVXLCDM"
	for _, char := range str {
		if !strings.ContainsRune(romanChars, char) {
			return false
		}
	}
	return true
}

func isValidArabicNumber(num int) bool {
	return num >= 1 && num <= 10
}

func isValidRomanNumber(roman string) bool {
	validRomanNumbers := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for _, validNum := range validRomanNumbers {
		if roman == validNum {
			return true
		}
	}
	return false
}

func add(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		return a + b.(int)
	case string:
		result := romanToArabic(a) + romanToArabic(b.(string))
		if result < 1 {
			panic("Результат меньше единицы. Римские числа не могут быть отрицательными.")
		}
		return result
	}
	return nil
}

func subtract(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		return a - b.(int)
	case string:
		result := romanToArabic(a) - romanToArabic(b.(string))
		if result < 1 {
			panic("Результат меньше единицы. Римские числа не могут быть отрицательными.")
		}
		return result
	}
	return nil
}

func multiply(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		return a * b.(int)
	case string:
		result := romanToArabic(a) * romanToArabic(b.(string))
		if result < 1 {
			panic("Результат меньше единицы. Римские числа не могут быть отрицательными.")
		}
		return result
	}
	return nil
}

func divide(a, b interface{}) interface{} {
	switch a := a.(type) {
	case int:
		return a / b.(int)
	case string:
		result := romanToArabic(a) / romanToArabic(b.(string))
		if result < 1 {
			panic("Результат меньше единицы. Римские числа не могут быть отрицательными.")
		}
		return result
	}
	return nil
}

func arabicToRoman(arabic string) string {
	arabicNum, err := strconv.Atoi(arabic)
	if err != nil {
		panic("Ошибка при конвертации арабского числа в римское")
	}
	if !isValidArabicNumber(arabicNum) {
		panic("Неверное арабское число. Поддерживаемые значения: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10")
	}
	romanValues := []struct {
		Value int
		Symbol string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}
	romanNum := ""
	for _, rv := range romanValues {
		for arabicNum >= rv.Value {
			romanNum += rv.Symbol
			arabicNum -= rv.Value
		}
	}
	return romanNum
}

func main() {
	fmt.Println("====================================================")
	fmt.Println("Приветствую дорогой пользователь :)\n Добро пожаловать в игру калькулятор")
	fmt.Println("====================================================")
	fmt.Println("Введите выражение в формате \"a операция b\" :\n")
	fmt.Println("Подсказка: Тобиш,для коректной работы вам нужно ввести числа и операнд через пробел(например, 1 + 4):\n>")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	// Разбиваем введенную строку на операнды и операцию
	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic("Выдача паники.Неверный формат ввода. Пожалуйста, введите выражение в формате \"a операция b\"")
	}

	// Проверяем, что оба операнда имеют одинаковый формат (либо оба арабские, либо оба римские)
	if (isRomanNumber(parts[0]) && !isRomanNumber(parts[2])) || (!isRomanNumber(parts[0]) && isRomanNumber(parts[2])) {
		panic("Выдача паники.Калькулятор умеет работать только с арабскими или римскими цифрами одновременно")
	}

	var (
		a, b interface{}
		err  error
	)

	if isRomanNumber(parts[0]) {
		if !isRomanNumber(parts[2]) {
			panic("Выдача паники.Нельзя выполнять операции между арабскими и римскими числами одновременно")
		}
		if !isValidRomanNumber(parts[0]) || !isValidRomanNumber(parts[2]) {
			panic("Выдача паники.Введены неверные римские числа. Поддерживаемые значения: I, II, III, IV, V, VI, VII, VIII, IX, X")
		}
		a = parts[0]
		b = parts[2]
	} else {
		a, err = strconv.Atoi(parts[0])
		if err != nil || !isValidArabicNumber(a.(int)) {
			panic("Выдача паники.Введено неверное арабское число. Поддерживаемые значения: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10")
		}
		b, err = strconv.Atoi(parts[2])
		if err != nil || !isValidArabicNumber(b.(int)) {
			panic("Выдача паники.Введено неверное арабское число. Поддерживаемые значения: 1, 2, 3, 4, 5, 6, 7, 8, 9, 10")
		}
	}

	var result interface{}
	switch parts[1] {
	case "+":
		result = add(a, b)
	case "-":
		result = subtract(a, b)
	case "*":
		result = multiply(a, b)
	case "/":
		if b.(int) == 0 {
			panic("Выдача паники.Деление на ноль невозможно")
		}
		result = divide(a, b)
	default:
		panic("Выдача паники.Неподдерживаемая операция. Допустимые операции: +, -, *, /")
	}

	switch result := result.(type) {
	case int:
		fmt.Println("Ваш результат:", result)
	case string:
		fmt.Println("Ваш результат:", result)
	}
}
