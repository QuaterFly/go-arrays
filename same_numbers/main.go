package main

import "fmt"

func main(){
	fmt.Println("hello from application")
	var numbers = [6] int64 {5,2,5,8,6,7}
	var saneNumbersAmount int8
	for index,value := range numbers{
		for i:=index + 1; i<len(numbers); i++ {
			if value==numbers[i]{
				saneNumbersAmount++
				break
			}
		}
	}
	if saneNumbersAmount!=0{
		saneNumbersAmount++
	}
	fmt.Println(saneNumbersAmount)
}
