package main

import "fmt"

/*
    поведенческий паттерн, который позволяет добавлять в программу новые операции, не изменяя классы объектов,
	над которыми эти операции могут выполняться.

	Используется, когда:
	- Когда новое поведение имеет смысл только для некоторых классов из существующей иерархии
	Плюсы:
	Упрощает добавление операций, работающих со сложными структурами объектов.
	Объединяет родственные операции в одном классе.

	Минусы:
	усложняет расширение иерархии классов, поскольку новые классы обычно требуют добавления нового метода visit для каждого посетителя
*/

type Device interface {
	GetModel() string
	Accept(Visitor)
}

// Смартфон
type SmartPhone struct {
	model  string
	cpu    string
	charge int
	isGPS  bool
}

func NewSmartPhone(model, cpu string, charge int, isGPS bool) *SmartPhone {
	return &SmartPhone{
		model:  model,
		cpu:    cpu,
		charge: charge,
		isGPS:  isGPS,
	}
}

func (p *SmartPhone) GetModel() string {
	return p.model
}

// Применить новое поведение
func (p *SmartPhone) Accept(v Visitor) {
	v.VisitPhone(p)
}

type SmartWatch struct {
	model         string
	cpu           string
	charge        int
	isPulseSensor bool
}

func NewSmartWatch(model, cpu string, charge int, isPulseSensor bool) *SmartWatch {
	return &SmartWatch{
		model:         model,
		cpu:           cpu,
		charge:        charge,
		isPulseSensor: isPulseSensor,
	}
}

func (w *SmartWatch) GetModel() string {
	return w.model
}

func (w *SmartWatch) Accept(v Visitor) {
	v.VisitWatch(w)
}

// Интерфейс встраиваемого класса
type Visitor interface {
	VisitWatch(*SmartWatch)
	VisitPhone(*SmartPhone)
}

// Индикация заряда устройства, как посетитель - встраиваемый модуль
type IndicatorCharge struct {
}

func NewIndicatorCharge() *IndicatorCharge {
	return &IndicatorCharge{}
}

func (i *IndicatorCharge) VisitWatch(w *SmartWatch) {
	fmt.Printf("Charge of %s watch is %d\n", w.GetModel(), w.charge)
}
func (i *IndicatorCharge) VisitPhone(p *SmartPhone) {
	fmt.Printf("Charge of %s phone is %d\n", p.GetModel(), p.charge)
}

// Проверка утсановлен ли доп модуль устройства, как посетитель - встраиваемый модуль
type AdditionalModules struct {
	is bool
}

func NewAdditionalModules() *AdditionalModules {
	return &AdditionalModules{false}
}
func (i *AdditionalModules) VisitWatch(w *SmartWatch) {
	i.is = w.isPulseSensor
}
func (i *AdditionalModules) VisitPhone(p *SmartPhone) {
	i.is = p.isGPS
}

func main() {
	phone := NewSmartPhone("Honor 20", "Kirin 780", 70, true)
	watch := NewSmartWatch("Apple series 3", "A4", 20, true)

	charge := NewIndicatorCharge()
	module := NewAdditionalModules()

	phone.Accept(charge)
	phone.Accept(module)
	fmt.Printf("Additional module in %s : %v\n", phone.GetModel(), module.is)

	watch.Accept(charge)
	watch.Accept(module)
	fmt.Printf("Additional module in %s : %v\n", watch.GetModel(), module.is)
}
