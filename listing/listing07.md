Что выведет программа? Объяснить вывод программы.

```go
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func asChan(vs ...int) <-chan int {
	c := make(chan int)

	go func() {
		for _, v := range vs {
			c <- v
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		}

		close(c)
	}()
	return c
}

func merge(a, b <-chan int) <-chan int {
	c := make(chan int)
	go func() {
		for {
			select {
			case v := <-a:
				c <- v
			case v := <-b:
				c <- v
			}
		}
	}()
	return c
}

func main() {

	a := asChan(1, 3, 5, 7)
	b := asChan(2, 4 ,6, 8)
	c := merge(a, b )
	for v := range c {
		fmt.Println(v)
	}
}
```

Ответ:
```
1, 2, 3, 4, 5, 6, 7, 8, 0, 0, 0............
Функция merge в селекте ожидает поступления данных в канал а и b, из-за задержки данные поступают поочередно и посылаются в канал с.
В main главная рутина читает данные из канала с и выводит. После отправки всех значений в каналы а и b и их считывания,
каналы закрываются в функциях asChan, после каналы закрыты, а данные в них "всегда есть", это значения типа данных канала по умолчанию(int),
Select больше не блокируется в ожидание данных в каналах, а читает значения по умолчанию, чтобы этого избежать нужно
внутри case проверять закрыть ли канал и делать break

```
