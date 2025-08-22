package services

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/application/command"
	"github.com/tranvu1111/go-students-new/internal/application/mapper"
	"github.com/tranvu1111/go-students-new/internal/application/query"

	// "github.com/tranvu1111/go-students-new/internal/application/common"
	"github.com/tranvu1111/go-students-new/internal/application/interfaces"
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
	"github.com/tranvu1111/go-students-new/internal/domain/repositories"
)


type StudentService struct {
	repo				repositories.StudentRepository
	idempotencyRepo 	repositories.IdempotencyRepository
}

func NewStudentService(	sr repositories.StudentRepository,	ir repositories.IdempotencyRepository) interfaces.StudentService  {
	return  &StudentService{
		repo: sr,
		idempotencyRepo: ir,
	}
}

func (s *StudentService) CreateStudent(studentCommand *command.CreateStudentCommand)(*command.CreateStudentCommandResult, error){
	ctx := context.Background()

	if studentCommand.IdempotencyKey != "" {
		existingRecord , err := s.idempotencyRepo.FindByKey(ctx, studentCommand.IdempotencyKey)
		if err != nil {
			return nil, err
		}

		if existingRecord != nil {
			var result command.CreateStudentCommandResult
			if err := json.Unmarshal([]byte(existingRecord.Response) , &result) ;err != nil {
				return nil, err
			}
			return &result, nil
		}
	}

	var idempotencyRecord *entities.IdempotencyRecord		
	if studentCommand.IdempotencyKey != "" {
		requestJSON, _ := json.Marshal(studentCommand)
		idempotencyRecord = entities.NewIdempotencyRecord(studentCommand.IdempotencyKey,string(requestJSON))
	}

	var newStudent = entities.NewStudent(
		studentCommand.FirstName,
		studentCommand.LastName,
		studentCommand.DateOfBirth,
		studentCommand.Email,
		studentCommand.Phone,
		studentCommand.Major,
		studentCommand.EnrollmentDate,
	)

	validatedStudent, err := entities.NewValidatedStudent(newStudent)
	if err != nil {
		return nil, err
	}

	_, err = s.repo.Create(validatedStudent)
	if err != nil {
		return nil, err
	}


	result := command.CreateStudentCommandResult{
		Result: mapper.NewStudentResultFromValidatedEntity(validatedStudent),
	}

	if idempotencyRecord != nil {
		responseJSON, _ := json.Marshal(result)
		idempotencyRecord.SetResponse(string(responseJSON), 200)
		_, err = s.idempotencyRepo.Create(ctx, idempotencyRecord)
		if err != nil {
			// Log error but don't fail the operation
			// In production, you might want to handle this differently
		}
	}

	return &result, nil
}

func (s *StudentService) FindAllStudent() (*query.StudentQueryListResult, error) {
	storedStudents ,err := s.repo.FindAll()
	if err != nil {
		return nil, err
	}

	var queryResult query.StudentQueryListResult
	for _, student := range storedStudents {
		queryResult.Result = append(queryResult.Result, mapper.NewStudentResultFromEntity(student))
	}

	return &queryResult, nil
}

func(s *StudentService) FindStudentById(id uuid.UUID)(*query.StudentQueryResult, error){
	student , err := s.repo.FindById(id)
	if err != nil {
		return nil, err
	}

	var queryResult query.StudentQueryResult
	queryResult.Result = mapper.NewStudentResultFromEntity(student)
	return &queryResult, nil
}

func(s *StudentService) UpdateStudent(updateCommand *command.UpdateStudentCommand)(*command.UpdateStudentCommandResult, error){
	ctx := context.Background()

	if updateCommand.IdempotencyKey != "" {
		existingRecord , err := s.idempotencyRepo.FindByKey(ctx, updateCommand.IdempotencyKey)
		if err != nil {
			return nil, err
		}

		if existingRecord != nil {
			var result command.UpdateStudentCommandResult
			if err := json.Unmarshal([]byte(existingRecord.Response) , &result) ;err != nil {
				return nil, err
			}
			return &result, nil
		}
	}

	var idempotencyRecord *entities.IdempotencyRecord		
	if updateCommand.IdempotencyKey != "" {
		requestJSON, _ := json.Marshal(updateCommand)
		idempotencyRecord = entities.NewIdempotencyRecord(updateCommand.IdempotencyKey,string(requestJSON))
	}

	storedStudent , err := s.repo.FindById(updateCommand.StudentId)
	if err != nil {
		return nil, err
	}

	if storedStudent == nil {
		return nil, errors.New("not found this student in order to update")
	}

	validUpdateStudent, err := entities.NewValidatedStudent(storedStudent)
	if err != nil {
		return nil, errors.New("this error came from find by ID")
	}

	if err := storedStudent.UpdateNewFields(updateCommand.DateOfBirth , updateCommand.Phone , updateCommand.Major) ; err != nil {
		return nil, err
	}

	validUpdateStudent, err = entities.NewValidatedStudent(storedStudent)
	if err != nil {
		return nil, errors.New("this error came from update fields in storedStudent")
	}
	_, err = s.repo.Update(validUpdateStudent)
	if err != nil {
		return nil, err
	}

	result := command.UpdateStudentCommandResult{
		Result: mapper.NewStudentResultFromValidatedEntity(validUpdateStudent),
	}

	if idempotencyRecord != nil {
		responseJSON, _ := json.Marshal(result)
		idempotencyRecord.SetResponse(string(responseJSON), 200)
		_, err = s.idempotencyRepo.Create(ctx, idempotencyRecord)
		if err != nil {
			// Log error but don't fail the operation
			// In production, you might want to handle this differently
		}
	}

	return &result, nil
}

func(s *StudentService)DeleteStudent(id uuid.UUID)(error) {
	return s.repo.Delete(id)
}

