package pattern

import "strings"

/*
	Реализовать паттерн «фасад».
Объяснить применимость паттерна, его плюсы и минусы,а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Facade_pattern
*/

/*
Это структурный паттерн проектирования, уровня объекта
Представляет собой высокоуровневый интерфейс(т.е. доступ) к сложной системе типов

Разбиение системы на подсистемы позволяет упростить процесс разработки,
а также помогает максимально снизить зависимости одной подсистемы от другой.
Однако использовать такие подсистемы становиться довольно сложно.
Один из способов решения этой проблемы является паттерн Facade.
Задача, сделать единый объект, через который можно было бы взаимодействовать с подсистемами.

+ изолирует клиентов от поведения сложной подсистемы давая доступ к высокоуровневому интерфейсу
- сам интерфейс фасада может стать суперобъектом
*/

// Фасад - структура GoodLife включающая в себя Eat, Sleep, Work
type Food struct {
}

func (f *Food) Eat() string {
	return "Eat"
}

type Bed struct {
}

func (b *Bed) Sleep() string {
	return "Sleep"
}

type Job struct {
}

func (j *Job) Work() string {
	return "Work"
}

type GoodLife struct {
	food *Food
	bed  *Bed
	job  *Job
}

func NewGoodLife() *GoodLife {
	return &GoodLife{
		food: &Food{},
		bed:  &Bed{},
		job:  &Job{},
	}
}

func (life *GoodLife) Go() string {
	spendTime := []string{life.food.Eat(), life.bed.Sleep(), life.job.Work()}
	return strings.Join(spendTime, " ")
}

// Пример работы
/* func main() {
	oneLife := NewGoodLife().Go()
	fmt.Println(oneLife)
} */
