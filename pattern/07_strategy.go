package pattern

import "sort"

/*
	Реализовать паттерн «стратегия».
Объяснить применимость паттерна, его плюсы и минусы, а также реальные примеры использования данного примера на практике.
	https://en.wikipedia.org/wiki/Strategy_pattern
*/

/*
Поведенческий паттерн проектирования уровня объекта
Определяет схожие алгоритмы и помещает их в свою отдельную структуру
Применим когда нужно использовать разные варианты алгоритма внутри одного объекта
Когда есть множество схожих объектов отличающихся поведением
+ Замена алгоритмов "на лету"
+ Изолирует код от бизнеслогики
+ Уход от наследования
+ Реалезует принцип открытости/закрытости
- Усложнение кода за счет доп объектов
- Клиент должен знать разницу между стратегиями что бы выбрать нужную
*/

// Стратегия с различными алгоритмами сортировки
type Strategy interface {
	Sort([]int) []int
	ShowName() string
}

type Sorter struct {
	Strategy
}

func (s *Sorter) SetAlgorithm(strat Strategy) {
	s.Strategy = strat
}

type BubbleSort struct {
	Name string
}

func (s *BubbleSort) ShowName() string {
	return s.Name
}

func (s *BubbleSort) Sort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	for i := 0; i < len(slice)-1; i++ {
		for j := i + 1; j < len(slice); j++ {
			if slice[i] > slice[j] {
				slice[i], slice[j] = slice[j], slice[i]
			}
		}
	}
	return slice
}

type SeletionSort struct {
	Name string
}

func (s *SeletionSort) ShowName() string {
	return s.Name
}

func (s *SeletionSort) Sort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	for i := 0; i < len(slice)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(slice); j++ {
			if slice[j] < slice[minIdx] {
				minIdx = j

			}
		}
		slice[i], slice[minIdx] = slice[minIdx], slice[i]
	}
	return slice
}

type InsertionSort struct {
	Name string
}

func (s *InsertionSort) ShowName() string {
	return s.Name
}

func (s *InsertionSort) Sort(slice []int) []int {
	for i := 0; i < len(slice); i++ {
		key := slice[i]
		j := i - 1
		for j >= 0 && slice[j] > key {
			slice[j+1] = slice[j]
			j--
		}
		slice[j+1] = key
	}
	return slice
}

type QuickSort struct {
	Name string
}

func (s *QuickSort) ShowName() string {
	return s.Name
}

func (s *QuickSort) Sort(slice []int) []int {
	if len(slice) < 2 {
		return slice
	}

	pivot := slice[0]
	var less = make([]int, 0, len(slice)/2)
	var greater = make([]int, 0, len(slice)/2)
	for _, num := range slice[1:] {
		if num <= pivot {
			less = append(less, num)
		} else {
			greater = append(greater, num)
		}
	}
	res := append(s.Sort(less), pivot)
	res = append(res, s.Sort(greater)...)
	return res
}

type PackageSort struct {
	Name string
}

func (s *PackageSort) ShowName() string {
	return s.Name
}

func (s *PackageSort) Sort(slice []int) []int {
	sort.Ints(slice)
	return slice
}

// Пример рабты
/* func main() {
	size := 50000
	slice := make([]int, size)
	mp := make(map[int]struct{}, size)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
	loop:
		for {
			r := rand.Intn(size)
			if _, ok := mp[r]; !ok {
				slice[i] = r
				mp[r] = struct{}{}
				break loop
			}
		}
	}

	sortTypes := []Strategy{
		&BubbleSort{Name: "Bubble    "},
		&SeletionSort{Name: "Selection "},
		&InsertionSort{Name: "Insertion "},
		&QuickSort{Name: "Quick     "},
		&PackageSort{Name: "Package   "},
	}

	s := Sorter{}
	for _, sort := range sortTypes {
		cop := make([]int, len(slice))
		copy(cop, slice)

		s.SetAlgorithm(sort)
		start := time.Now()
		_ = s.Sort(cop)
		dur := time.Since(start)
		fmt.Printf("%s duration: %v\n", sort.ShowName(), dur)
	}
} */
