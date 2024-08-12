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
	data, err := parseInput(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	a := strings.Trim(data[0], "\"")
	operator := data[1]
	b := strings.Trim(data[2], "\"")

	if len(a) > 10 {
		fmt.Println("Ошибка: строка не должна содержать более 10 символов.")
		return
	}

	var result string

	switch operator {
	case "+":
		result = sumStrings(a, b)
	case "-":
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
		if num > len(a) {
			fmt.Println("Ошибка: длина строки меньше делителя.")
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

func parseInput(input string) ([]string, error) {
	parts := make([]string, 0, 3)

	start := strings.Index(input, "\"")
	end := strings.Index(input[start+1:], "\"") + start + 1
	if start == -1 || end == -1 || end <= start {
		return nil, fmt.Errorf("Некорректный формат ввода. Ожидается строка в кавычках.")
	}
	parts = append(parts, input[start:end+1])

	rest := strings.TrimSpace(input[end+1:])
	if len(rest) < 2 {
		return nil, fmt.Errorf("Некорректный формат ввода. Ожидается оператор.")
	}
	operator := string(rest[0])
	if operator != "+" && operator != "-" && operator != "*" && operator != "/" {
		return nil, fmt.Errorf("Некорректный ввод операции.")
	}
	parts = append(parts, operator)
	rest = strings.TrimSpace(rest[1:])

	if strings.HasPrefix(rest, "\"") {
		start = 0
		end = strings.Index(rest[1:], "\"") + 1
		if end <= start {
			return nil, fmt.Errorf("Некорректный формат ввода. Ожидается строка в кавычках.")
		}
		parts = append(parts, rest[start:end+1])
	} else {
		parts = append(parts, rest)
	}

	return parts, nil
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
