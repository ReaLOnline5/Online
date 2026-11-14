package main

import "fmt"

func main() {
	// Задача 1: Проверка числа (положительное, отрицательное или ноль)
	var number int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)

	if number > 0 {
		fmt.Println("Number is positive")
	} else if number < 0 {
		fmt.Println("Number is negative")
	} else {
		fmt.Println("Number is zero")
	}

	fmt.Println()

	// Задача 2: Чётные числа с помощью цикла for
	var n int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	for i := 0; i <= n; i += 2 {
		fmt.Println(i)
	}
}
