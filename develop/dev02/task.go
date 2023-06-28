package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
=== Задача на распаковку ===

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:
	- "a4bc2d5e" => "aaaabccddddde"
	- "abcd" => "abcd"
	- "45" => "" (некорректная строка)
	- "" => ""
Дополнительное задание: поддержка escape - последовательностей
	- qwe\4\5 => qwe45 (*)
	- qwe\45 => qwe44444 (*)
	- qwe\\5 => qwe\\\\\ (*)

В случае если была передана некорректная строка функция должна возвращать ошибку. Написать unit-тесты.

Функция должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func Unpack(str string) (string, error) {
	if str == "" {
		return "", nil
	}

	if str[0] < 'a' || str[0] > 'z' {
		return str, fmt.Errorf("invalide string")
	}

	res := strings.Builder{}
	res.Grow(len(str))

	for i := 0; i <= len(str)-1; i++ {
		if str[i] >= 'a' && str[i] <= 'z' {
			res.WriteByte(str[i])
		} else if str[i] == '\\' {
			res.WriteByte(str[i+1])
			i++
		} else if str[i] >= '1' && str[i] <= '9' {
			n, err := strconv.Atoi(string(str[i]))
			if err != nil {
				panic(err)
			}

			s := strings.Repeat(string(str[i-1]), n-1)

			res.WriteString(s)
		}
	}

	return res.String(), nil
}

func main() {
	ans, err := Unpack("qwe\\45")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(ans)
	}
}
