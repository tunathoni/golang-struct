package morestring

// Student is a struct
type Student struct {
	Name    string
	Age     int
	Address string
	StudentClass
}

var allStudents = []Student{}

// SetStudent is a function to set Student Detail
func SetStudent(Name string, Age int, Address string, ClassName string) {
	newStudent := Student{}
	newStudent.Name = Name
	newStudent.Age = Age
	newStudent.Address = Address
	newStudent.StudentClass.Name = ClassName
	allStudents = append(allStudents, newStudent)
}

// GetStudentDetail is a fucntion to get Student Detail
func GetStudentDetail() []Student {
	return allStudents
}

// GetAgeAcerage to get average of all students
func GetAgeAcerage() int {
	total := 0
	for _, item := range allStudents {
		total += item.Age
	}

	average := total / len(allStudents)

	return average
}
