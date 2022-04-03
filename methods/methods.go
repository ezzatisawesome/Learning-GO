package main

import "fmt"

type Student struct {
	name   string
	grades []int
	age    int
}

func (s Student) getAge() int {
	return s.age
}

func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func main() {
	s1 := Student{"ezzat", []int{90, 50, 80, 100}, 35}
	fmt.Println(s1.getAge())
	s1.setAge(25)
	fmt.Println(s1.age)
	fmt.Println(s1.getAverageGrade())
}
