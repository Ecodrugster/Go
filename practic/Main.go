package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Grade float64
}

type Course struct {
	Title    string
	Students []Student
}

func (c *Course) AddStudent(s Student) {
	c.Students = append(c.Students, s)
}

func (c Course) PrintStudents() {
	fmt.Println("Курс:", c.Title)
	for _, s := range c.Students {
		fmt.Printf("Имя: %s | Возраст: %d | Балл: %.2f\n",
			s.Name, s.Age, s.Grade)
	}
}

func (c Course) Count() int {
	return len(c.Students)
}

func (c Course) TopStudent() *Student {
	if len(c.Students) == 0 {
		return nil
	}
	top := c.Students[0]

	for i := range c.Students {
		if c.Students[i].Grade > top.Grade {
			top = c.Students[i]
		}
	}
	return &top
}
func (s Student) Info() string {
	return fmt.Sprintf("Name: %s, Age: %d, Grade: %.2f", s.Name, s.Age, s.Grade)
}

func (s *Student) Birthday() {
	s.Age++
}

func FindByName(students []Student, name string) *Student {
	for i := range students {
		if students[i].Name == name {
			return &students[i]
		}
	}
	return nil
}

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Shop struct {
	Name     string
	Products []Product
}

func (s *Shop) AddProduct(p Product) {
	s.Products = append(s.Products, p)
}

func (s Shop) PrintProducts() {
	fmt.Println("Магазин: ", s.Name)
	for _, p := range s.Products {
		fmt.Printf("Id: %d | %s | Price: %.2f\n", p.ID, p.Name, p.Price)
	}
}

func (s Shop) FindById(id int) *Product {
	for i := range s.Products {
		if s.Products[i].ID == id {
			return &s.Products[i]
		}
	}
	return nil
}

func (s Shop) TotalPrice() float64 {
	var sum float64
	for _, p := range s.Products {
		sum += p.Price
	}
	return sum
}

type User struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {

	// 1
	// student1 := Student{Name: "Alice", Age: 20, Grade: 3.5}
	// student2 := Student{Name: "Bob", Age: 22, Grade: 3.8}

	// fmt.Println("До изменения: ")
	// fmt.Println(student1.Info())
	// fmt.Println(student2.Info())

	// student1.Birthday()

	// fmt.Println("\nПосле изменения: ")
	// fmt.Println(student1.Info())
	// fmt.Println(student2.Info())

	// 2
	// students := []Student{
	// 	{Name: "Alice", Age: 20, Grade: 3.5},
	// 	{Name: "Bob", Age: 22, Grade: 3.8},
	// 	{Name: "Charlie", Age: 21, Grade: 3.2},
	// 	{Name: "Diana", Age: 23, Grade: 3.9},
	// 	{Name: "Erasil", Age: 21, Grade: 3.7},
	// }

	// if len(students) == 0 {
	// 	fmt.Println("Список пуст")
	// 	return
	// }
	// oldest := students[0]
	// for _, s := range students {
	// 	if s.Age > oldest.Age {
	// 		oldest = s
	// 	}
	// }

	// fmt.Println("Самый старший студент:")
	// fmt.Printf("%s (%d лет)\n", oldest.Name, oldest.Age)

	// var sum float64
	// for _, s := range students {
	// 	sum += s.Grade
	// }
	// average := sum / float64(len(students))

	// fmt.Printf("Средний балл группы: %.2f\n", average)

	// found := FindByName(students, "Diana")

	// if found != nil {
	// 	fmt.Println("Студент найден:", found.Name, found.Age, found.Grade)
	// } else {
	// 	fmt.Println("Студент не найден")
	// }

	// 3
	// course := Course{Title: "Go Programming"}

	// course.AddStudent(Student{"Alice", 20, 3.5})
	// course.AddStudent(Student{"Bob", 22, 3.8})
	// course.AddStudent(Student{"Charlie", 21, 4.2})

	// course.PrintStudents()

	// fmt.Println("Количество студентов:", course.Count())

	// top := course.TopStudent()
	// if top != nil {
	// 	fmt.Printf("Лучший студент: %s с баллом %.2f\n", top.Name, top.Grade)
	// }

	// 4
	// 	shop := Shop{Name: "ElectroMart"}

	// 	shop.AddProduct(Product{ID: 1, Name: "TV", Price: 60000})
	// 	shop.AddProduct(Product{ID: 2, Name: "Laptop", Price: 80000})
	// 	shop.AddProduct(Product{ID: 3, Name: "Smartphone", Price: 50000})

	// 	shop.PrintProducts()

	// 	fmt.Println("Общая стоймость товаров: ", shop.TotalPrice())

	// 	found := shop.FindById(2)
	// 	if found != nil {
	// 		fmt.Printf("Найден товар: %s с ценой %.2f\n", found.Name, found.Price)
	// 	} else {
	// 		fmt.Println("Товар не найден")
	// 	}

	// 5 1)
	// user := User{
	// 	ID:    1,
	// 	Name:  "Alice",
	// 	Email: "alice@gmail.com",
	// 	Age:   20,
	// }

	// jsonData, err := json.Marshal(user)
	// if err != nil {
	// 	fmt.Println("Ошибка : ", err)
	// 	return
	// }

	// fmt.Println("Struct to JSON: ")
	// fmt.Println(string(jsonData))

	// 5 2)
	// jsonString := `{"id":1,"name":"Alice","email":"alice@gmail.com","age":20}`

	// var user User
	// err := json.Unmarshal([]byte(jsonString), &user)
	// if err != nil {
	// 	fmt.Println("Ошибка: ", err)
	// 	return
	// }
	// fmt.Println("\nJSON to Struct: ")
	// fmt.Printf("%+v\n", user)

}
