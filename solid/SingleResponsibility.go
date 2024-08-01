package main

import (
	"errors"
	"fmt"
)

func ProccessPaymentAndSendNotifications(paymentMethod string, amount float64) error {
	switch paymentMethod {
	case "masterCard":
		fmt.Printf("Обрабатывается платеж с MasterCard на сумму %.2f\n", amount)
		fmt.Println("Отправляются уведомления после обработки MasterCard")
	case "visa":
		fmt.Printf("Обрабатывается платеж с Visa на сумму %.2f\n", amount)
		fmt.Println("Отправляются уведомления после обработки Visa")
	default:
		return errors.New("Неизвестный метод оплаты")
	}

	return nil
}

func main() {
	// Пример использования функции
	err := ProccessPaymentAndSendNotifications("masterCard", 100.50)
	if err != nil {
		fmt.Println("Ошибка обработки платежа:", err)
	}
}
