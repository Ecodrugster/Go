package main

import "fmt"

func main() {
	//1
	// arr := [10]int{1, 5, 8, 12, 3, 7, 14, 6, 9, 2}

	// fmt.Println("\n1. Элементы массива:")
	// for i := 0; i < len(arr); i++ {
	// 	fmt.Println(arr[i])
	// }
	//2
	// arr := [10]int{1, 5, 8, 12, 3, 7, 14, 6, 9, 2}
	// sum := 0

	// for _, v := range arr {
	// 	sum += v
	// }

	// fmt.Println("\n2. Сумма элементов:", sum)
	//3
	// arr := [10]int{1, 5, 8, 12, 3, 7, 14, 6, 9, 2}

	// min := arr[0]
	// max := arr[0]

	// for _, v := range arr {
	// 	if v < min {
	// 		min = v
	// 	}
	// 	if v > max {
	// 		max = v
	// 	}
	// }

	// fmt.Println("\n3. Минимум:", min)
	// fmt.Println("   Максимум:", max)
	//4
	// arr := [10]int{1, 5, 8, 12, 3, 7, 14, 6, 9, 2}
	// count := 0

	// for _, v := range arr {
	// 	if v%2 == 0 {
	// 		count++
	// 	}
	// }

	// fmt.Println("\n4. Количество чётных чисел:", count)
	//5
	// arr := [10]int{1, 5, 8, 12, 3, 7, 14, 6, 9, 2}

	// fmt.Println("\n5. Умножение каждого элемента на 2:")
	// for _, v := range arr {
	// 	fmt.Println(v * 2)
	// }
	//6
	// s := []int{1, 2, 3, 4, 5}
	// s = append(s, 6, 7, 8)

	// fmt.Println("\n6. Слайс после append:", s)
	//7
	// s := []int{10, 20, 30, 40, 50}
	// index := 2

	// s = append(s[:index], s[index+1:]...)

	// fmt.Println("\n7. Слайс после удаления:", s)
	//8
	// s := []int{10, 20, 30, 40, 50}
	// sum := 0

	// for _, v := range s {
	// 	sum += v
	// }

	// average := float64(sum) / float64(len(s))
	// fmt.Println("\n8. Среднее арифметическое:", average)
	//9
	// names := []string{"Алексей", "Иван", "Мария", "Оля", "Дмитрий"}

	// shortest := names[0]

	// for _, name := range names {
	// 	if len(name) < len(shortest) {
	// 		shortest = name
	// 	}
	// }

	// fmt.Println("\n9. Самое короткое имя:", shortest)
	//10
	// s := []int{5, 12, 7, 20, 3, 15, 9}
	// var result []int

	// for _, v := range s {
	// 	if v > 10 {
	// 		result = append(result, v)
	// 	}
	// }

	// fmt.Println("\n10. Числа больше 10:", result)
	//11
	s := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	frequency := make(map[int]int)

	for _, v := range s {
		frequency[v]++
	}

	maxCount := 0
	mostFrequent := 0

	for num, count := range frequency {
		if count > maxCount {
			maxCount = count
			mostFrequent = num
		}
	}

	fmt.Println("Самое часто встречающееся число:", mostFrequent)
}
