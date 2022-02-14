package main

import "fmt"

type Device interface {
	GetModel() string
	Accept(Visitor)
}

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

type Visitor interface {
	VisitWatch(*SmartWatch)
	VisitPhone(*SmartPhone)
}

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
