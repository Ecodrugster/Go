package main

import "fmt"

type User struct {
	name string
	age  int
}

func (u User) GetName() string {
	return u.name
}

func (u User) GetAge() int {
	return u.age
}

func (u *User) SetName(name string) {
	if name == "" {
		fmt.Println("Ошибка: имя не может быть пустым")
		return
	}
	u.name = name
}

func (u *User) SetAge(age int) {
	if age < 0 {
		fmt.Println("Ошибка: возраст не может быть отрицательным")
		return
	}
	u.age = age
}

type Participant interface {
	Info()
	Act()
}

type Student struct {
	User
	Score int
}

func (s *Student) Info() {
	fmt.Println("Student:", s.GetName(), "Score:", s.Score)
}

func (s *Student) Act() {
	fmt.Println(s.GetName(), "учится")
}

type Teacher struct {
	User
}

func (t *Teacher) Info() {
	fmt.Println("Teacher:", t.GetName())
}

func (t *Teacher) Act() {
	fmt.Println(t.GetName(), "преподает")
}

func (t *Teacher) GradeStudent(student *Student, points int, maxScore int) {
	student.Score += points
	if student.Score > maxScore {
		student.Score = maxScore
	}
}

type Course struct {
	Title       string
	MaxScore    int
	MaxStudents int
	Students    []*Student
	Teacher     *Teacher
}

func (c *Course) AddStudent(s *Student) {
	if len(c.Students) >= c.MaxStudents {
		fmt.Println("Лимит студентов достигнут")
		return
	}
	c.Students = append(c.Students, s)
}

func (c Course) Info() {
	fmt.Println("Course:", c.Title)
	fmt.Println("MaxScore:", c.MaxScore)
	fmt.Println("Students count:", len(c.Students))
	fmt.Println()
}

func AverageScore(students []*Student) float64 {
	if len(students) == 0 {
		return 0
	}

	sum := 0
	for _, s := range students {
		sum += s.Score
	}

	return float64(sum) / float64(len(students))
}

func BestStudent(students []*Student) *Student {
	if len(students) == 0 {
		return nil
	}

	best := students[0]

	for _, s := range students {
		if s.Score > best.Score {
			best = s
		}
	}

	return best
}

func main() {

	teacher := &Teacher{}
	teacher.SetName("Ivan")
	teacher.SetAge(35)

	s1 := &Student{}
	s1.SetName("Ali")
	s1.SetAge(20)

	s2 := &Student{}
	s2.SetName("Dana")
	s2.SetAge(21)

	s3 := &Student{}
	s3.SetName("Amina")
	s3.SetAge(19)

	course := Course{
		Title:       "Go Course",
		MaxScore:    100,
		MaxStudents: 3,
		Teacher:     teacher,
	}

	course.AddStudent(s1)
	course.AddStudent(s2)
	course.AddStudent(s3)

	course.Info()

	var participants []Participant = []Participant{s1, s2, s3, teacher}

	for _, p := range participants {
		p.Info()
		p.Act()
	}

	fmt.Println()

	teacher.GradeStudent(s1, 90, course.MaxScore)
	teacher.GradeStudent(s2, 70, course.MaxScore)
	teacher.GradeStudent(s3, 95, course.MaxScore)

	avg := AverageScore(course.Students)
	fmt.Println("Average score:", avg)

	best := BestStudent(course.Students)
	if best != nil {
		fmt.Println("Best student:", best.GetName(), "Score:", best.Score)
	}
}
