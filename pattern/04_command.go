package main

import "fmt"

/*
	Паттерн превращает запросы в отдельные объекты,
	отделяя запросы от класса отправителя, позволяя осуществлять работу
	с запросами в рантайме, реализуя различные операции: выстраивание очередей
	команд, хранение истории, реализации отмены и другое.
	Например, может использоваться для разделения слоя графического интерфейса,
	от слоя бизнес-логики, которые будут общаться друг с другом посредством
	объектов команд: отправитель (графический интерфейс) будет вызывать нужную команду,
	а получатель (бизнес-логика) будет делать нужное действие.
	При этом детали будут скрытых от обоих узлов.
	Плюсы:
	- Убирается прямая связь между отправителями и исполнителями запросов
	- Позволяет удобно реализовывать различные операции: отмена и повтор запросов,
	отложенный запуск запросов, выстраивание очереди запросов.
	Минусы:
	- Усложняет код из-за необходимости реализации дополнительных классов
	Реализовать паттерн можно для создания взаимодействия между
	кнопкой на пульте (интерфейсом, отправителем) и телевизором (бизнес-логикой, приемником).
	Команды On и Off будут включать и выключать телевизор.
	Пример
	Разрабатываем библиотеку графического меню и хотите, чтобы пользователи могли использовать меню в разных
	приложениях, не меняя каждый раз код ваших классов.
*/

// Интерфейс команды
type Command interface {
	Execute()
}

// Кнопка, содержащая команду
type Button struct {
	cmd Command
}

func NewButton(cmd Command) *Button {
	return &Button{cmd}
}

// Выполнить команду кнопки
func (b *Button) Press() {
	b.cmd.Execute()
}

// Реализация команды для транспорта
type TurnOnCommand struct {
	v Vehicle
}

func NewTurnOnCommand(v Vehicle) *TurnOnCommand {
	return &TurnOnCommand{v}
}

// выполнение команды
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

// Интерфейс транспорта
type Vehicle interface {
	StartEngine()
	StopEngine()
}

// Автомобиль
type Car struct {
	isRunEngine bool
}

func NewCar() *Car {
	return &Car{}
}

// Запуск двигателя
func (v *Car) StartEngine() {
	fmt.Println("Engine is starting")
	v.isRunEngine = true
}

// Остановить двигатель
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
