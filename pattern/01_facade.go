package pattern

import "fmt"

type bankEmitter struct {
}

type bankAcquire struct {
}

type paymentSystem struct {
}

type paymentDetails struct {
}

func (payDet *paymentDetails) Set() {
	fmt.Println("Payment details payment details are set")
}
func (payDet *paymentDetails) TransmitToAcquire() {
	fmt.Println("Data transmit to Acquire Bank")
}

func (paySys *paymentSystem) RequestWriteOffToEmitter() {
	fmt.Println("Request for write-off funds: OK!")
}

func (paySys *paymentSystem) RequestCreditToEmitter() {
	fmt.Println("Request for credit funds: OK!")
}
func (bank *bankAcquire) CheckPayment() {
	fmt.Println("Check payment details")
}
func (bank *bankAcquire) TransmitToPaySystem() {
	fmt.Println("Payment transmit toPaySystem")
}
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
