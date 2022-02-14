package main

import "fmt"

type Command interface {
	Execute()
}

type Button struct {
	cmd Command
}

func NewButton(cmd Command) *Button {
	return &Button{cmd}
}

func (b *Button) Press() {
	b.cmd.Execute()
}

type TurnOnCommand struct {
	v Vehicle
}

func NewTurnOnCommand(v Vehicle) *TurnOnCommand {
	return &TurnOnCommand{v}
}

func (to *TurnOnCommand) Execute() {
	to.v.StartEngine()
}

type TurnOffCommand struct {
	v Vehicle
}

func NewTurnOffCommand(v Vehicle) *TurnOffCommand {
	return &TurnOffCommand{v}
}

func (to *TurnOffCommand) Execute() {
	to.v.StopEngine()
}

type Vehicle interface {
	StartEngine()
	StopEngine()
}

type Car struct {
	isRunEngine bool
}

func NewCar() *Car {
	return &Car{}
}

func (v *Car) StartEngine() {
	fmt.Println("Engine is starting")
	v.isRunEngine = true
}

func (v *Car) StopEngine() {
	fmt.Println("Engine is stopping")
	v.isRunEngine = false
}

func main() {
	car := NewCar()
	onCmd := NewTurnOnCommand(car)
	offCmd := NewTurnOffCommand(car)
	btnOn := NewButton(onCmd)
	btnOff := NewButton(offCmd)

	btnOn.Press()
	fmt.Println(car.isRunEngine)
	btnOff.Press()
	fmt.Println(car.isRunEngine)

}
