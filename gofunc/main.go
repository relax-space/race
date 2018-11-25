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
	//Race1()
	//2.has race
	//Race2()
	//3.no race
	//NoRace()
	//4.has race
	//Race3()
	//5.NoRace2
	NoRace2()
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

func NoRace() (newFruit *Fruit) {
	f := &Fruit{Name: "Apple"}
	go func(fNew Fruit) {
		fNew.Name = "Pear"
		newFruit = &Fruit{Name: "Pear"}
	}(*f)
	time.Sleep(2 * time.Second)
	fmt.Printf("no race:exp:Apple,act:%v", f.Name)
	return
}

func Race3() (newFruit Fruit) {
	go func() {
		newFruit = Fruit{Name: "Pear"}
	}()
	//time.Sleep(2 * time.Second)
	fmt.Printf("race3:exp:Pear,act:%v", newFruit.Name)
	return
}

func NoRace2() (newFruit *Fruit) {
	fChan := make(chan *Fruit)
	go func() {
		newFruit = &Fruit{Name: "Pear"}
		fChan <- newFruit
	}()
	newFruit = <-fChan
	fmt.Printf("race3:exp:Pear,act:%v", newFruit.Name)
	return
}
