package main

import (
	"fmt"
)

type Human struct {
	name   string
	age    int
	weight float32
}

type Student struct {
	Human
	speciality string
}

func main() {
	mark := Student{Human{"Mark", 18, 60}, "Computer Science"}
	fmt.Println("His name is", mark.name)
	fmt.Println("His age is", mark.age)
	fmt.Println("His weight is", mark.weight)
	fmt.Println("His Speciality is", mark.speciality)

	mark.speciality = "AI"
	fmt.Println("He changed his speciality to ", mark.speciality)
	mark.age += 40
	fmt.Println("Mark has becoome old and his age has become ", mark.age)

}
