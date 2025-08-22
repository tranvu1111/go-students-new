package entities

import (
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	
)

func TestNewStudent(t *testing.T) {
	date_of_birth := time.Date(2003, time.March,11,02,0,0,0,time.UTC)
	enrollment_date := time.Now()
	phone := "0947531799"
	major := "Computer Science"


	s := NewStudent(
		"tran",
		"vu",
		&date_of_birth,
		"tranvu123@gmail.com",
		&phone,
		&major,
		enrollment_date,
	)

	if s.StudentID == uuid.Nil {
		t.Errorf("Expected no-nil studentID, got nil value")
	}

	if s.FirstName == "" {
		t.Errorf("Expected non-empty first name, got empty")
	}

	if s.FirstName != "tran" {
		t.Errorf("Expected 'tran' first name, got %s", s.FirstName	)
	}

	if s.LastName == "" {
		t.Errorf("Expected non-empty last name, got empty")
	}

	if s.LastName != "vu" {
		t.Errorf("Expected 'vu' last name, got %s", s.LastName)
	}

	if s.DateOfBirth != nil && s.DateOfBirth != &date_of_birth {
		t.Errorf("Expect %v but got %v", date_of_birth, s.DateOfBirth)
	}

	if s.Email != "tranvu123@gmail.com" {
		t.Errorf("Expected Email 'tranvu123@gmail.com' but got %v", s.Email)
	}

	if s.Phone ==  nil ||  s.Phone != &phone {
		t.Errorf("Expected Phone '0947531799' but got %v", *s.Phone)
	}

	if s.Major == nil || s.Major != &major {
		t.Errorf("Expected major 'Computer Science' but got %v", *s.Major)
	}

	if !s.EnrollmentDate.Equal(enrollment_date){
		t.Errorf("Expected enrollment_date is %v but got %v ", enrollment_date, s.EnrollmentDate)
	}

	student := NewStudent("Jane", "Smith",nil,"jane@example.com", nil, nil, enrollment_date)
	if student.DateOfBirth != nil {
		t.Errorf("Expected DateOfBirth nil, got '%v'", student.DateOfBirth)
	}
	if student.Phone != nil {
		t.Errorf("Expected Phone nil, got '%v'", student.Phone)
	}
	if student.Major != nil {
		t.Errorf("Expected Major nil, got '%v'", student.Major)
	}
	
}

func TestStudent_Validate(t *testing.T) {
	validPhone := "0442312300"
	validMajor := "ECM"
	emptyPhone := ""


	future_dob := time.Now().AddDate(1, 0, 0)

	validStudent := &Student{
		StudentID:      uuid.New(),
		FirstName:      "John",
		LastName:       "Doe",
		DateOfBirth:    &time.Time{}, // Use a valid time
		Email:          "john.doe@example.com",
		Phone:          &validPhone,
		Major:          &validMajor,
		EnrollmentDate: time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	testCases := []struct {
		name_case	string
		student 	*Student
		expectedErr error
	}{
		{
			name_case:          "Valid Student",
			student:       validStudent,
			expectedErr: nil,
		},
		{
			name_case: "Missing First Name",
			student: &Student{
				StudentID:      uuid.New(),
				LastName:       "Doe",
				Email:          "john.doe@example.com",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: errors.New("Must have first name."),
		},
		{
			name_case: "Missing Last Name",
			student: &Student{
				StudentID:      uuid.New(),
				FirstName:       "Doe",
				Email:          "john.doe@example.com",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: errors.New("Must have last name."),
		},
		{
			name_case: "Missing ID",
			student: &Student{
				
				FirstName:       "John",
				LastName:    	"Doe",	
				Email:          "john.doe@example.com",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: errors.New("Student ID can't be nil"),
		},
		{
			name_case: "Missing Email",
			student: &Student{
				StudentID:      uuid.New(),
				FirstName:       "John",
				LastName:    	"Doe",	
				Email:          "",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: errors.New("Email can't be empty"),
		},
		{
			name_case: "Invalid Email",
			student: &Student{
				StudentID:      uuid.New(),
				FirstName:       "John",
				LastName:    	"Doe",	
				Email:          "sadfawefawfe111",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: errors.New("Invalid email"),
		},
		{
			name_case: "Invalid Email to long at the end",
			student: &Student{
				StudentID:      uuid.New(),
				FirstName:       "John",
				LastName:    	"Doe",	
				Email:          "sadfawe@gamil.cavsm",
				EnrollmentDate: time.Now(),
				CreatedAt:      time.Now(),
				UpdatedAt:      time.Now(),
			},
			expectedErr: errors.New("Invalid email"),
		},
		{
			name_case: "Zero date of birth",
			student:&Student{
					StudentID:     	 uuid.New(),
					FirstName:       "John",
					LastName:    	"Doe",	
					Email:          "john.doe@example.com",
					DateOfBirth: 	&future_dob,
					EnrollmentDate: time.Now(),
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				},
				
			
			expectedErr: errors.New("Invalid date of birth"),
		},
		{
			name_case: "Empty phone number",
			student:&Student{
					StudentID:     	 uuid.New(),
					FirstName:       "John",
					LastName:    	"Doe",	
					Email:          "john.doe@example.com",
					Phone:			&emptyPhone,
					EnrollmentDate: time.Now(),
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Now(),
				},
			expectedErr: errors.New("Phone cannot be an empty string if provided"),
		},
		{
			name_case: "Create At is zero",
			student:&Student{
					StudentID:     	 uuid.New(),
					FirstName:       "John",
					LastName:    	"Doe",	
					Email:          "john.doe@example.com",					
					EnrollmentDate: time.Now(),
					CreatedAt:      time.Time{},
					UpdatedAt:      time.Now(),
				},
			expectedErr: errors.New("CreatedAt is required and cannot be zero"),
		},
		{
			name_case: "UpdatedAt is zero",
			student:&Student{
					StudentID:     	 uuid.New(),
					FirstName:       "John",
					LastName:    	"Doe",	
					Email:          "john.doe@example.com",					
					EnrollmentDate: time.Now(),
					CreatedAt:      time.Now(),
					UpdatedAt:      time.Time{},
				},
			expectedErr: errors.New("UpdatedAt is required and cannot be zero"),
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.name_case, func(t *testing.T){
			err := tc.student.validate()

			if tc.expectedErr != nil {
				if err == nil {
					t.Errorf("Expected non-nil but got %s", err.Error())
				}

				if err.Error() != tc.expectedErr.Error() {
					t.Errorf("unexpected error message: got %v, want %v", err, tc.expectedErr)
				}
			} else {
				if err != nil {
					t.Errorf("Expected no error but got %v", err)
				}
			}
		})
	} 


}