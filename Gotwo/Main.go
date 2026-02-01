package main

import "fmt"

func main() {
	menu()
}

func inRange(x int, min int, max int) bool {
	return x >= min && x <= max
}

func ageCategory(age int) string {
	if age < 12 {
		return "Child"
	} else if age <= 17 {
		return "Teen"
	} else if age <= 64 {
		return "Adult"
	}
	return "Senior"
}

func compareThree(a int, b int, c int) string {
	if a == b && b == c {
		return "all equal"
	} else if a == b || a == c || b == c {
		return "two equal"
	}
	return "all different"
}

func countDivisors(n int) int {
	count := 0
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			count++
		}
	}
	return count
}

func firstEven() int {
	var x int
	for {
		fmt.Print("Введите число: ")
		fmt.Scan(&x)
		if x%2 == 0 {
			return x
		}
	}
}

func averageUntilZero() float64 {
	sum := 0
	count := 0
	var x int

	for {
		fmt.Print("Введите число (0 — конец): ")
		fmt.Scan(&x)
		if x == 0 {
			break
		}
		sum += x
		count++
	}

	if count == 0 {
		return 0
	}

	return float64(sum) / float64(count)
}

func readPositive() int {
	var x int
	for {
		fmt.Print("Введите положительное число: ")
		fmt.Scan(&x)
		if x > 0 {
			return x
		}
	}
}

func safeCompare(a int, b int) (string, bool) {
	if a == b {
		return "equal", true
	} else if a > b {
		return "a greater", true
	}
	return "b greater", true
}

func menu() {
	var choice int

	for {
		fmt.Println("\n=== МЕНЮ ===")
		fmt.Println("1 - Проверка диапазона")
		fmt.Println("2 - Категория возраста")
		fmt.Println("3 - Количество делителей")
		fmt.Println("4 - Среднее до нуля")
		fmt.Println("0 - Выход")
		fmt.Print("Выберите пункт: ")

		fmt.Scan(&choice)

		switch choice {

		case 1:
			var x, min, max int
			fmt.Print("Введите x: ")
			fmt.Scan(&x)
			fmt.Print("Введите min: ")
			fmt.Scan(&min)
			fmt.Print("Введите max: ")
			fmt.Scan(&max)
			fmt.Println("Результат:", inRange(x, min, max))

		case 2:
			var age int
			fmt.Print("Введите возраст: ")
			fmt.Scan(&age)
			fmt.Println("Категория:", ageCategory(age))

		case 3:
			var n int
			fmt.Print("Введите число: ")
			fmt.Scan(&n)
			fmt.Println("Количество делителей:", countDivisors(n))

		case 4:
			fmt.Println("Среднее:", averageUntilZero())

		case 0:
			fmt.Println("Выход из программы")
			return

		default:
			fmt.Println("Неверный пункт меню")
		}
	}
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isPowerOfTwo(n int) bool {
	if n <= 0 {
		return false
	}
	for n > 1 {
		if n%2 != 0 {
			return false
		}
		n /= 2
	}
	return true
}
