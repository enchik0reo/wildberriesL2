package pattern

import "fmt"

/*
	Реализовать паттерн «цепочка вызовов».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Chain-of-responsibility_pattern
*/

/*
Поведенческий патерн проектирования уровня объекта.
Позволяет передавать выполнение запроса последовательно по цепочке
Обработчики, которые по очереди получают запрос, сами решают, обрабатывать его или нет.
+ Уменьшает зависимость между клиентом и обработчиками, каждый обработчик выполняет свою логику
+ Реализует принцип единственной ответственности
+ Реализует принцип открытости для расширения закртости для изменения
- Если запрос не обработан ни одним обработчиком, то он просто теряется
*/

// Взаимодействие нескольких сервисов по цепочке

// Абстрактный класс
type Handler interface {
	SendRequest(count int) string
	SetNext(Handler)
}

// Конкретная структура, реализующая абстрактный класс
type ConcreteHandlerA struct {
	Name string
	next Handler
}

// Логика конкретного обработчика
func (h *ConcreteHandlerA) SendRequest(count int) string {
	if count == 0 {
		return fmt.Sprintf("Count [%d], forced stop, Handler %s", count, h.Name)
	} else if h.next != nil {
		fmt.Printf("Count [%d] processed, Handler %s\n", count, h.Name)
		count -= 1
		return h.next.SendRequest(count)
	} else {
		return fmt.Sprintf("Count save [%d], second Handler %s", count, h.Name)
	}
}

// Задание следующего в цепочке обработчика
func (h *ConcreteHandlerA) SetNext(handler Handler) {
	h.next = handler
}

type ConcreteHandlerB struct {
	Name string
	next Handler
}

func (h *ConcreteHandlerB) SendRequest(count int) string {
	if count == 0 {
		return fmt.Sprintf("Count [%d], forced stop, Handler %s", count, h.Name)
	} else if h.next != nil {
		fmt.Printf("Count [%d] processed, Handler %s\n", count, h.Name)
		count -= 1
		return h.next.SendRequest(count)
	} else {
		return fmt.Sprintf("Count save [%d], second Handler %s", count, h.Name)
	}
}

func (h *ConcreteHandlerB) SetNext(handler Handler) {
	h.next = handler
}

type ConcreteHandlerC struct {
	Name string
	next Handler
}

func (h *ConcreteHandlerC) SendRequest(count int) string {
	if count == 0 {
		return fmt.Sprintf("Count [%d], forced stop, Handler %s", count, h.Name)
	} else if h.next != nil {
		fmt.Printf("Count [%d] processed, Handler %s\n", count, h.Name)
		count -= 1
		return h.next.SendRequest(count)
	} else {
		return fmt.Sprintf("Count save [%d], second Handler %s", count, h.Name)
	}
}

func (h *ConcreteHandlerC) SetNext(handler Handler) {
	h.next = handler
}

// Пример работы
/* func main() {
	handler1 := &ConcreteHandlerA{Name: "A"}
	handler2 := &ConcreteHandlerB{Name: "B"}
	handler3 := &ConcreteHandlerC{Name: "C"}

	handler1.SetNext(handler2)
	handler2.SetNext(handler3)

	fmt.Println(handler1.SendRequest(3))
} */
