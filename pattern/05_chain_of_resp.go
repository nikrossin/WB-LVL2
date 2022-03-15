package main

import "fmt"

/*
Идея шаблона заключается в организации рекуррентного конвейера из обработчиков, в котором каждый обработчик
может либо обработать поступившее сообщение (например, только сообщения определенного типа), либо
делегировать обработку следующему в конвейере обработчику. Возможен еще вариант обработки и последующей передачи.
При этом, клиенту, чтобы инициировать обработку того или иного сообщения достаточно лишь передать
его первому в конвейере обработчику.
Плюсы:
	- Разнесение клиента и обработчиков, уменьшение их зависимости
	- Реализация принципа единственной ответственности
Минусы:
	- Создание дополнительных объектов, усложнение кода
	- Запрос может быть не обработан

*/

//Интерфейс каждого департамента
type Department interface {
	Execute(*Cargo)
	SetNext(Department)
}

// Отделение почты
type PostOffice struct {
	next Department
}

// выполняет свои действие отделение согласно критериям конвеера
func (p *PostOffice) Execute(cargo *Cargo) {
	if !cargo.isPaid {
		cargo.isPaid = true
		fmt.Println("PostOffice: Cargo is paid")
	}
	fmt.Println("PostOffice: Cargo has been successfully sent")
	p.next.Execute(cargo)

}

// Установить след обработчика
func (p *PostOffice) SetNext(d Department) {
	p.next = d
}

type Security struct {
	next Department
}

func (p *Security) Execute(cargo *Cargo) {
	if !cargo.isCheck {
		cargo.isCheck = true
		fmt.Println("Security: Cargo is checked!")
	}
	fmt.Println("Security: Cargo has been sent to next point")
	p.next.Execute(cargo)

}

func (p *Security) SetNext(d Department) {
	p.next = d
}

type SortingCenter struct {
	next Department
}

func (p *SortingCenter) Execute(cargo *Cargo) {
	if cargo.isSorting {
		fmt.Println("SortingCenter: Cargo has already been sorted in this center - additional verification")
		return
	}
	cargo.isSorting = true
	fmt.Println("SortingCenter: Cargo has been sorted")
	fmt.Println("SortingCenter: Cargo has been sent to next point")
	p.next.Execute(cargo)

}

func (p *SortingCenter) SetNext(d Department) {
	p.next = d
}

type Cargo struct {
	isPaid    bool
	isCheck   bool
	isSorting bool
}

func main() {
	cargo := &Cargo{}

	post := &PostOffice{}
	security1 := &Security{}
	security2 := &Security{}
	sort := &SortingCenter{}

	post.SetNext(security1)
	security1.SetNext(sort)
	sort.SetNext(security2)
	security2.SetNext(sort)

	post.Execute(cargo)

}
