package common

import (
	"time"
	"github.com/google/uuid"
)

type StudentResult struct {
	StudentID 		uuid.UUID 
	FirstName 		string 
	LastName 		string 
	DateOfBirth 	*time.Time 
	Email 			string 	
	Phone 			*string 
	Major 			*string
	CreatedAt 		time.Time
	UpdatedAt 		time.Time 
	EnrollmentDate 	time.Time 
}