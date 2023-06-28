package pattern

/*
	Реализовать паттерн «состояние».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/State_pattern
*/

/*
Паттерн State относится к поведенческим паттернам уровня объекта.
Паттерн State позволяет объекту изменять свое поведение в зависимости от внутреннего состояния.
Поведение объекта изменяется настолько, что создается впечатление, будто изменился класс объекта.

Паттерн должен применяться:
- когда поведение объекта зависит от его состояния
- поведение объекта должно изменяться во время выполнения программы
- состояний достаточно много и использовать для этого условные операторы, разбросанные по коду, достаточно затруднительно

+ Избавляет от множество услоных операторов
- Неоравданно усложняет код если состояний мало и они редко меняются
*/

/*
Структура MobilePhone, представляет собой объектно-ориентированное представление состояния;
Абстрактный класс MobileAlertStater, определяющий интерфейс различных состояний;
Структуры Sound, Vibration, Off реализуют одно из поведений, ассоциированное с определенным состоянием.
*/

type MobileAlertStater interface {
	Alert() string
}

type MobilePhone struct {
	state MobileAlertStater
}

func (a *MobilePhone) Alert() string {
	return a.state.Alert()
}

func (a *MobilePhone) SetAlertState(state MobileAlertStater) {
	a.state = state
}

func NewMobile() *MobilePhone {
	return &MobilePhone{state: &MobileAlertSound{}}
}

type MobileAlertSound struct {
}

func (a *MobileAlertSound) Alert() string {
	return "Ring ring... Ring ring...\n"
}

type MobileAlertVibration struct {
}

func (a *MobileAlertVibration) Alert() string {
	return "Vrrr... Brrr...\n"
}

type MobileAlertOff struct {
}

func (a *MobileAlertOff) Alert() string {
	return "...\n"
}

// Пример использования
/* func main() {
	mobile := NewMobile()
	alert := mobile.Alert()

	mobile.SetAlertState(&MobileAlertVibration{})
	alert += mobile.Alert()

	mobile.SetAlertState(&MobileAlertOff{})
	alert += mobile.Alert()

	fmt.Print(alert)
} */
