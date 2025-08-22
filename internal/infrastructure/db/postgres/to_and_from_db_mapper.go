package postgres

import "github.com/tranvu1111/go-students-new/internal/domain/entities"

func toDBStudent(validStudent *entities.ValidatedStudent) *DBStudent {
	return &DBStudent{
		StudentID: 		validStudent.StudentID,
		FirstName: 		validStudent.FirstName,
		LastName: 		validStudent.LastName,
		DateOfBirth: 	validStudent.DateOfBirth,
		Email: 			validStudent.Email,
		Phone: 			validStudent.Phone,
		Major: 			validStudent.Major,
		EnrollmentDate: validStudent.EnrollmentDate,
		CreatedAt: 		validStudent.CreatedAt,
		UpdatedAt: 		validStudent.UpdatedAt,
	}
}

func fromDBStudent(dbStudent *DBStudent) *entities.Student {
	var s = &entities.Student{
		StudentID: dbStudent.StudentID,
		FirstName: dbStudent.FirstName,
		LastName: dbStudent.LastName,
		DateOfBirth: dbStudent.DateOfBirth,
		Email: dbStudent.Email,
		Phone: dbStudent.Phone,
		Major: dbStudent.Major,
		EnrollmentDate: dbStudent.EnrollmentDate,
		CreatedAt: dbStudent.CreatedAt,
		UpdatedAt: dbStudent.UpdatedAt,
	}
	return s
}