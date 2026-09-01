package main

import (
	"fmt"
)

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}
	v2 = Vertex{X: 1}
	v3 = &Vertex{}
	p  = &Vertex{1, 2}
)

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	var k []int = primes[1:4]
	fmt.Println(k)
	fmt.Println(primes)

	fmt.Println(v1, p, v2, v3)

	i, j := 42, 2701

	p := &i
	fmt.Println(*p)
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37
	fmt.Println(j)

	fmt.Println(Vertex{1, 2})
	v := Vertex{1, 2}
	v.X = 4
	v.Y = 5
	fmt.Println(v.X)
	fmt.Println(v.Y)
}
