package main

import "fmt"

func proccessPayment(cardType string, amount int) (string, error) {
	switch cardType {
	case "visa":
		fmt.Println("платіж обробляється")
		return fmt.Sprintf("visa %d", amount), nil
	}
	return "", nil
}

func main() {
	res, err := proccessPayment("visa", 100)
	if err != nil {
		fmt.Println("помилка")
	}
	fmt.Println(res)

}
