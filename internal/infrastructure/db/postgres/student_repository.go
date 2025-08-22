package postgres

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
	"github.com/tranvu1111/go-students-new/internal/domain/repositories"
	"gorm.io/gorm"
)

type GormStudentRepo struct {
	db *gorm.DB
}

func NewGormStudentRepo(db *gorm.DB) repositories.StudentRepository {
	return &GormStudentRepo{db:db}
}


func (repo *GormStudentRepo) Create(student *entities.ValidatedStudent) (*entities.Student,error) {
	dbStudent := toDBStudent(student)

	if err := repo.db.AutoMigrate(&DBStudent{}); err != nil {
		log.Fatalf("Fail to auto migrate Postgres schema: %v", err)
	}

	if err := repo.db.Create(dbStudent).Error; err != nil {
		return nil, err
	}

	return repo.FindById(dbStudent.StudentID)
}

func (repo *GormStudentRepo) FindById(id uuid.UUID) (*entities.Student, error) {
	var dbStudent DBStudent
	if err := repo.db.First(&dbStudent, id).Error; err != nil {
		return nil, err
	}

	// Map back to domain entity
	return fromDBStudent(&dbStudent), nil
}


func (repo *GormStudentRepo) FindAll() ([]*entities.Student , error) {
	var dbStudents []DBStudent
	if err := repo.db.Find(&dbStudents).Error;err != nil {
		return nil, err
	}

	students := make([]*entities.Student, len(dbStudents))

	for i, dbStudent := range dbStudents {
		students[i] = fromDBStudent(&dbStudent)
		fmt.Println(students[i])
	}
	return students,nil
}

func (repo *GormStudentRepo) Update(student *entities.ValidatedStudent) (*entities.Student, error) {
	dbStudent := *toDBStudent(student)

	// if err := repo.db.AutoMigrate(&DBStudent{}); err != nil {
	// 	log.Fatalf("Fail to auto migrate Postgres schema: %v", err)
	// }
	if err := repo.db.Model(&DBStudent{}).Where("student_id = ?", dbStudent.StudentID).Omit("student_id").Updates(dbStudent).Error; err != nil {
		return nil, err
	}

	return repo.FindById(dbStudent.StudentID)
		
}

func (repo *GormStudentRepo) Delete(id uuid.UUID) error {
	return repo.db.Delete(&DBStudent{}, id).Error
}


