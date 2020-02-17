package main

var students = []*Student{}

type Student struct {
	Id    string
	Name  string
	Grade int32
}

func GetStudents() []*Student {
	return students
}

func SelectStudent(id string) *Student {
	for _, each := range students {
		if each.Id == id {
			return each
		}
	}

	return nil
}
func init() {
	students = append(students, &Student{Id: "s001", Name: "foo", Grade: 2})
	students = append(students, &Student{Id: "s002", Name: "bar", Grade: 2})
	students = append(students, &Student{Id: "s003", Name: "foobar", Grade: 3})
}
