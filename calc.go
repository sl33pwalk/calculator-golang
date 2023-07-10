package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func romantic(x string) int {

	arabic, _ := strconv.Atoi(x)

	if _, err := strconv.Atoi(x); err == nil {
		return arabic
	}

	rome := 0
	for i := (len(x) - 1); i >= 0; i-- {
		if string(x[i]) == "I" {
			rome += 1
		}

		if string(x[i]) == "V" {
			rome += 5
		}

		if string(x[i]) == "X" {
			rome += 10
		}

		if (i + 1) <= (len(x) - 1) {
			if string(x[i+1]) == "V" || string(x[i+1]) == "X" {
				rome -= 2
			}
		}
	}

	return rome
}

func expression(x int, operator string, y int) int {

	switch op := operator; op {
	case "+":
		result := x + y
		return result

	case "-":
		result := x - y
		return result
	case "*":
		result := x * y
		return result
	case "/":
		result := x / y
		return int(result)
	}
	return 1
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func integerToRoman(number int) string {
	maxRomanNumber := 3999
	if number > maxRomanNumber {
		return strconv.Itoa(number)
	}

	conversions := []struct {
		value int
		digit string
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

	var roman strings.Builder
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman.WriteString(conversion.digit)
			number -= conversion.value
		}
	}

	return roman.String()
}

func whatNumSystem(a string) string {
	if _, err := strconv.Atoi(a); err == nil {
		return "arabic"
	}
	return "roman"
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Введите выражение")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		splitted := strings.Fields(text)

		if len(splitted) != 3 {
			fmt.Println("Неверный формат ввода")
			return
		}

		if whatNumSystem(splitted[0]) != whatNumSystem(splitted[2]) {
			fmt.Println("Должна использоваться только одна система счисления")
			return
		}

		if romantic(splitted[0]) <= 10 &&
			romantic(splitted[0]) >= 1 &&
			romantic(splitted[2]) <= 10 &&
			romantic(splitted[2]) >= 1 &&
			(romantic(splitted[2])%1) == 0 &&
			(romantic(splitted[0])%1) == 0 {

			if expression(romantic(splitted[0]),
				splitted[1],
				romantic(splitted[2])) != 1 {

				if _, err := strconv.Atoi(splitted[0]); err == nil {
					if _, err := strconv.Atoi(splitted[2]); err == nil {
						fmt.Println(expression(romantic(splitted[0]),
							splitted[1],
							romantic(splitted[2])))
					}
				} else {
					if expression(romantic(splitted[0]),
						splitted[1],
						romantic(splitted[2])) > 0 {
						fmt.Println(integerToRoman(expression(romantic(splitted[0]),
							splitted[1],
							romantic(splitted[2]))))
					} else {
						fmt.Println("результат работы меньше единицы")
						return
					}
				}

			} else {
				fmt.Println("Неправильно введён арифметический оператор")
				os.Exit(1)
			}

		} else {
			fmt.Println("Калькулятор принимает на вход только числа от 1 до 10 включительно")
			os.Exit(1)
		}

	}

}
