package main

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

var variantOperation = map[string]func(numbers []float64) float64{
	"AVG": calculateAvg,
	"SUM": calculateSum,
	"MED": calculateMed,
}

func main() {
	operation := getUserInputOperation()
	numbers := getUserInputNumber()

	result, err := operationWithNumbers(numbers, operation)

	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}

func getUserInputOperation() string {
	var operaition string
	for {
		fmt.Print("Введите необходимую операцию (AVG - среднее, SUM - сумму, MED - медиану): ")
		fmt.Scan(&operaition)

		_, err := checkOperation(operaition)

		if err != nil {
			fmt.Println(err)
			continue
		}

		return strings.ToUpper(operaition)
	}
}

func checkOperation(operation string) (bool, error) {
	if strings.ToUpper(operation) == "AVG" || strings.ToUpper(operation) == "SUM" || strings.ToUpper(operation) == "MED" {
		return true, nil
	}

	return false, errors.New("Введена недоступная операция. Повторите ввод")
}

func getUserInputNumber() string {
	var numbers string
	fmt.Print("Введите числа через запятую:")
	fmt.Scan(&numbers)

	return numbers
}

func operationWithNumbers(numbers string, operation string) (float64, error) {
	arrNumbers := strings.Split(numbers, ",")
	var arrForCalc []float64

	for _, value := range arrNumbers {
		number, err := strconv.ParseFloat(strings.TrimSpace(value), 64)
		if err != nil {
			return 0, errors.New("Введено некорректное число")
		}
		arrForCalc = append(arrForCalc, number)
	}

	if len(arrForCalc) == 0 {
		return 0, errors.New("Не введено чисел")
	}

	return variantOperation[operation](arrForCalc), nil
}

func calculateAvg(numbers []float64) float64 {
	result := 0.0

	for _, value := range numbers {
		result += value
	}

	return result / float64(len(numbers))
}

func calculateSum(numbers []float64) float64 {
	result := 0.0

	for _, value := range numbers {
		result += value
	}
	return result
}

func calculateMed(numbers []float64) float64 {
	arrLenght := len(numbers)
	sort.Float64s(numbers)

	if arrLenght%2 == 1 {
		return numbers[arrLenght/2]
	}

	return (numbers[arrLenght/2-1] + numbers[arrLenght/2]) / 2
}
