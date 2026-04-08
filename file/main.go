package main

import (
	"bufio"
	"fmt"
	"os"
)

type Student struct {
	Name  string
	Age   int
	Group string
}

// 4
type Book struct {
	Title string
	Author string
	Year int
}

func (b Book) Display() {
	fmt.Println("Title:", b.Title)
	fmt.Println("Author:", b.Author)
	fmt.Println("Year:", b.Year)
}
func (b Book) SaveToFile(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	data := fmt.Sprintf("Title: %s\nAuthor: %s\nYear: %d\n",
		b.Title, b.Author, b.Year)
	_, err = file.WriteString(data)
	return err
}
func (b Book) ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

// func student
// 6
func (s Student) Validate() error {
	if s.Name == "" {
		return fmt.Errorf("Имя не может быть пустым")
	}
	if s.Age < 0 {
		return fmt.Errorf("Возраст не может быть отрицательным")
	}
	if s.Group == "" {
		return fmt.Errorf("Группа не может быть пустой")
	}
	return nil
}

// 1
func (s Student) Display() {
	fmt.Println("Name:", s.Name)
	fmt.Println("Age:", s.Age)
	fmt.Println("Group:", s.Group)
}

// 2
func (s Student) SaveToFile(filename string) error {
	if err := s.Validate(); err != nil {
		return err
	}
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	data := fmt.Sprintf("Имя: %s\nВозраст: %d\nГруппа: %s\n",
		s.Name, s.Age, s.Group)

	_, err = file.WriteString(data)
	return err
}

// 3
func (s Student) ReadFile(filename string) error {
	file, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return scanner.Err()
}

// 5
	type FileManager struct {
		FileName string
	}

	func (fm FileManager) Write(data string) error {
		file, err := os.Create(fm.FileName)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(data)
		return err
	}
	func (fm FileManager) Read() (string, error) {
		file, err := os.Open(fm.FileName)
		if err != nil {
			return "", err
		}
		defer file.Close()
		var data string
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			data += scanner.Text() + "\n"
		}
		return data, scanner.Err()
	}
	// 7
	type Note struct {
		Text string
	}
	func AddNote(filename string) error {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите текст заметки:")
		text, err := reader.ReadString('\n')
		if err != nil {
			return err
		}
		file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			return err
		}
		defer file.Close()
		_, err = file.WriteString(text)
		if err != nil {
			return err
		}
		return nil
	}
	func ShowNote(filename string) (string, error) {
		file, err := os.Open(filename)
		if err != nil {
			return "", err
		}
		defer file.Close()
		var data string
		scanner := bufio.NewScanner(file)
		fmt.Println("Текст заметки:")
		for scanner.Scan() {
			data += scanner.Text() + "\n"
		}
		return data, scanner.Err()
	}



func main() {

	// 1
	student := Student{"Ali", 17, "Go-101"}
	student.Display()

	// // 2
	err := student.SaveToFile("student.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Студент сохранен в файл")
	}
	// 3
	fmt.Println("\nЧтение student.txt:")
	err = student.ReadFile("student.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	// 4
	book := Book{"1984", "George Orwell", 1949}

	err = book.SaveToFile("book.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}

	fmt.Println("\nЧтение book.txt:")
	err = book.ReadFile("book.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	}
	// 5
	fileManager := FileManager{"file.txt"}
	data, err := fileManager.Read()
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Данные из файла:", data)
	}
	err = fileManager.Write("Hello, World!")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Данные записаны в файл")
	}
	// 7
	err = AddNote("notes.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Заметка добавлена")
	}

	data, err = ShowNote("notes.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Текст заметки:", data)
	}
}