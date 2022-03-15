Что выведет программа? Объяснить вывод программы. Рассказать про внутреннее устройство слайсов и что происходит при передачи их в качестве аргументов функции.

```go
package main

import (
	"fmt"
)

func main() {
	var s = []string{"1", "2", "3"}
	modifySlice(s)
	fmt.Println(s)
}

func modifySlice(i []string) {
	i[0] = "3"
	i = append(i, "4")
	i[1] = "5"
	i = append(i, "6")
}
```

Ответ:
```
Слайс внутри себя представляет некую структуру с полями int значений cap и len, а также указателем на массив.
При передаче слайса в функцию len и cap передадутся по значению,и указатель на массив также передаст свое значение.
Получим копию слайса, которая ссылается на тот же массив.
Сначала внутренний массив изменит нулевой элемент на 3 (для обоих слайсов), далее, т.к cap=len , функция append выделит
новый массив для слайса внутри функции ( в 2 раза больший чем до этого массив) и изменит 1 элемент и спокойно добавит новый элемент
со значеием 6. В основной функции мы выводим значение "внешнего" слайса, с тем самым первоначальным массивом значение которого стало [3,2,3]

```