package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
=== Утилита grep ===

Реализовать утилиту фильтрации (man grep)

Поддержать флаги:
-A - "after" печатать +N строк после совпадения
-B - "before" печатать +N строк до совпадения
-C - "context" (A+B) печатать ±N строк вокруг совпадения
-c - "count" (количество строк)
-i - "ignore-case" (игнорировать регистр)
-v - "invert" (вместо совпадения, исключать)
-F - "fixed", точное совпадение со строкой, не паттерн
-n - "line num", печатать номер строки

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

func grep(args []string) error {
	if len(args) < 2 {
		return errors.New("usage: want string, filename, flag")
	}

	search := os.Args[1]
	if len(search) == 0 {
		return errors.New("invalid string")
	}

	filename := os.Args[2]
	strs, err := readFile(filename)
	if err != nil {
		return err
	}

	res := make([]string, 0, len(strs))

	var (
		A, B, C       int
		c, i, v, F, n bool
	)

	fs := flag.NewFlagSet(filename, flag.ContinueOnError)
	fs.IntVar(&A, "A", 0, "\"after\" print +N strings after match")
	fs.IntVar(&B, "B", 0, "\"before\" print +N strings before match")
	fs.IntVar(&C, "C", 0, "\"context\" (A+B) print ±N strings around match")
	fs.BoolVar(&c, "c", false, "\"count\" strings count")
	fs.BoolVar(&i, "i", false, "\"gnore-case\" ignore string's case")
	fs.BoolVar(&v, "v", false, "\"invert\" exclude instead of matching")
	fs.BoolVar(&F, "F", false, "\"fixed\" exact match with string (not pattern)")
	fs.BoolVar(&n, "n", false, "\"line num\", print number of string")
	fs.Parse(os.Args[3:])

	var needPrint = true

	switch {
	case A != 0:
		res = After(search, strs, A)
	case B != 0:
		res = Before(search, strs, B)
	case C != 0:
		res = Context(search, strs, C)
	case c:
		Count(search, strs)
		needPrint = false
	case i:
		res = IgnoreCase(search, strs)
	case v:
		res = Invert(search, strs)
	case F:
		res = Fixed(search, strs)
	case n:
		res = LineNum(search, strs)
	default:
		res = Def(search, strs)
	}

	if needPrint {
		Prnt(search, res, i, v)
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

	r := bufio.NewReader(fl)

	for {
		str, err := r.ReadString('\n')
		strs = append(strs, strings.Trim(str, "\n"))
		if err != nil {
			if err == io.EOF {
				break
			} else {
				return nil, err
			}
		}
	}

	return strs, nil
}

func After(search string, strs []string, a int) []string {
	res := make([]string, 0, len(strs))
	mp := make(map[string]struct{}, len(strs))

	for i := 0; i <= len(strs)-1; i++ {
		if strings.Contains(strs[i], search) {
			if a < len(strs[i:]) {
				for _, s := range strs[i : i+a+1] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			} else {
				for _, s := range strs[i:] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			}
		}
	}
	return res
}

func Before(search string, strs []string, b int) []string {
	res := make([]string, 0, len(strs))
	mp := make(map[string]struct{}, len(strs))

	for i := 0; i <= len(strs)-1; i++ {
		if strings.Contains(strs[i], search) {
			if b < len(strs[:i+1]) {
				for _, s := range strs[i-b : i+1] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			} else {
				for _, s := range strs[:i+1] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			}
		}
	}
	return res
}

func Context(search string, strs []string, c int) []string {
	res := make([]string, 0, len(strs))
	mp := make(map[string]struct{}, len(strs))

	for i := 0; i <= len(strs)-1; i++ {
		if strings.Contains(strs[i], search) {
			if c < len(strs[i:]) && c < len(strs[:i+1]) {
				for _, s := range strs[i-c : i+c+1] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			} else if c < len(strs[i:]) {
				for _, s := range strs[i : i+c+1] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			} else if c < len(strs[:i+1]) {
				for _, s := range strs[i-c : i+1] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			} else {
				for _, s := range strs[:] {
					if _, ok := mp[s]; !ok {
						res = append(res, s)
						mp[s] = struct{}{}
					}
				}
			}
		}
	}
	return res
}

func Count(search string, strs []string) {
	res := make([]string, 0, len(strs))

	for _, str := range strs {
		b := strings.Contains(str, search)
		if b {
			res = append(res, str)
		}
	}

	fmt.Println(len(res))
}

func IgnoreCase(search string, strs []string) []string {
	res := make([]string, 0, len(strs))
	for _, str := range strs {
		b := strings.Contains(strings.ToLower(str), strings.ToLower(search))
		if b {
			res = append(res, str)
		}
	}

	return res
}

func Invert(search string, strs []string) []string {
	res := make([]string, 0, len(strs))
	for _, str := range strs {
		b := strings.Contains(str, search)
		if !b {
			res = append(res, str)
		}
	}

	return res
}

func Fixed(search string, strs []string) []string {
	res := make([]string, 0, len(strs))
	for _, str := range strs {
		if str == search {
			res = append(res, str)
		}
	}

	return res
}

func LineNum(search string, strs []string) []string {
	var (
		ColorGreen string = "\u001b[32m"
		ColorBlue         = "\u001b[36m"
		ColorReset        = "\u001b[0m"
	)

	res := make([]string, 0, len(strs))

	for i, str := range strs {
		b := strings.Contains(str, search)
		if b {
			res = append(res, fmt.Sprintf("%s%d%s%s:%s%s", ColorGreen, i+1, ColorReset, ColorBlue, ColorReset, str))
		}
	}

	return res
}

func Def(search string, strs []string) []string {
	res := make([]string, 0, len(strs))
	for _, str := range strs {
		b := strings.Contains(str, search)
		if b {
			res = append(res, str)
		}
	}

	return res
}

func Prnt(search string, strs []string, I bool, v bool) {
	if v {
		for _, s := range strs {
			fmt.Println(s)
		}
		return
	}

	var (
		ColorRed   string = "\u001b[31;1m"
		ColorReset        = "\u001b[0m"
	)

	res := make([]string, len(strs))
	copy(res, strs)

	for k, str := range strs {
		for i := 0; i <= len(str); i++ {
			if i >= len(search) {
				if I {
					if strings.EqualFold(str[i-(len(search)):i], search) {
						res[k] = fmt.Sprintf("%s%s%s%s%s", str[:i-(len(search))], ColorRed, str[i-(len(search)):i], ColorReset, str[i:])
						str = res[k]
						i = i + len(ColorRed) + len(ColorReset)
					}
				} else {
					if str[i-(len(search)):i] == search {
						res[k] = fmt.Sprintf("%s%s%s%s%s", str[:i-(len(search))], ColorRed, search, ColorReset, str[i:])
						str = res[k]
						i = i + len(ColorRed) + len(ColorReset)
					}
				}
			}
		}
	}

	for _, s := range res {
		fmt.Println(s)
	}
}

func main() {
	if err := grep(os.Args[1:]); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
