package mapper

import (
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
	"github.com/tranvu1111/go-students-new/internal/application/common"
)

func NewStudentResultFromValidatedEntity(validatedStudent *entities.ValidatedStudent) *common.StudentResult {
	return NewStudentResultFromEntity(&validatedStudent.Student)
}

func NewStudentResultFromEntity(student *entities.Student) *common.StudentResult {
	if student == nil {
		return nil
	}

	return &common.StudentResult{
		StudentID: student.StudentID,
		FirstName: student.FirstName,
		LastName: student.LastName,
		DateOfBirth: student.DateOfBirth,
		Email: student.Email,
		Phone: student.Phone,
		Major: student.Major,
		CreatedAt: student.CreatedAt,
		UpdatedAt: student.UpdatedAt,
		EnrollmentDate: student.EnrollmentDate,
	}
}