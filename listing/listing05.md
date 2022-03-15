Что выведет программа? Объяснить вывод программы.

```go
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
```

Ответ:
```
error

Функция test вернет указатель на структуру, которая реализует интерфейс error. Указатель будет со значением nil.
Т.к customError реализует интерфейс err , то результат функции можно записать в переменную err. Аналогично заданию 3,
интерфейс не будет иметь значения [nil,nil], поскольку только динамическое значение nil., и в результате проверки условия,
будет true.

```
