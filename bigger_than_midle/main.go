package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var number, avarage int64
	var class = [20]int64{}
	for i := 0; i < 20; i++ {
		class[i] = rand.Int63n(50) + 150
		avarage += class[i]
	}
	avarage = avarage / 20
	for i := 0; i < 20; i++ {
		if class[i] > avarage {
			number++
		}
	}
	fmt.Println("Рост учащихся:",class)
	fmt.Println("Средний рост в классе:", avarage)
	fmt.Println("Кол-во учащихся чей рост выше среднего в классе:", number)
}
