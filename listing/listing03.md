Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}

Ответ:
<nil>
false

Сам по себе интерфейс это структура с двумя полями, данными об объекте находящимся в интерфейсе и указателем на
сами данные. Интерфейс представляет собой "контракт" (набор методов) который должен иметь объект чтобы
реализовывать его (в го неявная "утинная типизация").
Сравнение с nil будет true только когда тип и значение будут nil.
Пустой интерфейс без типа и значения это <nil> <nil> == nil.
Интерфейс с нил элементом определённого типа это <type> <nil> != nil.
Интерфейс с не нил элементом это <type> <value> != nil.