package main

import "fmt"

type Course struct {
	Id    int
	Title string
	Price int
}

type Student struct {
	Name    string
	Courses []Course
}

var catalog = make(map[int]Course)
var students = make(map[string]Student)

// 1
func AddCourse(c Course) {
	catalog[c.Id] = c
}

// 2
func AddStudent(name string) {
	students[name] = Student{
		Name:    name,
		Courses: []Course{},
	}
}

// 3
func Enroll(name string, courseId int) {
	student, ok1 := students[name]
	course, ok2 := catalog[courseId]

	if !ok1 || !ok2 {
		return
	}

	student.Courses = append(student.Courses, course)
	students[name] = student
}

// 4
func TotalCost(name string) int {
	student, ok := students[name]
	if !ok {
		return 0
	}

	total := 0
	for _, course := range student.Courses {
		total += course.Price
	}

	return total
}

// 5
func PrintStudent(name string) {
	student, ok := students[name]
	if !ok {
		return
	}

	fmt.Println("Student:", student.Name)
	fmt.Println("Courses:")

	for _, course := range student.Courses {
		fmt.Println(course.Title, "-", course.Price)
	}

	fmt.Println("Total:", TotalCost(name))
	fmt.Println()
}

func main() {

	// 1
	AddCourse(Course{1, "Go Basics", 20000})
	AddCourse(Course{2, "SQL", 15000})
	AddCourse(Course{3, "Web", 18000})

	// 2
	AddStudent("Ali")
	AddStudent("Dana")

	// 3
	Enroll("Ali", 1)
	Enroll("Ali", 2)

	Enroll("Dana", 2)
	Enroll("Dana", 3)

	// 4
	PrintStudent("Ali")
	PrintStudent("Dana")
}
