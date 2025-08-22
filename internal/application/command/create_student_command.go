package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/application/common"
)

type CreateStudentCommand struct {
	IdempotencyKey	string
	StudentId		uuid.UUID
	FirstName       string 			
	LastName 		string 			
	DateOfBirth 	*time.Time 
	Email 			string 	
	Phone 			*string 
	Major 			*string 
	EnrollmentDate 	time.Time 
}


type CreateStudentCommandResult struct {
	Result *common.StudentResult
}

