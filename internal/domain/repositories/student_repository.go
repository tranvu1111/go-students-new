package repositories

import (
	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
)

type StudentRepository interface {

	Create(student *entities.ValidatedStudent) (*entities.Student, error)
	FindById(id uuid.UUID) (*entities.Student, error)
	FindAll() ([]*entities.Student, error)
	Update(student *entities.ValidatedStudent) (*entities.Student, error)
	Delete(id uuid.UUID) error

}