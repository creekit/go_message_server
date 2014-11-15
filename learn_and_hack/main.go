package main

import "fmt"

type person struct {
	name string
	age  int
}

func Older(p1, p2 person) (person, int) {
	if p1.age < p2.age {
		return p2, p2.age - p1.age
	} else {
		return p1, p1.age - p2.age
	}
}

type student struct {
	person
	major string
}

func testStruct() {
	var p person
	p.name = "creekit"
	p.age = 1

	//p2 := person{"creek2", 2}

	//p3 := person{age: 3, name: "creek3"}

	p4 := new(person)
	p4.age = 4
	p4.name = "creekit"

	p5 := &person{name: "creek", age: 5}

	fmt.Printf("the person's name:%v, age:%v\n", p4.name, p4.age)
	fmt.Printf("the person's name:%v, age:%v\n", p5.name, p5.age)

	olderP, diff_age := Older(*p4, *p5)
	fmt.Printf("older person name:%s, diff:%v\n", olderP.name, diff_age)

}

func testAnounymousStruct() {
	s1 := student{person{name: "creekit", age: 1}, "english"}
	s1.age = 2
	s1.person = person{name: "creekit_again", age: 3}
	s1.person.age = 4

	fmt.Printf("student name:%v, age:%v, major:%v\n", s1.name, s1.age, s1.major)
}

func test() {
	fmt.Println("in test")

	//testStruct()
	testAnounymousStruct()
}

func main() {
	fmt.Println("hello, go")
	test()

}
