package main

/*
=== Поиск анаграмм по словарю ===

Напишите функцию поиска всех множеств анаграмм по словарю.
Например:
'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
'листок', 'слиток' и 'столик' - другому.

Входные данные для функции: ссылка на массив - каждый элемент которого - слово на русском языке в кодировке utf8.
Выходные данные: Ссылка на мапу множеств анаграмм.
Ключ - первое встретившееся в словаре слово из множества
Значение - ссылка на массив, каждый элемент которого, слово из множества. Массив должен быть отсортирован по возрастанию.
Множества из одного элемента не должны попасть в результат.
Все слова должны быть приведены к нижнему регистру.
В результате каждое слово должно встречаться только один раз.

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"sort"
	"strings"
)

func Anagrams(words []string) map[string][]string {
	if len(words) < 2 {
		return nil
	}

	res := make(map[string][]string, len(words))

	for i := 0; i < len(words)-1; i++ {
		words[i] = strings.ToLower(words[i])
		if _, ok := res[words[i]]; !ok {
			res[words[i]] = append(res[words[i]], words[i])
			for j := i + 1; j < len(words); j++ {
				if len(words[i]) == len(words[j]) {
					var chk bool
					for _, ri := range words[i] {
						chk = false
						for _, rj := range words[j] {
							if ri == rj {
								chk = true
								break
							}
						}
						if !chk {
							break
						}
					}
					if chk {
						res[words[i]] = append(res[words[i]], words[j])
						res[words[j]] = []string{}
					}
				}
			}
		}
	}

	for key, anagram := range res {
		if len(anagram) < 2 {
			delete(res, key)
			continue
		}
		sort.Slice(anagram, func(i, j int) bool { return anagram[i] < anagram[j] })
	}

	return res
}

func main() {
	fmt.Println(Anagrams([]string{"слиток", "пятак", "столик", "пятка", "тяпка", "листок", "арбуз"}))
}
