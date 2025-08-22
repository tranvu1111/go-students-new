package command

import (
	"time"

	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/application/common"
)

type UpdateStudentCommand struct {
	IdempotencyKey	string
	StudentId		uuid.UUID
	DateOfBirth 	*time.Time 	 	
	Phone 			*string 
	Major 			*string 
	
}

type UpdateStudentCommandResult struct {
	Result *common.StudentResult

}
