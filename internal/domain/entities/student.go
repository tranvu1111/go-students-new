package entities

import (
	"errors"
	"regexp"
	"time"

	"github.com/google/uuid"
)

type Student struct {
	StudentID 		uuid.UUID 
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

func NewStudent(first_name string, last_name string, date_of_birth *time.Time, email string,
	phone *string, major *string, enrollment_date time.Time ) *Student {
	return &Student	{	
		StudentID:			uuid.New() ,
		FirstName:			first_name ,
		LastName:			last_name ,
		DateOfBirth:		date_of_birth ,
		Email:				email 	,
		Phone:				phone ,
		Major:				major ,
		EnrollmentDate:		enrollment_date,
		CreatedAt: 			time.Now(),
		UpdatedAt: 			time.Now(),
	}	
}

func (s *Student) validate() error {	
	if s.FirstName == "" {
		return errors.New("Must have first name.")
	}

	if s.LastName == "" {
		return errors.New("Must have last name.")

	}

	if s.StudentID == uuid.Nil {
		return errors.New("Student ID can't be nil")
	}

	if s.Email == ""{
		return errors.New("Email can't be empty")

	}

	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	if !emailRegex.MatchString(s.Email) {
		return errors.New("Invalid email")
	}

	if s.EnrollmentDate.IsZero() {
		return errors.New("The enrollment date can't be zero")
	}

	if s.DateOfBirth != nil && s.DateOfBirth.After(time.Now()) {
		return errors.New("Invalid date of birth")
	}

	if s.Phone != nil && *s.Phone == "" {
		return errors.New("Phone cannot be an empty string if provided")
	}

	if s.Major != nil && *s.Major == "" {
		return errors.New("The major cannot be an empty string if provided")
	}

	if s.CreatedAt.IsZero() {
		return errors.New("CreatedAt is required and cannot be zero")
	}
	if s.UpdatedAt.IsZero() {
		return errors.New("UpdatedAt is required and cannot be zero")
	}

	return nil
 

} 

func (s *Student) UpdateNewFields(dob *time.Time, phone *string, major *string) error {
	s.DateOfBirth = dob
	s.Phone = phone
	s.Major = major
	s.UpdatedAt = time.Now()

	return s.validate()

}

