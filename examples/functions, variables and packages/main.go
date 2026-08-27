package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func add(x int, y int) int { // Можно использовать просто (x, y int). Для упрощения.
	return x + y
}
func swap(x, y string) (string, string) { // Функция свапа, помогает менять местами две строчки
	return y, x
}
func split(sum int) (int, int) {
	x := sum * 4 / 9
	y := sum - x
	return x, y
}
func main() { // Здесь играюсь с различными функицями, пробую использовать то, что наимпортил.
	functions() // Пока убрал в коммент, чтобы работала только часть flow.
	flow()
	sqrt()
}
func functions() {
	fmt.Println("Hello, Slavique")
	fmt.Println("The time is", time.Now())
	fmt.Println("My favorite number is", rand.Intn(10))
	fmt.Println(add(25, 25))

	a, b := swap("hello", "world")
	fmt.Println(a, b)

	splitResultX, splitResultY := split(10)
	fmt.Println(splitResultX, splitResultY)

	var c, python bool
	var i, java int
	fmt.Println(i, c, python, java)

	var k, l int = 1, 2 // Переменные с инициализаторами, приравнимаем переменным свои свойства.
	var swift, cSharp, javaScript = true, false, "yes!"
	fmt.Println(k, l, swift, cSharp, javaScript)
}
func flow() {

	sum := 5
	for i := 0; i < 10; i++ { // Выводит числа от 0 до 9
		sum += i
	}
	fmt.Println(sum)

	for sum < 1000 {
		sum += sum // Удваивает sum пока он меньше 100
	}
	fmt.Println(sum)
}
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}
