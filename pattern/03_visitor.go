package pattern

import "strings"

/*
	Реализовать паттерн «посетитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Visitor_pattern
*/

/*
Паттерн Visitor относится к поведенческим паттернам уровня объекта.
+ Паттерн Visitor позволяет обойти набор элементов (объектов) с разнородными интерфейсами,
+ а также позволяет добавить новый метод в класс объекта, при этом, не изменяя сам класс этого объекта.

Требуется для реализации:
    Абстрактный класс Visitor, описывающий интерфейс визитера;
    Структура Person, реализующая конкретного визитера. Реализует методы для обхода конкретного элемента;
    Структура City, реализующая структуру, в которой хранятся элементы для обхода;
    Абстрактный класс Place, реализующий интерфейс элементов структуры City;
    Структура Zoo, реализующий элемент структуры City;
	Структура Theater, реализующий элемент структуры City;
    Структура Circus, реализующий элемент структуры City.
*/

type Visitor interface {
	VisitZoo(z *Zoo) string
	VisitTheater(t *Theater) string
	VisitCircus(c *Circus) string
}

type Place interface {
	Accept(v Visitor) string
}

type Person struct {
}

func (p *Person) VisitZoo(z *Zoo) string {
	return z.SeeAnimal()
}

func (p *Person) VisitCircus(c *Circus) string {
	return c.SeeShow()
}

func (p *Person) VisitTheater(t *Theater) string {
	return t.SeeMovie()
}

type Zoo struct {
}

func (z *Zoo) Accept(v Visitor) string {
	return v.VisitZoo(z)
}

func (z *Zoo) SeeAnimal() string {
	return "See animal..."
}

type Theater struct {
}

func (t *Theater) Accept(v Visitor) string {
	return v.VisitTheater(t)
}

func (t *Theater) SeeMovie() string {
	return "See play..."
}

type Circus struct {
}

func (c *Circus) Accept(v Visitor) string {
	return v.VisitCircus(c)
}

func (c *Circus) SeeShow() string {
	return "See show..."
}

type City struct {
	places []Place
}

func (c *City) Add(p Place) {
	c.places = append(c.places, p)
}

func (c *City) Accept(v Visitor) string {
	var res = strings.Builder{}
	for _, place := range c.places {
		res.WriteString(place.Accept(v))
	}

	return res.String()
}

// Пример работы
/* func main() {
	city := new(City)

	city.Add(&Zoo{})
	city.Add(&Theater{})
	city.Add(&Circus{})

	fmt.Println(city.Accept(&Person{}))
} */
