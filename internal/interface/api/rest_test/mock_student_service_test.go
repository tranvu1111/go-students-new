package rest_test

import (
	// "time"

	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/tranvu1111/go-students-new/internal/application/command"
	"github.com/tranvu1111/go-students-new/internal/application/mapper"
	"github.com/tranvu1111/go-students-new/internal/application/query"
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
)

type MockStudentService struct {
	mock.Mock
}

func(m *MockStudentService) CreateStudent(studentCommand *command.CreateStudentCommand) (*command.CreateStudentCommandResult,error) {
	args := m.Called(studentCommand)
	var result command.CreateStudentCommandResult

	// var now = time.Now()

	var student = entities.NewStudent(	
		studentCommand.FirstName,        			
		studentCommand.LastName, 		 			
		studentCommand.DateOfBirth, 	 
		studentCommand.Email, 			 	
		studentCommand.Phone, 			 
		studentCommand.Major, 			 
		studentCommand.EnrollmentDate, 		
	)

	var validatedStudent, err = entities.NewValidatedStudent(student)
	if err != nil {
		return nil, err
	}

	
	result.Result =  mapper.NewStudentResultFromValidatedEntity(validatedStudent)

	return &result , args.Error(1)

}

func (m *MockStudentService) FindAllStudent()(*query.StudentQueryListResult, error){
	args := m.Called()

	studentQueryListResult := &query.StudentQueryListResult{}

	for _, s := range args.Get(0).([]*entities.Student){
		studentQueryListResult.Result = append(studentQueryListResult.Result, mapper.NewStudentResultFromEntity(s))
	}

	return studentQueryListResult, args.Error(1)

}

func (m *MockStudentService) FindStudentById(id uuid.UUID)(*query.StudentQueryResult, error){
	args := m.Called(id)

	studentQueryResult := &query.StudentQueryResult{
		Result: mapper.NewStudentResultFromEntity(args.Get(0).(*entities.Student)),
	}

	return studentQueryResult, args.Error(1)
}

func(m *MockStudentService) UpdateStudent(updateCommand *command.UpdateStudentCommand) (*command.UpdateStudentCommandResult, error) {
	args := m.Called(updateCommand)

	var result command.UpdateStudentCommandResult
	studentID , _ := uuid.Parse("6f69799c-1eb2-4266-b28c-9762a4d02129")
	
	enrollment_date := time.Date(2023, 3, 11, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	student := entities.Student{
		StudentID: studentID,
		FirstName: "tran",
		LastName: "vu",
		DateOfBirth: updateCommand.DateOfBirth,
		Email: "tranvu333@gmail.com",
		Phone: updateCommand.Phone,
		Major: updateCommand.Major,
		EnrollmentDate: enrollment_date,
		CreatedAt: now,
		UpdatedAt: now,
	}

	validStudent , err := entities.NewValidatedStudent(&student)
	if err != nil {
		return nil, err
	}

	result.Result = mapper.NewStudentResultFromValidatedEntity(validStudent)
	return &result , args.Error(1)
}

func(m *MockStudentService) DeleteStudent(id uuid.UUID)(error) {
	return nil
}