package main

import "fmt"

type Stringer interface {
	String() string
}

type google struct {
	workingYear int
	baseSalary int
}

type mercari struct {
	workingYear int
	baseSalary int
}

type apple struct {
	workingYear int
	baseSalary int
}

func (g google) String() string {
	return "Hello Google"
}

func (m mercari) String() string {
	return "Hello Mercari"
}

func (m apple) String() string{
	return "Hello Apple"
}

func main() {
	ueno := google{
		workingYear: 5,
		baseSalary: 100,
	}

	tomoya := mercari{
		workingYear: 4,
		baseSalary: 200,
	}

	yoko := apple{
		workingYear: 5,
		baseSalary: 100,
	}

	workers := []Stringer{ueno, tomoya, yoko}
	fmt.Println(workers)
}





