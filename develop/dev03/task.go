package main

/*
=== Утилита sort ===

Отсортировать строки (man sort): на входе подается файл из несортированными
строками, на выходе — файл с отсортированными.

Основное

Поддержать ключи

-k — указание колонки для сортировки (слова в строке могут выступать в качестве колонок,
по умолчанию разделитель - пробел)
-n — сортировать по числовому значению
-r — сортировать в обратном порядке
-u — не выводить повторяющиеся строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func sortFile(args []string) error {
	if len(args) < 1 {
		return errors.New("you must pass a file name")
	}

	filename := os.Args[1]
	strs, err := readFile(filename)
	if err != nil {
		return err
	}

	var (
		k       int
		n, r, u bool
	)

	fs := flag.NewFlagSet(filename, flag.ContinueOnError)
	fs.IntVar(&k, "k", 0, "sort by column")
	fs.BoolVar(&n, "n", false, "sort by number")
	fs.BoolVar(&r, "r", false, "sort in reverse")
	fs.BoolVar(&u, "u", false, "sort by trimming duplicate strings")
	fs.Parse(os.Args[2:])

	switch {
	case k != 0 || n || r || u:
		if k != 0 {
			strs = ColumnSort(strs, k)
		}

		if n {
			strs = NumSort(strs, k)
		}

		if r {
			strs = ReverseSort(strs, k, n)
		}

		if u {
			strs = UniqSort(strs, k, n, r)
		}

	default:
		strs = DefaultSort(strs)
	}

	if err = createFile(filename, strs); err != nil {
		return err
	}
	return nil
}

func readFile(filename string) ([]string, error) {
	fl, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fl.Close()

	var strs []string
	sc := bufio.NewScanner(fl)

	for sc.Scan() {
		str := sc.Text()
		strs = append(strs, str)
	}

	return strs, nil
}

func ColumnSort(strs []string, k int) []string {
	if len(strs) < 2 {
		return strs
	}
	if k < 2 {
		res := make([]string, len(strs))
		copy(res, strs)
		sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
		return res
	}

	var res = make([]string, len(strs))
	copy(res, strs)
	var midleRes []string
	var itog = make([]string, 0, len(strs))
	var mp = make(map[string][]string)

	for i := 0; i <= len(res)-1; i++ {
		strin := strings.Split(res[i], " ")
		if len(strin) >= k {
			ss := strin[k-1]
			if _, ok := mp[ss]; !ok {
				midleRes = append(midleRes, ss)
				mp[ss] = append(mp[ss], res[i])
				copy(res[i:], res[i+1:])
				i--
				res = res[:len(res)-1]
			} else {
				mp[ss] = append(mp[ss], res[i])
				copy(res[i:], res[i+1:])
				i--
				res = res[:len(res)-1]
			}
		}
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	sort.Slice(midleRes, func(i, j int) bool { return midleRes[i] < midleRes[j] })

	itog = append(itog, res...)

	for _, j := range midleRes {
		itog = append(itog, mp[j]...)
	}

	return itog
}

func NumSort(strs []string, k int) []string {
	if len(strs) < 2 {
		return strs
	}

	if k < 2 {
		k = 1
	}

	var res = make([]string, len(strs))
	copy(res, strs)
	var midleRes []int
	var itog = make([]string, 0, len(strs))
	var mp = make(map[int][]string)

	for i := 0; i <= len(res)-1; i++ {
		strin := strings.Split(res[i], " ")
		if len(strin) >= k {
			ss := strin[k-1]
			if ss[0] >= '0' && ss[0] <= '9' {
				var runs []rune
				for _, r := range ss {
					if r >= '0' && r <= '9' {
						runs = append(runs, r)
					} else {
						break
					}
				}

				num, err := strconv.Atoi(string(runs))
				if err != nil {
					fmt.Fprintln(os.Stderr, err)
				}

				if _, ok := mp[num]; !ok {
					midleRes = append(midleRes, num)
					mp[num] = append(mp[num], res[i])
					copy(res[i:], res[i+1:])
					i--
					res = res[:len(res)-1]
				} else {
					mp[num] = append(mp[num], res[i])
					copy(res[i:], res[i+1:])
					i--
					res = res[:len(res)-1]
				}
			}
		}
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	sort.Slice(midleRes, func(i, j int) bool { return midleRes[i] < midleRes[j] })

	itog = append(itog, res...)

	for _, j := range midleRes {
		itog = append(itog, mp[j]...)
	}

	return itog
}

func ReverseSort(strs []string, k int, n bool) []string {
	if len(strs) < 2 {
		return strs
	}

	if k != 0 || n {
		res := make([]string, 0, len(strs))
		for i := len(strs) - 1; i >= 0; i-- {
			res = append(res, strs[i])
		}
		return res
	}

	res := make([]string, len(strs))
	copy(res, strs)
	sort.Slice(res, func(i, j int) bool { return res[i] > res[j] })
	return res
}

func UniqSort(strs []string, k int, n bool, r bool) []string {
	if len(strs) < 2 {
		return strs
	}

	res := make([]string, 0, len(strs))
	mp := make(map[string]struct{})

	for _, j := range strs {
		if _, ok := mp[j]; !ok {
			res = append(res, j)
			mp[j] = struct{}{}
		}
	}

	if k != 0 || n || r {
		return res
	}

	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func DefaultSort(strs []string) []string {
	res := make([]string, len(strs))
	copy(res, strs)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return res
}

func createFile(filename string, strs []string) error {
	fl, err := os.Create("sorted " + filename)
	if err != nil {
		return err
	}
	defer fl.Close()

	for i, j := range strs {
		if i == len(strs)-1 {
			_, err := fl.WriteString(j)
			if err != nil {
				return err
			}
		} else {
			_, err := fl.WriteString(j + "\n")
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func main() {
	if err := sortFile(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
