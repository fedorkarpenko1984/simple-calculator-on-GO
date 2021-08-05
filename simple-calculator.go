package main

import (
	"fmt"
	"os"
)

func simpleCalc() {
	var firstNum float64
	var secondNum float64
	var operand string
	var repeat string

	fmt.Println("Введите, подтверждая каждый ввод клавишей Enter:")
	fmt.Println("Первое число")
	fmt.Println("Оператор в формате:")
	fmt.Println(" + : сложение")
	fmt.Println(" - : вычитание")
	fmt.Println(" * : умножение")
	fmt.Println(" / : деление")
	fmt.Println("Второе число")
	fmt.Fscan(os.Stdin, &firstNum)
	fmt.Fscan(os.Stdin, &operand)
	fmt.Fscan(os.Stdin, &secondNum)

	switch operand {
	case "+":
		fmt.Println("Результат")
		fmt.Println(firstNum + secondNum)
	case "-":
		fmt.Println("Результат")
		fmt.Println(firstNum - secondNum)
	case "*":
		fmt.Println("Результат")
		fmt.Println(firstNum * secondNum)
	case "/":
		fmt.Println("Результат")
		fmt.Println(firstNum / secondNum)
	default:
		fmt.Println("Некорректный оператор")
	}

	fmt.Println("Еще посчитаем? y - согласие")
	fmt.Fscan(os.Stdin, &repeat)

	if repeat == "y" || repeat == "Y" {
		simpleCalc()
	} else {
		fmt.Println("Пока-пока")
	}

}

func main() {
	fmt.Println("Здравствуйте!")
	simpleCalc()
}
