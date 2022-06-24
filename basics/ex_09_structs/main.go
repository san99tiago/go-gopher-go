// Golang program to illustrate structs
// Santiago Garcia Arango

// Including main package
package main

import "fmt"

func main() {
	fmt.Println("\n********** EXAMPLE 1 (simple structs)**********")
	fmt.Println("\n--> [OPERATION] Create 'santi' and 'moni' of type 'Person'")
	var santi Person = Person{
		name:      "Santiago",
		last_name: "Garcia",
		age:       22,
	}
	var moni Person = Person{"Monica", "Hill", 23}
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "santi", santi, santi)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "santi.name", santi.name, santi.name)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "moni", moni, moni)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "moni.name", moni.name, moni.name)

	fmt.Println("\n--> [OPERATION 'update age of santi'] santi.age = 23")
	santi.age = 23
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "santi", santi, santi)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "santi.age", santi.age, santi.age)

	fmt.Println("\n********** EXAMPLE 2 (pointer of struct)**********")
	fmt.Println("\n--> [OPERATION] Create 'frailejon' of type '&Person'")
	frailejon := &Person{
		name:      "Ernesto",
		last_name: "Perez",
		age:       60,
	}
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "frailejon", frailejon, frailejon)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "frailejon.age", frailejon.age, frailejon.age)

	fmt.Println("\n--> [OPERATION 'update age of frailejon'] frailejon.age = 70")
	frailejon.age = 70
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "frailejon", frailejon, frailejon)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "frailejon.age", frailejon.age, frailejon.age)

	fmt.Println("\n********** EXAMPLE 3 (struct with pointer of other struct)**********")
	a_cool_student := Student{
		university: "EIA University",
		gpa:        4.73,
		person:     &santi,
	}
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "a_cool_student", a_cool_student, a_cool_student)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "a_cool_student.person", a_cool_student.person, a_cool_student.person)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "a_cool_student.person.name", a_cool_student.person.name, a_cool_student.person.name)
}

type Person struct {
	name      string
	last_name string
	age       uint32
}

type Student struct {
	university string
	gpa        float32
	person     *Person
}
