package main

import "fmt"

// type Book struct {
// 	Title       string
// 	Author      string
// 	pages       int
// 	isAvailable bool
// }

// func newBook(title, author string, pages int) Book {
// 	if pages <= 0 {
// 		pages = 1
// 	}

// 	return Book{
// 		Title:       title,
// 		Author:      author,
// 		pages:       pages,
// 		isAvailable: true,
// 	}
// }

// func (b Book) Info() {
// 	fmt.Println("Книга:", b.Title)
// 	fmt.Println("Автор:", b.Author)
// 	fmt.Println("Страниц:", b.pages)
// 	fmt.Println("Доступна:", b.isAvailable)
// 	fmt.Println()
// }

// func (b *Book) Borrow() {
// 	if !b.isAvailable {
// 		fmt.Println("Книга уже выдана")
// 		return
// 	}
// 	b.isAvailable = false
// 	fmt.Println("Книга выдана")
// }

// func (b *Book) ReturnBook() {
// 	b.isAvailable = true
// 	fmt.Println("Книга возвращена")
// }

// func (b Book) GetPages() int {
// 	return b.pages
// }

// func (b *Book) SetPages(p int) {
// 	if p <= 0 {
// 		fmt.Println("Ошибка: страницы должны быть > 0")
// 		return
// 	}
// 	b.pages = p
// }

// // 2
// type Worker interface {
// 	Work() string
// 	GetName() string
// }

// type Programmer struct {
// 	Name     string
// 	Language string
// }

// func (p Programmer) Work() string {
// 	return "Программист " + p.Name + " пишет код на " + p.Language
// }

// func (p Programmer) GetName() string {
// 	return p.Name
// }

// type Designer struct {
// 	Name string
// 	Tool string
// }

// func (d Designer) Work() string {
// 	return "Дизайнер " + d.Name + " делает макет в " + d.Tool
// }

// func (d Designer) GetName() string {
// 	return d.Name
// }

// func ShowWork(w Worker) {
// 	fmt.Println("Имя:", w.GetName())
// 	fmt.Println(w.Work())
// 	fmt.Println()
// }

// 3

type Product struct {
	Name     string
	price    float64
	Quantity int
}

func newProduct(name string, price float64, quantity int) Product {
	if price < 0 {
		price = 0
	}
	if quantity < 0 {
		quantity = 0
	}

	return Product{
		Name:     name,
		price:    price,
		Quantity: quantity,
	}
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p *Product) SetPrice(newPrice float64) {
	if newPrice < 0 {
		fmt.Println("Ошибка: цена не может быть отрицательной")
		return
	}
	p.price = newPrice
}

func (p *Product) Buy(amount int) {
	if amount <= 0 {
		fmt.Println("Ошибка: некорректное количество")
		return
	}
	if amount > p.Quantity {
		fmt.Println("Недостаточно товара")
		return
	}

	p.Quantity -= amount
	fmt.Println("Покупка успешна")
}

func (p *Product) Restock(amount int) {
	if amount > 0 {
		p.Quantity += amount
	}
}

func (p Product) Info() {
	fmt.Println("Товар:", p.Name)
	fmt.Println("Цена:", p.price)
	fmt.Println("Остаток:", p.Quantity)
	fmt.Println()
}

func main() {

	// 1
	// book := newBook("Go Basics", "John", 200)
	// book.Info()

	// book.Borrow()
	// book.Borrow()
	// book.ReturnBook()

	// book.SetPages(300)
	// book.Info()

	// // 2
	// p1 := Programmer{"Ali", "Go"}
	// p2 := Programmer{"Dana", "Java"}

	// d1 := Designer{"Amina", "Figma"}
	// d2 := Designer{"Olga", "Photoshop"}

	// ShowWork(p1)
	// ShowWork(p2)
	// ShowWork(d1)
	// ShowWork(d2)

	// 3
	product := newProduct("Phone", 1000, 5)
	product.Info()

	product.Buy(2)
	product.Buy(10)

	product.SetPrice(1200)
	product.Restock(3)

	product.Info()
}
