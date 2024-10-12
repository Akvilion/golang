package main

import (
	"fmt"
	"sync"
)

func R(ch chan int) {
	for i := range ch {
		fmt.Println(i)
	}
}

func One() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
		// defer R(ch)
	}()

	// нічого не виведе так як one завершить свою роботу не чекаючи відпрацювання горутин
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
	}()

}

func Two() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		close(ch)
	}()

	// Це точно відпрацює тому що ми тут проходимось по каналу без горутини

	for i := range ch {
		fmt.Println(i)
	}
}

func Three() {

	ch := make(chan int)

	// помилка fatal error: all goroutines are asleep - deadlock! через не закритий канал.
	// Це призводить до того, що основна горутина намагається отримати значення з каналу, але ніхто більше не надсилає дані, і канал залишається відкритим

	go func() {
		for i := 0; i < 1000; i++ {
			ch <- i
		}
		// close(ch)  // тому потрібно закрити канал
	}()

	// цикл for i := range ch читає значення з каналу. Цикл працює, доки канал не буде закритий
	for i := range ch {
		fmt.Println(i)
	}
}

func Four() {
	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go func(n int) { // горутини працюють паралельно, і порядок надсилання даних у канал не гарантується
			ch <- n
		}(i)

	}

	for i := range ch {
		fmt.Println(i)
	} // тому буде виводити 10 55 2 998 18 ... і ще й помилку бо канал не закритий

}

func Five() {
	var ws sync.WaitGroup

	// це працює ідеально, проте порядок викоду чисел різний 10 55 2 998 18 ...

	ch := make(chan int)
	for i := 0; i < 1000; i++ {
		go func(n int) { // горутини працюють паралельно, і порядок надсилання даних у канал не гарантується
			ch <- n
			defer ws.Done() // defer значить що після виконання горутити запуститься ws.Done() - зменшення на 1 (аналог finaly)
		}(i)

	}

	go func() {
		ws.Wait()
		close(ch)
	}()

	ws.Add(1000)
	for i := range ch {
		fmt.Println(i)
	} // тому буде виводити 10 55 2 998 ....
}

func main() {
	One()
	// Two()
	// Three()
	// Four()
	// Five()
}
