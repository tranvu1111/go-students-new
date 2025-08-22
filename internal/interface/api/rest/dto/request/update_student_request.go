package request

import (
	"time"

	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/application/command"
)

// JsonTime is a custom type to handle JSON unmarshaling of time strings.


type UpdateStudentResquest struct {
	IdempotencyKey string		`json:"IdempotencyKey"`
	StudentId      uuid.UUID	`json:"StudentId"`
	DateOfBirth    *JsonTime	`json:"DateOfBirth"`
	Phone          *string		`json:"Phone"`
	Major          *string		`json:"Major"`
}

func (ur *UpdateStudentResquest) ToUpdateStudentCommand() (*command.UpdateStudentCommand , error ) {
	var dateOfBirth *time.Time
	if ur.DateOfBirth != nil {
		// Dereference the JsonTime pointer, convert it to time.Time,
		// and then get a new pointer to that value.
		convertedDate := time.Time(*ur.DateOfBirth)
		dateOfBirth = &convertedDate
	}

	return &command.UpdateStudentCommand{
		IdempotencyKey:	ur.IdempotencyKey,
		StudentId:		ur.StudentId,
		DateOfBirth: 	dateOfBirth, 	 	
		Phone: 			ur.Phone,
		Major: 			ur.Major,
	},nil


}

