package interfaces

import (
	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/application/command"
	"github.com/tranvu1111/go-students-new/internal/application/query"
)

type StudentService interface {
	CreateStudent(studentCommand *command.CreateStudentCommand)(*command.CreateStudentCommandResult, error)
	FindAllStudent()(*query.StudentQueryListResult, error)
	FindStudentById(id uuid.UUID)(*query.StudentQueryResult, error)
	UpdateStudent(updateCommand *command.UpdateStudentCommand)(*command.UpdateStudentCommandResult, error)
	DeleteStudent(id uuid.UUID)(error)
}