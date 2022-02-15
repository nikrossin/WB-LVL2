package main

import "fmt"

type Vehicle interface {
	Get() string
}

type Plane struct {
	model        string
	modification string
	seats        int
}

func (p *Plane) Get() string {
	return fmt.Sprintf("Model:%v Mod:%v Seats:%v", p.model, p.modification, p.seats)
}

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
