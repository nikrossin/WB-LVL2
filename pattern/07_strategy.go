package main

import "fmt"

type MedicalKit interface {
	Use(*Character)
}

type Bandage struct {
	power int
}

func newBandage() MedicalKit {
	return &Bandage{10}
}

func (b *Bandage) Use(c *Character) {
	c.health += b.power
}

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

func (c *Character) UseMedicalKit() {
	c.medicalKit.Use(c)
}

func NewCharacter(name string) *Character {
	return &Character{name, nil, 5}
}

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
