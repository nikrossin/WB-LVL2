package main

import "fmt"

/*

	Фабричный метод —это класс, в котором есть один метод с большим условным оператором, выбирающим создаваемый продукт.
	Применимость:
	- Когда заранее неизвестны типы и зависимости объектов, с которыми должен работать ваш код.
	Фабричный метод отделяет код производства продуктов от остального кода, который эти продукты использует.
	- Когда вы хотите дать возможность пользователям расширять части вашего фреймворка или библиотеки.
	Преимущества:
	- Выделяет код производства продуктов в одно место, упрощая поддержку кода.
	- Упрощает добавление новых продуктов в программу.
	Кросс-платформенная программа может показывать одни и те же элементы интерфейса, выглядящие чуточку по-другому в
	различных операционных системах. В такой программе важно, чтобы все создаваемые элементы всегда соответствовали
	текущей операционной системе. Вы бы не хотели, чтобы программа, запущенная на Windows, вдруг начала показывать
	чекбоксы в стиле macOS.

*/
// Интерфейс транспортного средства
type Vehicle interface {
	Get() string
}

// Класс самолета
type Plane struct {
	model        string
	modification string
	seats        int
}

func (p *Plane) Get() string {
	return fmt.Sprintf("Model:%v Mod:%v Seats:%v", p.model, p.modification, p.seats)
}

// Производство моели самолета
type Airbus struct {
	Plane
}

func NewAirbus() Vehicle {
	return &Airbus{
		Plane{
			"A320",
			"NEO",
			188,
		},
	}
}

// Модель Боинга
type Boing struct {
	Plane
}

func NewBoing() Vehicle {
	return &Boing{
		Plane{
			"737",
			"MAX 8",
			167,
		},
	}
}

// Фабрика производства по заданному типу
func Factory(pType string) (Vehicle, error) {
	switch pType {
	case "Airbus":
		return NewAirbus(), nil
	case "Boing":
		return NewBoing(), nil
	default:
		return nil, fmt.Errorf("Error type of Vehicle")
	}

}

func main() {
	airbus, _ := Factory("Airbus")
	boing, _ := Factory("Boing")
	fmt.Println(airbus.Get())
	fmt.Println(boing.Get())
}
