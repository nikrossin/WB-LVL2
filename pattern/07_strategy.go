package main

import "fmt"

/*
  Стратегия - это поведенческий паттерн проектирования, который определяет семейство схожих алгоритмов и помещает каждый из них в собственный класс,
  после чего алгоритмы можно взаимозаменять прямо во время исполнения программы.
  Плюсы:
	- Возможность замены алгоритмов в рантайме
	- Отделение алгоритмов от остальной логики, сокрытие самих алгоритмов
  Минусы:
	- Усложнение кода, засчет введения дополнительных объектов
	- Клиент должен знать в чем состоит отличие алгоритмов, чтобы выбрать нужный
*/

// Интерфейс аптечки
type MedicalKit interface {
	Use(*Character)
}

// Тип аптечки - бинт
type Bandage struct {
	power int
}

func newBandage() MedicalKit {
	return &Bandage{10}
}

func (b *Bandage) Use(c *Character) {
	c.health += b.power
}

// Тип аптечки - большая аптечка
type FullKit struct {
	power int
}

func newFullKit() MedicalKit {
	return &FullKit{30}
}
func (f *FullKit) Use(c *Character) {
	c.health += f.power
}

type Character struct {
	name       string
	medicalKit MedicalKit
	health     int
}

// Метод применения аптечки
func (c *Character) UseMedicalKit() {
	c.medicalKit.Use(c)
}

func NewCharacter(name string) *Character {
	return &Character{name, nil, 5}
}

// "взять" аптечку
func (c *Character) GetKit(m MedicalKit) {
	c.medicalKit = m
}

func (c *Character) Get() string {
	return fmt.Sprintf("Name: %v Medical: %v Health: %v", c.name, c.medicalKit, c.health)
}

func main() {
	character := NewCharacter("solder")
	kit1 := newBandage()
	kit2 := newFullKit()

	fmt.Println(character.Get())
	character.GetKit(kit1)
	character.UseMedicalKit()
	fmt.Println(character.Get())
	character.GetKit(kit2)
	character.UseMedicalKit()
	fmt.Println(character.Get())

}
