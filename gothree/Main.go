package main

import "fmt"

func changeFirst(s []int) {
	s[0] = 999
}
func addElement(s []int) {
	s = append(s, 100)
	fmt.Println("Внутри функции:", s)
}

func main() {

	// 1
	// var arr [5]int
	// arr[0] = 10
	// arr[1] = 20
	// arr[2] = 30
	// arr[3] = 40
	// arr[4] = 50

	// fmt.Println("Массив:", arr)
	// fmt.Println("Длина:", len(arr))
	//2
	// arr := [4]int{1, 2, 3, 4}
	// fmt.Println("До изменения:", arr)

	// arr[1] = 20
	// arr[3] = 40

	// fmt.Println("После изменения:", arr)
	//3
	// a := [3]int{1,2,3}
	// b := a
	// a[0] = 100
	// fmt.Println("a:", a)
	// fmt.Println("b:", b)
	//4
	// var s []int

	// fmt.Println("Значение:", s)
	// fmt.Println("len:" , len(s))
	// fmt.Println("cap:" , cap(s))
	// fmt.Println("s == nil:", s == nil)
	//5
	// s := []int{1,2,3}
	// fmt.Println("До append len:", len(s), "cap:", cap(s))
	// s = append(s,4)
	// fmt.Println("После append len:", len(s), "cap:", cap(s))
	//6
	// s := make ([]int , 3, 5)

	// s[0] = 1
	// s[1] = 2
	// s[2] = 3

	// fmt.Println("До append len:", len(s), "cap:", cap(s))
	// s = append(s,4,5)
	// fmt.Println("После append len:", len(s), "cap:", cap(s))
	//7
	// arr := [6]int {10,20,30,40,50,60}

	// s := arr[2:5]
	// s[0] = 999

	// fmt.Println("Массив:", arr)
	// fmt.Println("Slice:", s)
	//8
	// arr := [6]int {10,20,30,40,50,60}
	// s := arr[2:5]

	// newSlice := make([]int, len(s))
	// copy(newSlice,s)

	// newSlice[0] = 777
	// fmt.Println("Массив:", arr)
	// fmt.Println("Old Slice:", s)
	// fmt.Println("New Slice:", newSlice)
	// 	9
	// s := []int{1, 2, 3}

	// fmt.Println("До:", s)

	// changeFirst(s)

	// fmt.Println("После:", s)
	//10
	s := make([]int, 3, 3)
	s[0], s[1], s[2] = 1, 2, 3

	fmt.Println("До функции:", s)

	addElement(s)

	fmt.Println("После функции:", s)
}
