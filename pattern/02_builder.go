// уменьшения размера конструктора создаваемого объекта;
// создания немного отличающихся в значениях, но одинаковых в конструкции объектов.
/*
 Паттерн Строитель предлагает вынести конструирование объекта за пределы его собственного класса, поручив это дело
 отдельным объектам, называемым строителями.

плюсы:
 Создание объекта пошаговое.
 Позволяет использовать один и тот же код для создания различных но схожих объектов.
минусы:
 Усложняет код программы из-за введения дополнительных классов.

*/

package main

import "fmt"

// Класс объекта для конструирования
type Plane struct {
	brand string
	model string
	seats int
}

func NewPlane(br, md string, s int) *Plane {
	return &Plane{br, md, s}
}

// Выбор строителя
func SetBuilder(bType string) Builder {
	if bType == "Boing" {
		return NewBoingBuilder()
	} else if bType == "Airbus" {
		return NewAirbusBuilder()
	}
	return nil
}

// Интерфейс строителя
type Builder interface {
	setBrand()
	setModel(string)
	setSeats(int)
	build() *Plane
}

// Строитель Airbus
type AirbusBuilder struct {
	brand string
	model string
	seats int
}

func NewAirbusBuilder() *AirbusBuilder {
	return &AirbusBuilder{}
}

// Методы для "строительства" пошагово объекта
func (b *AirbusBuilder) setBrand() {
	b.brand = "Airbus"
}
func (b *AirbusBuilder) setModel(m string) {
	b.model = m
}
func (b *AirbusBuilder) setSeats(s int) {
	b.seats = s
}

// Завершить сборку объекта"
func (b *AirbusBuilder) build() *Plane {
	return NewPlane(b.brand, b.model, b.seats)
}

// СТроитель Boing
type BoingBuilder struct {
	brand string
	model string
	seats int
}

func NewBoingBuilder() *BoingBuilder {
	return &BoingBuilder{}
}
func (b *BoingBuilder) setBrand() {
	b.brand = "Boing"
}
func (b *BoingBuilder) setModel(m string) {
	b.model = m
}
func (b *BoingBuilder) setSeats(s int) {
	b.seats = s
}
func (b *BoingBuilder) build() *Plane {
	return NewPlane(b.brand, b.model, b.seats)
}

func main() {
	airbus := SetBuilder("Airbus")
	airbus.setBrand()
	airbus.setModel("A320")
	airbus.setSeats(180)
	plane1 := airbus.build()

	boing := SetBuilder("Boing")
	boing.setBrand()
	boing.setModel("737-800")
	boing.setSeats(167)
	plane2 := boing.build()

	fmt.Println(plane1, "", plane2)

}
