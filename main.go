package main

import (
	"fmt"
	"time"
)

type BaseEmploye struct {
	test string
}

type Employee struct {
	id   int
	name string
	BaseEmploye
}

func (s *Employee) setId(id *int) {

	s.id = *id
}

func main() {

	var x int
	x = 8
	y := 10
	fmt.Println(x)
	fmt.Println(y)
	m := make(map[string]int)

	m["Key"] = 6

	fmt.Println(m["Key"])

	s := []int{1, 2, 3, 45}

	for index, value := range s {

		fmt.Println(index, value)

	}

	c := make(chan string)

	go doSomething(c)

	<-c

}

func doSomething(c chan<- string) {

	fmt.Println("QLQ MALVADO MENOR")
	time.Sleep(3 * time.Second)

	c <- "Set Job"

}

func sum(values ...int) int {

	var totalSum int

	for _, valueNumber := range values {

		totalSum = +valueNumber

	}

	return totalSum

}
