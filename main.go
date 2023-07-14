package main

import (
	"fmt"
	"math/rand"
)

func swap(x, y string) (string, string) {
	return y, x
}

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

const (
	_      = iota
	KB int = 1 << (10 * iota) // 1
	MB                        // 2
	GB
	TB
)

func exe4() {

	fmt.Printf("%d  \n", KB)
	fmt.Printf("%d  \n", MB)
	fmt.Printf("%d  \n", GB)
	fmt.Printf("%d  \n", TB)
}

func exe5() {
	var zero int
	short := 9
	const (
		_ = iota
		a
		b
	)
	fmt.Print(zero)
	fmt.Print(short)
	fmt.Print(a)
	fmt.Print(b)
}

func main() {
	const (
		name = "David"
		age  = 23
	)
	t := "Roffel"
	fmt.Println("Hello World!")

	const i = int(2)
	var j = rand.Intn(4)

	t = "oinhoin"
	fmt.Printf("%s is %d years old vs %s.", name, age, t)
	fmt.Print(i * j)

	a, b := swap("hello", "world")
	fmt.Print("\n", a, b)

	fmt.Printf("\n %d \t %b", 1<<7, 1<<7)

	exe4()
}
