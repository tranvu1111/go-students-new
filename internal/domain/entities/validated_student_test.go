package entities

import (
	"testing"
	"time"
	"github.com/google/uuid"
)


func TestNewValidatedStudent(t *testing.T){
	validStudent := &Student{
		StudentID:      uuid.New(),
		FirstName:      "John",
		LastName:       "Doe",
		Email:          "john.doe@example.com",
		EnrollmentDate: time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	validatedStudent , err  := NewValidatedStudent(validStudent)

	if err != nil {
		t.Errorf("Expected a valid student but got, err : %s" , err.Error())
	}

	if !validatedStudent.IsValid() {
		t.Errorf("Expected a valid student but got an invalidated student")
	}

	invalidStudent := &Student{
		StudentID:      uuid.New(),
		FirstName:      "John",
		LastName:       "",
		Email:          "john.doe@example.com",
		EnrollmentDate: time.Now(),
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	validatedStudent, err = NewValidatedStudent(invalidStudent)
	if err == nil {
		t.Errorf("Expected a invalid student in return, but got no err")
	}

	if validatedStudent != nil {
		t.Errorf("Expected cannot create a valid student but still created")
	}

}