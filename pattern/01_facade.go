package pattern

//скрывает за фасадом сложность инициализации отдельных служб для клиента

/*
	Плюсы:
	- меньше кода, меньше ошибок, быстрее разработка.
	- упрощение работы с подсистемой для клиента
	Минусы:
	- дополнительная разработка.
	- Нужно хорошо продумать реализуемый набор интерфейсов для клиента
	Реальные примеры:
	Используя паттерн «Фасад», реализуем унифицированный интерфейс к некоторой подсистеме авторизации пользователей.
	Сама подсистема авторизации (в данном примере), безусловно не претендует на «сложную систему»
*/
import "fmt"

type bankEmitter struct {
}

type bankAcquire struct {
}

type paymentSystem struct {
}

type paymentDetails struct {
}

// Платежные данные
func (payDet *paymentDetails) Set() {
	fmt.Println("Payment details payment details are set")
}

// Передача транзакции в банк эквайер
func (payDet *paymentDetails) TransmitToAcquire() {
	fmt.Println("Data transmit to Acquire Bank")
}

// Списание средств у банка эмиттера
func (paySys *paymentSystem) RequestWriteOffToEmitter() {
	fmt.Println("Request for write-off funds: OK!")
}

// Зачисление денежных средств в другой банк эмиттер
func (paySys *paymentSystem) RequestCreditToEmitter() {
	fmt.Println("Request for credit funds: OK!")
}

// Проверка платежных данных
func (bank *bankAcquire) CheckPayment() {
	fmt.Println("Check payment details")
}

// Транзакция переходит "в руки" платежной системы"
func (bank *bankAcquire) TransmitToPaySystem() {
	fmt.Println("Payment transmit toPaySystem")
}

// Прием денежных средств
func (bank *bankEmitter) DepositFunds() {
	fmt.Println("Deposit funds")
}

type PaymentFacade struct {
	bankEm     *bankEmitter
	bankAc     *bankAcquire
	paySys     *paymentSystem
	payDetails *paymentDetails
}

func NewPaymentFacade() *PaymentFacade {
	return &PaymentFacade{
		&bankEmitter{},
		&bankAcquire{},
		&paymentSystem{},
		&paymentDetails{},
	}

}

func (f *PaymentFacade) TransmitFunds() {
	f.payDetails.Set()
	f.payDetails.TransmitToAcquire()
	f.bankAc.CheckPayment()
	f.bankAc.TransmitToPaySystem()
	f.paySys.RequestWriteOffToEmitter()
	f.paySys.RequestCreditToEmitter()
	f.bankEm.DepositFunds()
}

func main() {
	payFacade := NewPaymentFacade()

	payFacade.TransmitFunds()
}
