package main

import (
	"fmt"
	"golang/hello/morestring"
)

func main() {
	a := 1
	b := &a
	a = 2
	total := a + *b
	fmt.Println("hello world = ", total) // hello world =  4
	morestr := morestring.ReverseRunes("test")
	fmt.Println("hello ", morestr)
	morestring.SetStudent("Thoni", 25, "Malang", "A")
	morestring.SetStudent("Zahra", 23, "Malang", "B")
	morestring.SetStudent("Eko", 25, "Malang", "C")

	studenDetail := morestring.GetStudentDetail()
	for _, item := range studenDetail {
		fmt.Println("Name = ", item.Name)
		fmt.Println("Age = ", item.Age)
		fmt.Println("Address = ", item.Address)
		fmt.Println("Class = ", item.StudentClass.Name)
		fmt.Println("========================")
	}

	fmt.Println("Age Average = ", morestring.GetAgeAcerage())
}
