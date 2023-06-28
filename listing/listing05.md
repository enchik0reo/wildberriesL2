Что выведет программа? Объяснить вывод программы.

package main

type customError struct {
	msg string
}

func (e *customError) Error() string {
	return e.msg
}

func test() *customError {
	{
		// do something
	}
	return nil
}

func main() {
	var err error
	err = test()
	if err != nil {
		println("error")
		return
	}
	println("ok")
}

Ответ:
error

error это интерфейс с одним методом Error() string.
В го утиная типизация, поэтому customError удовлетворяет этому интерфейсу.
Функция test() возвращает значение nil, но типа customError, следовательно функция main получит вместо <nil><nil>:
*main.customError <nil>  и неравенство с <nil> будет true