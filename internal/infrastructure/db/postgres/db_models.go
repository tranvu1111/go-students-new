package postgres

import (
	"time"
	"github.com/google/uuid"
)

type DBStudent struct {
	StudentID 		uuid.UUID 		`gorm:"primaryKey"`
	FirstName 		string 
	LastName 		string 
	DateOfBirth 	*time.Time 
	Email 			string 	
	Phone 			*string 
	Major 			*string 
	EnrollmentDate 	time.Time 
	CreatedAt 		time.Time
	UpdatedAt 		time.Time
}

type DBIdempotencyRecord struct {
	ID         uuid.UUID	`gorm:"primaryKey"`
	Key        string		`gorm:"uniqueIndex"`
	Request    string
	Response   string
	StatusCode int
	CreatedAt  time.Time
}