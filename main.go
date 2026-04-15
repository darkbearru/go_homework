package main

import (
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"
)

const ErrorInvalidOperation = "...Неверное название операции"
const ErrorInvalidNumbers = "...Неверный перечень чисел"

var operations = []string{"AVG", "SUM", "MED"}

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	var numbers = make([]int, 0, 10)
	operation := userOperationInput("Выбери операцию:")
	numbers = readNumbers("Введи список чисел:", numbers)

	fmt.Println("Операция:", operation)
	fmt.Println("Числа:", numbers)

	switch operation {
	case "AVG":
		fmt.Println("Среднее:", operationAVG(numbers))
	case "SUM":
		fmt.Println("Сумма:", operationSUM(numbers))
	case "MED":
		fmt.Println("Медиана:", operationMED(numbers))
	}
}

func userOperationInput(message string) string {
	var operation string
	for {
		fmt.Printf(message+"(%s)\n", strings.Join(operations, ", "))
		_, err := fmt.Scan(&operation)
		if err != nil {
			fmt.Println(ErrorInvalidOperation)
			continue
		}
		operation = strings.ToUpper(operation)
		if slices.Index(operations, operation) == -1 {
			fmt.Println(ErrorInvalidOperation)
			continue
		}
		break
	}
	return operation
}

func readNumbers(message string, numbers []int) []int {
	var numbersList string
	var list []string
	fmt.Println(message)
	for {
		_, err := fmt.Scan(&numbersList)
		if err != nil {
			fmt.Println(ErrorInvalidNumbers)
			continue
		}
		list = strings.Split(numbersList, ",")
		if len(list) == 0 {
			fmt.Println(ErrorInvalidNumbers)
			continue
		}
		for _, number := range list {
			num, err := strconv.Atoi(number)
			if err != nil {
				fmt.Println(ErrorInvalidNumbers)
				continue
			}
			numbers = append(numbers, num)
		}
		break
	}
	return numbers
}

func operationAVG(numbers []int) float64 {
	sum := operationSUM(numbers)
	return float64(sum) / float64(len(numbers))
}

func operationSUM(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func operationMED(numbers []int) float64 {
	count := len(numbers)
	arr := make([]int, count)
	copy(arr, numbers)

	sort.Ints(arr)

	if count%2 == 1 {
		return float64(arr[count/2])
	}

	return float64(arr[count/2-1]+arr[count/2]) / 2.0
}
