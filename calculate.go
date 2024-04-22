package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Math operations
func Add(x, y int) int {
	return x + y
}

func Sub(x, y int) int {
	return x - y
}

func Mul(x, y int) int {
	return x * y
}

func Div(x, y int) int {
	return x / y
}

func ParseInt(x, y string) (int, int) {
	num1, _ := strconv.Atoi(x)
	num2, _ := strconv.Atoi(y)
	return num1, num2
}

func RomanToInt(s string) int {
	h := map[uint8]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}
	var v, lv, cv int
	for i := len(s) - 1; i >= 0; i-- {
		cv = h[s[i]]
		if cv < lv {
			v -= cv
		} else {
			v += cv
		}
		lv = cv
	}
	return v
}

func IntToRoman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
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
	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func Looperkal(x, plus, y string) {
	inted := [10]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	roman := [10]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	for i := 0; i < len(inted); i++ {
		for j := 0; j < len(inted); j++ {
			if x == inted[i] && y == inted[j] {
				num1, num2 := ParseInt(x, y)
				switch plus {
				case "+":
					fmt.Println(Add(num1, num2))
				case "-":
					fmt.Println(Sub(num1, num2))
				case "*":
					fmt.Println(Mul(num1, num2))
				case "/":
					fmt.Println(Div(num1, num2))
				default:
					fmt.Println("Error: не соответствует математической операции")
				}
			} else if x == roman[i] && y == roman[j] {
				num1 := RomanToInt(x)
				num2 := RomanToInt(y)
				switch plus {
				case "+":
					rom := Add(num1, num2)
					fmt.Println(IntToRoman(rom))
				case "-":
					rom := Sub(num1, num2)
					if rom < 1 {
						fmt.Print("Error: в римской системе нет отрицательных чисел и нуля")
					}
					fmt.Println(IntToRoman(rom))
				case "*":
					rom := Mul(num1, num2)
					fmt.Println(IntToRoman(rom))
				case "/":
					rom := Div(num1, num2)
					fmt.Println(IntToRoman(rom))
				default:
					fmt.Println("Error: не соответствует математической операции")
				}
			}
			if (x == roman[i] && y == inted[j]) || (x == inted[i] && y == roman[j]) {
				fmt.Println("Error: разные типы счисления")
			}
		}
	}
}

func main() {
	var x, plus, y string
	var line string
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line = sc.Text()
	arr := strings.Split(line, " ")
	if len(arr) != 3 {
		fmt.Println("Error: математическая операция должна состоять из двух элементов")
		return
	}
	x, plus, y = arr[0], arr[1], arr[2]
	x1, y1 := ParseInt(x, y)
	x2 := RomanToInt(x)
	y2 := RomanToInt(y)
	if x1 > 10 || y1 > 10 || x1 < 0 || y1 < 0 {
		fmt.Println("Error: числа находятся в диапозоне от 1 до 10")
	}
	if x2 > 10 || y2 > 10 || x2 < 0 || y2 < 0 {
		fmt.Println("Error: числа находятся в диапозоне от 1 до 10")
	}
	Looperkal(x, plus, y)
}
