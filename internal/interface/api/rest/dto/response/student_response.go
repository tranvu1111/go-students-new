package response

import (
	"time"

	// "github.com/google/uuid"
)

type StudentResponse struct {
	StudentID      	string
	FirstName      	string
	LastName       	string
	DateOfBirth 	*time.Time 	
	Email          	string
	Phone 			*string 
	Major 			*string 
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
	EnrollmentDate 	time.Time 	
}

type StudentResponseList struct {
	Students []*StudentResponse		`json:"Students"`
}