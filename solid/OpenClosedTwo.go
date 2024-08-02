package main

import (
	"errors"
	"fmt"
)

// суть принципу якщо ми схочемо добавити новий метод платежу треба буде міняти тіло ProcessPayment

// Функция processPayment принимает информацию о платеже и обрабатывает его.
func ProcessPayment(paymentMethod string, amount float64, accountStatus *string) (string, error) {
	switch paymentMethod {
	case "masterCard":
		fmt.Println("Обрабатывается платеж с MasterCard")
		return "Результат обработки платежа MasterCard", nil
	case "visa":
		fmt.Println("Обрабатывается платеж с Visa")
		return "Результат обработки платежа Visa", nil
	case "paypal":
		fmt.Println("Обрабатывается платеж через PayPal")
		if accountStatus == nil {
			return "", errors.New("Отстуствует accountStatus")
		}
		if err := CheckPayPalAccountStatus(*accountStatus); err != nil {
			return "", err
		}
		return "Результат обработки платежа paypal", nil
	default:
		return "", errors.New("Неизвестный метод оплаты")
	}
}

// Функция checkPayPalAccountStatus проверяет статус аккаунта пользователя PayPal.
func CheckPayPalAccountStatus(accountStatus string) error {
	if accountStatus == "active" {
		return nil
	}
	if accountStatus == "inactive" {
		return errors.New("Аккаунт PayPal неактивен")
	}

	return errors.New("Неизвестный статус PayPal аккаунта")
}

// Функция sendNotifications отправляет уведомления получателю.
func SendNotifications(paymentResult string) {
	// Отправка уведомлений по результатам
	fmt.Printf("Отправлены уведомления: %s\n", paymentResult)
}

func main() {
	accountStatus := "active"
	// Пример использования функций
	res, err := ProcessPayment("masterCard", 100.50, nil)
	if err != nil {
		fmt.Println("Ошибка обработки платежа:", err)
	} else {
		SendNotifications(res)
	}

	// Пример использования функций
	res, err = ProcessPayment("paypal", 100.50, &accountStatus) // тут &accountStatus accountStatus должен быть не просто string, а *string. То есть он должен быть указателем на string, чтобы мы могли передавать туда и string, и nil в качестве значения.
	if err != nil {
		fmt.Println("Ошибка обработки платежа:", err)
	} else {
		SendNotifications(res)
	}
}
