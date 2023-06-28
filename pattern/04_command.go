package pattern

import "strings"

/*
	Реализовать паттерн «комманда».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Command_pattern
*/

/*
Паттерн Command относится к поведенческим паттернам уровня объекта.
Паттерн Command позволяет представить запрос в виде объекта.
Из этого следует, что команда - это объект.
Такие запросы, например, можно ставить в очередь, отменять или возобновлять.

Паттерн Command отделяет объект, инициирующий операцию, от объекта, который знает, как ее выполнить.
Единственное, что должен знать инициатор, это как отправить команду.

Требуется для реализации:

    Базовый абстрактный класс Command описывающий интерфейс команды;
   	Структура TurnOnCommand, реализующая команду;
	Структура TurnOffCommand, реализующая команду;
	Класс Receiver, реализующий получателя и имеющий набор действий, которые команда можем запрашивать;
    Класс Sender, реализующий инициатора, записывающий команду и провоцирующий её выполнение;

Sender умеет складывать команды в стопку и инициировать их выполнение по какому-то событию.
Обратившись к Sender можно отменить команду, пока та не выполнена.

ConcreteCommand содержит в себе запросы к Receiver, которые тот должен выполнять.
В свою очередь Receiver содержит только набор действий (Actions),
которые выполняются при обращении к ним из ConcreteCommand.
*/

type Command interface {
	Execute() string
}

type TurnOnCommand struct {
	Receiver *Receiver
}

func (c *TurnOnCommand) Execute() string {
	return c.Receiver.TurnOn()
}

type TurnOffCommand struct {
	Receiver *Receiver
}

func (c *TurnOffCommand) Execute() string {
	return c.Receiver.TurnOff()
}

type Receiver struct {
}

func (r *Receiver) TurnOn() string {
	return "Toggle On"
}

func (r *Receiver) TurnOff() string {
	return "Toggle Off"
}

type Sender struct {
	commands []Command
}

func (s *Sender) SaveCommand(command Command) {
	s.commands = append(s.commands, command)
}

func (s *Sender) Execute() string {
	result := strings.Builder{}
	for _, command := range s.commands {
		result.WriteString(command.Execute() + "\n")
	}
	return result.String()
}

// Пример работы
/* func main() {
	sender := new(Sender)
	receiver := new(Receiver)

	sender.SaveCommand(&TurnOnCommand{Receiver: receiver})
	sender.SaveCommand(&TurnOffCommand{Receiver: receiver})

	fmt.Print(sender.Execute())
} */
