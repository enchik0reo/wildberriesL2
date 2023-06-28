package pattern

import "fmt"

/*
	Реализовать паттерн «строитель».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Builder_pattern
*/

/*
Пораджающий паттерн проектирования уровня объекта
Позволяет создавать сложные объекты поэтапно, на каждом этапе получается часть общего большого объекта

+ позволяет создавать пошагово и гибко большой объект из маленьких независимых частей
+ использует один и тот же код для создания различных объектов
+ изолирует сложный код сборки объекта от его бизнеслогики
- усложняет код программы из за введения доп структур и интерфейсов
- клиент будет привязан к конкретному интерфейсу строителя, при добавлении нового строителя его нужно реализовывать
*/

// Фабрика по производству ноутбуков

/*
Структура Factory, которая будет распоряжаться строителем и отдавать ему команды в определённом порядке;
Интерфейс Builder, который описывает команды, которые конкретный билдер обязан выполнять;
Конкретные билдеры AsusBuilder и MsiBuilder которые реализуют Builder и взаимодействуют со сложным объектом;
Сложный объект - структура Laptop, её будет "собирать" билдер.
*/

type Laptop struct {
	Brand       string
	Core        int
	Memory      int
	GraphicCard int
	Diagonal    float64
}

func (l *Laptop) Show() {
	fmt.Printf("%s | Core: [%d] Memory: [%d]Gb GraphicCard: [%d]Gb Diagonal: [%0.1f]inches\n", l.Brand, l.Core, l.Memory, l.GraphicCard, l.Diagonal)
}

type Builder interface {
	SetBrand()
	SetCore()
	SetMemory()
	SetGraphicCard()
	SetDiagonal()
	MakeLaptop() Laptop
}

type AsusBuilder struct {
	Laptop
}

func (a *AsusBuilder) SetBrand() {
	a.Brand = "Asus"
}

func (a *AsusBuilder) SetCore() {
	a.Core = 4
}

func (a *AsusBuilder) SetMemory() {
	a.Memory = 8
}

func (a *AsusBuilder) SetGraphicCard() {
	a.GraphicCard = 2
}

func (a *AsusBuilder) SetDiagonal() {
	a.Diagonal = 14
}

func (a *AsusBuilder) MakeLaptop() Laptop {
	return Laptop{
		Core:        a.Core,
		Brand:       a.Brand,
		Memory:      a.Memory,
		GraphicCard: a.GraphicCard,
		Diagonal:    a.Diagonal,
	}
}

type MsiBuilder struct {
	Laptop
}

func (m *MsiBuilder) SetBrand() {
	m.Brand = "MSI"
}

func (m *MsiBuilder) SetCore() {
	m.Core = 8
}

func (m *MsiBuilder) SetMemory() {
	m.Memory = 16
}

func (m *MsiBuilder) SetGraphicCard() {
	m.GraphicCard = 8
}

func (m *MsiBuilder) SetDiagonal() {
	m.Diagonal = 15.6
}

func (m *MsiBuilder) MakeLaptop() Laptop {
	return Laptop{
		Core:        m.Core,
		Brand:       m.Brand,
		Memory:      m.Memory,
		GraphicCard: m.GraphicCard,
		Diagonal:    m.Diagonal,
	}
}

// Asus or MSI
func MakeBuilder(builder string) Builder {
	switch builder {
	case "Asus":
		return &AsusBuilder{}
	case "MSI":
		return &MsiBuilder{}
	default:
		panic("invalid builder")
	}
}

type Factory struct {
	Builder Builder
}

func NewFactory(builder Builder) *Factory {
	return &Factory{
		Builder: builder,
	}
}

func (f *Factory) ChangeBuilder(builder Builder) {
	f.Builder = builder
}

func (f *Factory) Construct() Laptop {
	f.Builder.SetBrand()
	f.Builder.SetCore()
	f.Builder.SetMemory()
	f.Builder.SetGraphicCard()
	f.Builder.SetDiagonal()
	return f.Builder.MakeLaptop()
}

// Пример работы
/* func main() {
	asusBuilder := MakeBuilder("Asus")
	msiBuilder := MakeBuilder("MSI")

	factory := NewFactory(asusBuilder)

	laptop := factory.Construct()
	laptop.Show()

	factory.ChangeBuilder(msiBuilder)

	laptop = factory.Construct()
	laptop.Show()
} */
