package mypkg // назва йде за назвою папки
import "fmt"

type MyInterface interface {
	MethodWithoutParameters()      // метод 1 інтерфейс
	MethodWithParameter(float64)   // метод 2 інтерфейс
	MethodWithReturnValue() string // метод 3 інтерфейс
}

type MyType int // тип підтримує методи з інтерфейсу

func (m MyType) MethodWithoutParameters() { // метод 1 інтерфейс
	fmt.Println("MethodWithoutParameters called")
}
func (m MyType) MethodWithParameter(f float64) { // метод 2 інтерфейс
	fmt.Println("MethodWithParameter called with", f)
}
func (m MyType) MethodWithReturnValue() string { // метод 3 інтерфейс
	return "Hi from MethodWithReturnValue"
}
func (my MyType) MethodNotInInterface() { // метод 4 не в інтерфейсі
	fmt.Println("MethodNotInInterface called")
}
