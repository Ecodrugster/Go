package main

import "fmt"

type Payment interface {
	Pay(amount float64) bool
}

type Transaction struct {
	Method string
	Amount float64
	Fee    float64
	Total  float64
	Ok     bool
}

type CardPayment struct {
	Balance    float64
	FeePercent float64
	Limit      float64
	IsBlocked  bool
	FailCount  int
	History    []Transaction
}

func (c *CardPayment) Pay(amount float64) bool {

	if c.IsBlocked {
		fmt.Println("Карта заблокирована")
		c.addHistory("card", amount, 0, 0, false)
		return false
	}

	if amount <= 0 {
		c.addHistory("card", amount, 0, 0, false)
		return false
	}

	if amount > c.Limit {
		fmt.Println("Превышен лимит")
		c.FailCount++
		c.addHistory("card", amount, 0, 0, false)
		return false
	}

	fee := amount * c.FeePercent / 100
	total := amount + fee

	if total > c.Balance {
		fmt.Println("Недостаточно средств")
		c.FailCount++
		c.addHistory("card", amount, fee, total, false)
		return false
	}

	c.Balance -= total
	fmt.Printf("Оплата картой: %.2f\n", amount)

	c.FailCount = 0
	c.addHistory("card", amount, fee, total, true)

	return true
}

func (c *CardPayment) addHistory(method string, amount, fee, total float64, ok bool) {
	c.History = append(c.History, Transaction{
		Method: method,
		Amount: amount,
		Fee:    fee,
		Total:  total,
		Ok:     ok,
	})
}

type CashPayment struct {
	Balance    float64
	FeePercent float64
	Limit      float64
	History    []Transaction
}

func (c *CashPayment) Pay(amount float64) bool {

	if amount <= 0 {
		c.addHistory("cash", amount, 0, 0, false)
		return false
	}

	if amount > c.Limit {
		fmt.Println("Превышен лимит")
		c.addHistory("cash", amount, 0, 0, false)
		return false
	}

	fee := amount * c.FeePercent / 100
	total := amount + fee

	if total > c.Balance {
		fmt.Println("Недостаточно наличных")
		c.addHistory("cash", amount, fee, total, false)
		return false
	}

	c.Balance -= total
	fmt.Printf("Оплата наличными: %.2f\n", amount)

	c.addHistory("cash", amount, fee, total, true)

	return true
}

func (c *CashPayment) addHistory(method string, amount, fee, total float64, ok bool) {
	c.History = append(c.History, Transaction{
		Method: method,
		Amount: amount,
		Fee:    fee,
		Total:  total,
		Ok:     ok,
	})
}

type CryptoPayment struct {
	Balance    float64
	FeePercent float64
	Limit      float64
	History    []Transaction
}

func (c *CryptoPayment) Pay(amount float64) bool {

	if amount <= 0 {
		c.addHistory("crypto", amount, 0, 0, false)
		return false
	}

	if amount > c.Limit {
		fmt.Println("Превышен лимит")
		c.addHistory("crypto", amount, 0, 0, false)
		return false
	}

	fee := amount * c.FeePercent / 100
	total := amount + fee

	if total > c.Balance {
		fmt.Println("Недостаточно криптовалюты")
		c.addHistory("crypto", amount, fee, total, false)
		return false
	}

	c.Balance -= total
	fmt.Printf("Оплата криптой: %.2f\n", amount)

	c.addHistory("crypto", amount, fee, total, true)

	return true
}

func (c *CryptoPayment) addHistory(method string, amount, fee, total float64, ok bool) {
	c.History = append(c.History, Transaction{
		Method: method,
		Amount: amount,
		Fee:    fee,
		Total:  total,
		Ok:     ok,
	})
}

func ProcessPayment(p Payment, amount float64) {
	fmt.Println("\nНачинаем оплату...")

	success := p.Pay(amount)

	if success {
		fmt.Println("Оплата успешна")
	} else {
		fmt.Println("Оплата не прошла")
	}

	if card, ok := p.(*CardPayment); ok {
		if card.FailCount >= 3 {
			card.IsBlocked = true
			fmt.Println("Карта автоматически заблокирована (3 ошибки)")
		}
	}
}

func main() {

	card := &CardPayment{
		Balance:    1000,
		FeePercent: 2,
		Limit:      500,
	}

	ProcessPayment(card, -10)

	ProcessPayment(card, 2000)

	ProcessPayment(card, 600)

	ProcessPayment(card, 100)
	ProcessPayment(card, 100)

	ProcessPayment(card, 2000)
	ProcessPayment(card, 2000)
	ProcessPayment(card, 2000)

	ProcessPayment(card, 50)

	fmt.Println("\nИстория операций карты:")
	for _, t := range card.History {
		fmt.Printf("%+v\n", t)
	}

	// type Product struct {
	// 	Name  string
	// 	Price float64
	// 	Stock int
	// }

	// func IncreasePrice(product *Product) {
	// 	product.Price = product.Price * 1.10
	// }

	// func main() {
	// 	products := []Product{
	// 		{Name: "Phone", Price: 50000, Stock: 30},
	// 		{Name: "TV", Price: 120000, Stock: 50},
	// 		{Name: "Headphones", Price: 10000, Stock: 40},
	// 	}
	// 	//  1
	// 	fmt.Println("Все продукты:")
	// 	for _, p := range products {
	// 		fmt.Printf("Название: %s | Цена: %.2f | Количество: %d\n", p.Name, p.Price, p.Stock)
	// 	}
	// 	// 2
	// 	expensive := products[0]

	// 	for _, p := range products {
	// 		if p.Price > expensive.Price {
	// 			expensive = p
	// 		}
	// 	}
	// 	fmt.Printf("\nСамый дорогой продукт: %s | Цена: %.2f\n", expensive.Name, expensive.Price)
	// 	// 3
	// 	fmt.Println("\nУвеличение цены первого продукта на 10%:")
	// 	IncreasePrice(&products[0])

	// 	for _, p := range products {
	// 		fmt.Printf("Название: %s | Цена: %.2f | Количество: %d\n", p.Name, p.Price, p.Stock)
	// 	}
}
