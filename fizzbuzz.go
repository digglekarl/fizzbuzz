package main

import (
	"fmt"
	"strconv"
)

type fizzbuzzInterface interface {
	handle(*fizzbuzz)
	setNext(fizzbuzzInterface)
}

type divisibleBy3 struct {
	next fizzbuzzInterface
}

func (d *divisibleBy3) handle(f *fizzbuzz) {

	if f.num%3 == 0 {
		fmt.Println("Fizz")
		f.done = true
	}

	d.next.handle(f)

}

func (d *divisibleBy3) setNext(next fizzbuzzInterface) {
	d.next = next
}

type divisibleBy5 struct {
	next fizzbuzzInterface
}

func (d *divisibleBy5) handle(f *fizzbuzz) {
	if f.num%5 == 0 {
		fmt.Println("Buzz")
		f.done = true
	}

	d.next.handle(f)
}

func (d *divisibleBy5) setNext(next fizzbuzzInterface) {
	d.next = next
}

type divisibleBy3And5 struct {
	next fizzbuzzInterface
}

func (d *divisibleBy3And5) handle(f *fizzbuzz) {

	if f.num%3 == 0 && f.num%5 == 0 {
		fmt.Println("FizzBuzz")
		f.done = true
	}

	d.next.handle(f)
}

func (d *divisibleBy3And5) setNext(next fizzbuzzInterface) {
	d.next = next
}

type divisibleByNone struct {
	next fizzbuzzInterface
}

func (d *divisibleByNone) handle(f *fizzbuzz) {
	if f.done == false {
		fmt.Println(strconv.Itoa(f.num))
	}
}

func (d *divisibleByNone) setNext(next fizzbuzzInterface) {
	d.next = next
}

type fizzbuzz struct {
	num  int
	done bool
}

func main() {

	divisibleByNone := &divisibleByNone{}

	divisibleBy3And5 := &divisibleBy3And5{}
	divisibleBy3And5.setNext(divisibleByNone)

	divisibleBy5 := &divisibleBy5{}
	divisibleBy5.setNext(divisibleBy3And5)

	divisibleBy3 := &divisibleBy3{}
	divisibleBy3.setNext(divisibleBy5)

	for i := 1; i <= 100; i++ {
		fizzbuzz := &fizzbuzz{num: i, done: false}

		divisibleBy3.handle(fizzbuzz)
	}

}
