package request

import (
	"time"

	"github.com/tranvu1111/go-students-new/internal/application/command"
)

// JsonTime is a custom type to handle JSON unmarshaling of time strings.
type JsonTime time.Time

func (jt *JsonTime) UnmarshalJSON(b []byte) error {
	// The JSON value will be a string, so we need to trim the quotes.
	s := string(b)
	s = s[1 : len(s)-1]

	// Use time.Parse with the expected format (e.g., "YYYY-MM-DD").
	t, err := time.Parse("2006-01-02", s)
	if err != nil {
		return err
	}

	// Assign the parsed time to the JsonTime pointer.
	*jt = JsonTime(t)
	return nil
}

type CreateStudentRequest struct {
	IdempotencyKey string    `json:"IdempotencyKey"`
	FirstName      string    `json:"FirstName"`
	LastName       string    `json:"LastName"`
	DateOfBirth    *JsonTime `json:"DateOfBirth,omitempty"`
	Email          string    `json:"Email"`
	Phone          *string   `json:"Phone,omitempty"`
	Major          *string   `json:"Major,omitempty"`
	EnrollmentDate JsonTime  `json:"EnrollmentDate"`
}

func (req *CreateStudentRequest) ToCreateStudentCommand() (*command.CreateStudentCommand, error) {

	// Correctly handle the DateOfBirth pointer.
	var dateOfBirth *time.Time
	if req.DateOfBirth != nil {
		// Dereference the JsonTime pointer, convert it to time.Time,
		// and then get a new pointer to that value.
		convertedDate := time.Time(*req.DateOfBirth)
		dateOfBirth = &convertedDate
	}

	// Correctly convert the EnrollmentDate.
	enrollmentDate := time.Time(req.EnrollmentDate)
	return &command.CreateStudentCommand{
		IdempotencyKey: req.IdempotencyKey,
		FirstName:      req.FirstName,
		LastName:       req.LastName,
		DateOfBirth:    dateOfBirth,
		Email:          req.Email,
		Phone:          req.Phone,
		Major:          req.Major,
		EnrollmentDate: enrollmentDate,
	}, nil
}
