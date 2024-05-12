package main

import (
	"bufio"
	"flag"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func input(numbers *[]int) {
	scanner := bufio.NewScanner(os.Stdin) //Сканнер для отслеживания ввода в консоль
	for scanner.Scan() {
		input := scanner.Text()
		if input == ";" || input == "" {
			break
		}
		number, err := strconv.Atoi(input)
		if err != nil {
			println("Invalid input")
			continue
		}
		if number > 100000 || number < -100000 {
			println("The number cannot be greater than 100000 or less than -100000")
			continue
		}
		*numbers = append(*numbers, number)
	}
}

func mean(arr []int) float64 {
	sum := 0
	for j := 0; j < len(arr); j++ {
		sum += arr[j]
	}
	return float64(sum) / float64(len(arr))
}

func median(arr []int) float64 {
	sort.Ints(arr)
	if len(arr)%2 == 1 {
		return float64(arr[len(arr)/2])
	} else {
		return (float64(arr[len(arr)/2]) + float64(arr[len(arr)/2-1])) / 2
	}
}

func mode(arr []int) int {
	freq := make(map[int]int)
	for _, num := range arr {
		if v, ok := freq[num]; ok {
			freq[num] = v + 1
		} else {
			freq[num] = 1
		}
	}
	bestNum := arr[0]
	bestTimes := freq[bestNum]
	for num, times := range freq {
		if times > bestTimes || (times == bestTimes && num < bestNum) {
			bestTimes = times
			bestNum = num
		}
	}
	return bestNum
}

func sd(arr []int) float64 {
	var total float64
	mean := mean(arr)
	for _, num := range arr {
		total += math.Pow(float64(num)-mean, 2)
	}
	variance := total / float64(len(arr))
	res := math.Sqrt(variance)
	return res
}

func main() {
	fmt.Println("Для остановки ввода впишите знак ';' или пустую строку\nВведите число:")

	var meanFlag, medianFlag, modeFlag, sdFlag bool
	var numbers []int

	flag.BoolVar(&meanFlag, "mean", false, "Mean")
	flag.BoolVar(&medianFlag, "median", false, "Median")
	flag.BoolVar(&modeFlag, "mode", false, "Mode")
	flag.BoolVar(&sdFlag, "sd", false, "SD")

	flag.Parse() //Парсинг флагов из коммандной строки

	if meanFlag == false && medianFlag == false && modeFlag == false && sdFlag == false {
		meanFlag, medianFlag, modeFlag, sdFlag = true, true, true, true
	} //Если не было введено опций, выводятся все показатели

	input(&numbers)

	if meanFlag {
		fmt.Printf("Mean: %.2f\n", mean(numbers))
	}
	if medianFlag {
		fmt.Printf("Median: %.2f\n", median(numbers))
	}
	if modeFlag {
		fmt.Printf("Mode: %d\n", mode(numbers))
	}
	if sdFlag {
		fmt.Printf("SD: %.2f\n", sd(numbers))
	}
}
