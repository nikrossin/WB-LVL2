package main

import "fmt"

type Plane struct {
	brand string
	model string
	seats int
}

func NewPlane(br, md string, s int) *Plane {
	return &Plane{br, md, s}
}

func SetBuilder(bType string) Builder {
	if bType == "Boing" {
		return NewBoingBuilder()
	} else if bType == "Airbus" {
		return NewAirbusBuilder()
	}
	return nil
}

type Builder interface {
	setBrand()
	setModel(string)
	setSeats(int)
	build() *Plane
}

type AirbusBuilder struct {
	brand string
	model string
	seats int
}

func NewAirbusBuilder() *AirbusBuilder {
	return &AirbusBuilder{}
}
func (b *AirbusBuilder) setBrand() {
	b.brand = "Airbus"
}
func (b *AirbusBuilder) setModel(m string) {
	b.model = m
}
func (b *AirbusBuilder) setSeats(s int) {
	b.seats = s
}

func (b *AirbusBuilder) build() *Plane {
	return NewPlane(b.brand, b.model, b.seats)
}

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
