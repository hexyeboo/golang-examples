package main

import (
	"fmt"
	"runtime"
	"time"
)

func Sqrt(x float64) float64 { // Упражнение на поиск того, как z близок к x. (Не до конца понял)
	z := float64(1)
	z -= (z*z - x) / (2 * z)
	return z
}

func main() {

	fmt.Println(Sqrt(2))

	fmt.Print("Go запущен на ") // Первый вариант применения свича.
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("macOS.")
	case "linux":
		fmt.Println("Linux.")
	default: //
		fmt.Printf("%s.\n", os)

		fmt.Println("Когда суббота?") // Другое применение свича
		today := time.Now().Weekday()
		switch time.Saturday {
		case today + 0:
			fmt.Println("Сегодня.")
		case today + 1:
			fmt.Println("Завтра.")
		case today + 2:
			fmt.Println("Послезавтра.")
		default:
			fmt.Println("Еще не скоро.")
		}
	}

	// Switch
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Доброе утро!")
	case t.Hour() < 17:
		fmt.Println("Доброго полудня!.")
	default:
		fmt.Println("Добрый вечер.")
	}

	// Defer
	func() {
		defer fmt.Println("от 5 до 1")
		fmt.Println("Отсчитай")
	}()

	// Stacking defers
	func() {
		fmt.Println("Считаю...")
		for i := 1; i < 6; i++ {
			defer fmt.Println(i)
		}
	}()
	fmt.Println("Готово.")
}
