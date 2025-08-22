package entities



type ValidatedStudent struct {
	Student
	isValidated bool
}

func (vs *ValidatedStudent) IsValid() bool {
	return  vs.isValidated
}

func NewValidatedStudent(student *Student) (*ValidatedStudent , error) {
	if err := student.validate(); err != nil{
		return nil, err

	}
	return &ValidatedStudent{
		Student: *student,
		isValidated: true,
	}, nil
}
