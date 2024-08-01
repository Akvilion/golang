package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Pay(amount float64) (string, error)
}

type MasterCardPayment struct{}

func (m *MasterCardPayment) Pay(amount float64) (string, error) {
	fmt.Println("Обрабатывается платеж с MasterCard")
	// Дополнительный код для обработки платежа с MasterCard
	return "Результат обработки платежа MasterCard", nil
}

type VisaPayment struct{}

func (v *VisaPayment) Pay(amount float64) (string, error) {
	fmt.Println("Обрабатывается платеж с Visa")
	// Дополнительный код для обработки платежа с Visa
	return "Результат обработки платежа Visa", nil
}

type PaypalPayment struct{}

func (p *PaypalPayment) Pay(amount float64) (string, error) {
	fmt.Println("Обрабатывается платеж через PayPal")
	// Дополнительный код для обработки платежа с PayPal
	return "Результат обработки платежа PayPal", nil
}

// Функция checkAccountStatus проверяет статус аккаунта пользователя.
func (p *PaypalPayment) CheckAccountStatus(accountStatus string) error {
	if accountStatus == "active" {
		return nil
	}
	if accountStatus == "inactive" {
		return errors.New("Аккаунт PayPal неактивен")
	}

	return errors.New("Неизвестный статус PayPal аккаунта")
}

func SendNotifications(paymentResult string) {
	fmt.Printf("Отправлены уведомления: %s\n", paymentResult)
}

// Функция processPayment принимает информацию о платеже и обрабатывает его.
func ProcessPayment(p Payer, amount float64) (string, error) {
	return p.Pay(amount)
}

func main() {
	// Пример использования функций
	masterCard := &MasterCardPayment{}
	res, err := ProcessPayment(masterCard, 100.50)
	if err != nil {
		fmt.Println("Ошибка обработки платежа:", err)
	} else {
		SendNotifications(res)
	}

	// Пример использования функций
	payPal := &PaypalPayment{}
	accountStatus := "active"
	err = payPal.CheckAccountStatus(accountStatus)
	if err != nil {
		fmt.Println("Ошибка проверки аккаунта:", err)
	} else {
		res, err = ProcessPayment(payPal, 100.50)
		if err != nil {
			fmt.Println("Ошибка обработки платежа:", err)
		} else {
			SendNotifications(res)
		}
	}
}
