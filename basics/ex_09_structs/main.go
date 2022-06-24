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

	fmt.Println("\n********** EXAMPLE 3 (struct with pointer of other struct and methods)**********")
	var a_cool_student Student = Student{
		university: "EIA University",
		grades:     []float32{5.0, 4.5},
		person:     &santi,
	}
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "a_cool_student", a_cool_student, a_cool_student)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "a_cool_student.person", a_cool_student.person, a_cool_student.person)
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "a_cool_student.person.name", a_cool_student.person.name, a_cool_student.person.name)

	fmt.Println("\n--> [OPERATION] student_gpa = a_cool_student.get_gpa()")
	student_gpa := a_cool_student.get_gpa()
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "student_gpa", student_gpa, student_gpa)

	fmt.Println("\n--> [OPERATION] a_cool_student.add_grade(4.8)")
	a_cool_student.add_grade(4.8)
	fmt.Println("\n--> [OPERATION] student_gpa_updated := a_cool_student.get_gpa()")
	student_gpa_updated := a_cool_student.get_gpa()
	fmt.Printf("\t[%s] has type of: [%T] and its value is: [%v]\n", "student_gpa_updated", student_gpa_updated, student_gpa_updated)
}

// PERSON STRUCT
type Person struct {
	name      string
	last_name string
	age       uint32
}

// STUDENT STRUCT THAT USES THE PERSON'S STRUCT
type Student struct {
	university string
	grades     []float32
	person     *Person
}

// METHODS FOR STUDENT STRUCT (BE CAREFUL TO USE POINTERS FOR 'SETTERS')
func (s *Student) add_grade(grade float32) {
	s.grades = append(s.grades, grade)
}

func (s Student) get_gpa() float32 {
	var sum float32 = 0
	for i := 0; i < len(s.grades); i++ {
		sum += (s.grades[i])
	}
	avg := (float32(sum)) / (float32(len(s.grades)))
	return avg
}
