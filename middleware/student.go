package main

import "errors"

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) (*Student, error) {
	for _, student := range students {
		if student.Id == id {
			return student, nil
		}
	}
	return nil, errors.New("id not found")
}

func init() {
	students = append(students, &Student{Id: "s001", Name: "John", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "Ethan", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "Wick", Grade: 3})
}
