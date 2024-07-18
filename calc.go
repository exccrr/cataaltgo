package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	data := strings.Fields(input)

	if len(data) != 3 {
		fmt.Println("Некорректный формат ввода. Ожидается 3 части: строка, оператор, строка/число.")
		return
	}

	if _, err := strconv.Atoi(data[0]); err == nil {
		fmt.Println("Первое значение должно быть строкой в кавычках.")
		return
	}

	if !strings.HasPrefix(data[0], "\"") || !strings.HasSuffix(data[0], "\"") {
		fmt.Println("Выделите двойными кавычками строковые значения.")
		return
	}

	a := strings.Trim(data[0], "\"")
	operator := data[1]
	b := data[2]

	var result string

	switch operator {
	case "+":
		if !strings.HasPrefix(b, "\"") || !strings.HasSuffix(b, "\"") {
			fmt.Println("Выделите двойными кавычками строковые значения.")
			return
		}
		b = strings.Trim(b, "\"")
		result = sumStrings(a, b)
	case "-":
		if !strings.HasPrefix(b, "\"") || !strings.HasSuffix(b, "\"") {
			fmt.Println("Выделите двойными кавычками строковые значения.")
			return
		}
		b = strings.Trim(b, "\"")
		result = subtractStrings(a, b)
	case "*":
		num, err := strconv.Atoi(b)
		if err != nil || num < 1 || num > 10 {
			fmt.Println("Введите число от 1 до 10.")
			return
		}
		result = multiplyStringByNumber(a, num)
	case "/":
		num, err := strconv.Atoi(b)
		if err != nil || num < 1 || num > 10 {
			fmt.Println("Введите число от 1 до 10.")
			return
		}
		result = divideStringByNumber(a, num)
	default:
		fmt.Println("Некорректный ввод операции.")
		return
	}

	const maxLength = 40
	if len(result) > maxLength {
		fmt.Println(result[:maxLength] + "...")
	} else {
		fmt.Println("Результат:", result)
	}
}

func sumStrings(a, b string) string {
	return a + b
}

func subtractStrings(a, b string) string {
	return strings.Replace(a, b, "", 1)
}

func divideStringByNumber(a string, b int) string {
	if b == 0 {
		return ""
	}
	return a[:len(a)/b]
}
func multiplyStringByNumber(a string, b int) string {
	result := ""
	for i := 0; i < b; i++ {
		result += a
	}
	return result
}