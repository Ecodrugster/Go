package main

import "fmt"

func main() {
	var isRegistered bool = false
	var isLoggedIn bool = false

	var login string
	var password string

	for {
		fmt.Println("\n=== МЕНЮ ===")
		fmt.Println("1 – Регистрация")
		fmt.Println("2 – Вход")
		fmt.Println("3 – Проверка статуса")
		fmt.Println("4 – Выход")
		fmt.Print("Выберите пункт: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {

		case 1:
			if isRegistered {
				fmt.Println("Вы уже зарегистрированы")
			} else {
				fmt.Print("Введите логин: ")
				fmt.Scan(&login)

				fmt.Print("Введите пароль: ")
				fmt.Scan(&password)

				isRegistered = true
				fmt.Println("Регистрация успешна")
			}

		case 2:
			if !isRegistered {
				fmt.Println("Сначала нужно зарегистрироваться")
			} else if isLoggedIn {
				fmt.Println("Вы уже вошли в систему")
			} else {
				var inputLogin, inputPassword string

				fmt.Print("Логин: ")
				fmt.Scan(&inputLogin)

				fmt.Print("Пароль: ")
				fmt.Scan(&inputPassword)

				if inputLogin == login && inputPassword == password {
					isLoggedIn = true
					fmt.Println("Вход выполнен успешно")
				} else {
					fmt.Println("Неверный логин или пароль")
				}
			}

		case 3:
			if !isRegistered {
				fmt.Println("Статус: не зарегистрирован")
			} else if isLoggedIn {
				fmt.Println("Статус: авторизован")
			} else {
				fmt.Println("Статус: зарегистрирован, но не вошёл")
			}

		case 4:
			fmt.Println("Выход из программы")
			return

		default:
			fmt.Println("Неверный пункт меню")
		}
	}
}
