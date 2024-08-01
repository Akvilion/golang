package main

import (
	"errors"
	"fmt"
)

func ProcessPayment(paymentMethod string, amount float64) (string, error) {
	switch paymentMethod {
	case "masterCard":
		fmt.Println("Обрабатывается платеж с MasterCard")
		return "Результат обработки платежа MasterCard", nil
	case "visa":
		fmt.Println("Обрабатывается платеж с Visa")
		return "Результат обработки платежа Visa", nil
	default:
		return "", errors.New("Неизвестный метод оплаты")
	}
}

func SendNotifications(paymentResult string) {
	fmt.Printf("Отправлены уведомления: %s\n", paymentResult)
}

func main() {
	res, err := ProcessPayment("masterCard", 100.50)
	if err != nil {
		fmt.Println("Ошибка обработки платежа:", err)
	} else {
		SendNotifications(res)
	}
}
