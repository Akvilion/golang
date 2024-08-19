package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		// Скасовуємо операцію через 2 секунди
		time.Sleep(2 * time.Second)
		cancel()
	}()

	doSomething(ctx)
}

func doSomething(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Operation completed")
	case <-ctx.Done():
		// ctx.Done() канал сповіщає про скасування контексту
		fmt.Println("Operation canceled111:", ctx.Err()) // Оскільки ми скасовуємо операцію через 2 секунди, функція виведе:
	}
}
