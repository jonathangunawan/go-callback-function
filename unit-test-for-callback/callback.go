package main

import "fmt"

func main() {
	num1 := 2
	num2 := 5

	dep := Usecase{
		m: Main{},
		s: Sub{},
	}
	dep.SomeService(num1, num2)
}

type Usecase struct {
	m MainItf
	s SubItf
}

type UsecaseItf interface {
	SomeService(data1, data2 int) error
}

func (u Usecase) SomeService(data1, data2 int) error {
	result, err := u.m.Process(data1, data2, func(num1 int, num2 int) (int, error) {
		res, err := u.s.AnotherProcess(num1, num2)
		if err != nil {
			return 0, err
		}

		return res, nil
	})

	if err != nil {
		return fmt.Errorf("wrong")
	}

	fmt.Println("Number:", result)
	return nil
}

type Main struct{}

type MainItf interface {
	Process(num1 int, num2 int, callback func(int, int) (int, error)) (int, error)
}

func (m Main) Process(num1 int, num2 int, callback func(int, int) (int, error)) (int, error) {
	res, err := callback(num1, num2)
	if err != nil {
		return 0, err
	}

	return res, nil
}

type Sub struct{}

type SubItf interface {
	AnotherProcess(num1, num2 int) (int, error)
}

// only process if the input less than 10
func (s Sub) AnotherProcess(num1, num2 int) (int, error) {
	if num1 >= 10 || num2 >= 10 {
		return 0, fmt.Errorf("more than 10")
	}

	return num1 + num2, nil
}
