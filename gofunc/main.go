package main

import (
	"fmt"
	"time"
)

type Fruit struct {
	Name string
}

//go run --race main.go
func main() {
	//1.has race
	Race1()
	//2.has race
	Race2()
	//3.no race
	NoRace()

}

func Race1() {
	f := &Fruit{Name: "Apple"}
	go func() {
		f.Name = "Pear"
	}()
	time.Sleep(2 * time.Second)
	fmt.Printf("race1:%v", f.Name)
}

func Race2() {
	f := &Fruit{Name: "Apple"}
	go func(fNew *Fruit) {
		fNew.Name = "Pear"
	}(f)
	time.Sleep(2 * time.Second)
	fmt.Printf("race2:exp:Apple,act:%v", f.Name)
}

func NoRace() {
	f := &Fruit{Name: "Apple"}
	go func(fNew Fruit) {
		fNew.Name = "Pear"
	}(*f)
	time.Sleep(2 * time.Second)
	fmt.Printf("race3:exp:Apple,act:%v", f.Name)
}
