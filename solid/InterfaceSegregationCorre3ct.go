package main

import (
	"errors"
	"fmt"
)

// Интерфейс для метода оплаты
type Payer interface {
	Pay() (string, error)
}

// Интерфейс для предварительной проверки перед оплатой
type PreprocessChecker interface {
	PreprocessCheck() error
}

// Интерфейс, объединяющий методы из Payer и PreprocessChecker
type PaymentProcessor interface {
	Payer
	PreprocessChecker
}

// Структура для MasterCard платежа
type MasterCardPayment struct {
	Amount int
}

func (m *MasterCardPayment) Pay() (string, error) {
	fmt.Printf("Обрабатывается платеж с MasterCard c суммой %v\n", m.Amount)
	// Дополнительный код для обработки платежа с MasterCard
	return "Результат обработки платежа MasterCard", nil
}

func (m *MasterCardPayment) PreprocessCheck() error {
	return nil
}

// Структура для Visa платежа
type VisaPayment struct {
	Amount int
}

func (v *VisaPayment) Pay() (string, error) {
	fmt.Printf("Обрабатывается платеж с Visa c суммой %v\n", v.Amount)
	// Дополнительный код для обработки платежа с Visa
	return "Результат обработки платежа Visa", nil
}

func (m *VisaPayment) PreprocessCheck() error {
	return nil
}

// Структура для PayPal платежа
type PaypalPayment struct {
	Amount        int
	AccountStatus string
}

func (p *PaypalPayment) Pay() (string, error) {
	fmt.Printf("Обрабатывается платеж через PayPal c суммой %v\n", p.Amount)
	// Дополнительный код для обработки платежа с PayPal
	return "Результат обработки платежа PayPal", nil
}

func (p *PaypalPayment) PreprocessCheck() error {
	if p.AccountStatus == "active" {
		return nil
	}
	if p.AccountStatus == "inactive" {
		return errors.New("Аккаунт PayPal неактивен")
	}

	return errors.New("Неизвестный статус PayPal аккаунта")
}

// Функция sendNotifications отправляет уведомления получателю.
func SendNotifications(paymentResult string) {
	// Отправка уведомлений по результатам
	fmt.Printf("Отправлены уведомления: %s\n", paymentResult)
}

func PreprocessCheck(pc PreprocessChecker) error {
	return pc.PreprocessCheck()
}

// Функция processPayment принимает информацию о платеже и обрабатывает его.
func ProcessPayment(p Payer) (string, error) {
	return p.Pay()
}

func main() {
	payPal := &PaypalPayment{AccountStatus: "active", Amount: 500}
	masterCard := &MasterCardPayment{Amount: 200}

	payments := []PaymentProcessor{payPal, masterCard}

	for _, payment := range payments {
		err := PreprocessCheck(payment)
		if err != nil {
			fmt.Println("Ошибка первичной проверки платежа:", err)
			continue
		}

		res, err := ProcessPayment(payment)
		if err != nil {
			fmt.Println("Ошибка обработки платежа:", err)
			continue
		}

		SendNotifications(res)
	}
}
