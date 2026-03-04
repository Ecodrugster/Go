package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func IncreasePrice(product *Product) {
	product.Price = product.Price * 1.10
}

func main() {
	products := []Product{
		{Name: "Phone", Price: 50000, Stock: 30},
		{Name: "TV", Price: 120000, Stock: 50},
		{Name: "Headphones", Price: 10000, Stock: 40},
	}
	//  1
	fmt.Println("Все продукты:")
	for _, p := range products {
		fmt.Printf("Название: %s | Цена: %.2f | Количество: %d\n", p.Name, p.Price, p.Stock)
	}
	// 2
	expensive := products[0]

	for _, p := range products {
		if p.Price > expensive.Price {
			expensive = p
		}
	}
	fmt.Printf("\nСамый дорогой продукт: %s | Цена: %.2f\n", expensive.Name, expensive.Price)
	// 3
	fmt.Println("\nУвеличение цены первого продукта на 10%:")
	IncreasePrice(&products[0])

	for _, p := range products {
		fmt.Printf("Название: %s | Цена: %.2f | Количество: %d\n", p.Name, p.Price, p.Stock)
	}
}
