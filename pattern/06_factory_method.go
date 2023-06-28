package pattern

import (
	"fmt"
	"log"
)

/*
	Реализовать паттерн «фабричный метод».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Factory_method_pattern
*/

/*
Пораждающий паттерн проектирования который определяет общий интерфейс поведения для создаваемых объектов
+ Паттерн избаваляет от привязки к конкретному типу
+ Общий конструктор создания
+ Упрощает добавление новых объектов
+ Реализует принцип открытости/закрытости
- Может привести к созданию большого количества параллельных структур
- Привязываемся с суперобъекту - конструктору
*/

// Фабрика по производству напитков
type Drink string

const (
	Coala  Drink = "coala"
	Spirte Drink = "spirte"
	Fantan Drink = "fantan"
)

type Creator interface {
	CreateProduct(d Drink) Product
}

type Product interface {
	ShowProduct()
}

type Fabrica struct{}

func NewCreator() Creator {
	return &Fabrica{}
}

func (f *Fabrica) CreateProduct(d Drink) Product {
	switch d {
	case Coala:
		return NewCoalaProduct()
	case Spirte:
		return NewSpirteProduct()
	case Fantan:
		return NewFantanProduct()
	default:
		log.Println("Unknown drink")
		return nil
	}
}

type CoalaProduct struct {
	name  string
	color string
}

func NewCoalaProduct() Product {
	return &CoalaProduct{name: string(Coala), color: "Dark"}
}

func (c *CoalaProduct) ShowProduct() {
	fmt.Printf("%s %s\n", c.color, c.name)
}

type SpirteProduct struct {
	name  string
	color string
}

func NewSpirteProduct() Product {
	return &SpirteProduct{name: string(Spirte), color: "Сolorless"}
}

func (s *SpirteProduct) ShowProduct() {
	fmt.Printf("%s %s\n", s.color, s.name)
}

type FantanProduct struct {
	name  string
	color string
}

func NewFantanProduct() Product {
	return &SpirteProduct{name: string(Fantan), color: "Orange"}
}

func (f *FantanProduct) ShowProduct() {
	fmt.Printf("%s %s\n", f.color, f.name)
}

// Пример работы
/* func main() {
	drinks := []Drink{Coala, Spirte, Fantan}
	fabrica := NewCreator()

	for _, drink := range drinks {
		product := fabrica.CreateProduct(drink)

		if product != nil {
			product.ShowProduct()
		}
	}
}*/
