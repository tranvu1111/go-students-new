// postgres/postgres_test.go

package db_test

import (
	"errors"
	"fmt"
	"testing"
	"time"

	"github.com/glebarez/sqlite"
	"github.com/google/uuid"
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
	"github.com/tranvu1111/go-students-new/internal/infrastructure/db/postgres"
	"gorm.io/gorm"
)

// setupTestDB initializes an in-memory SQLite database for testing.
// It auto-migrates the DBStudent model.
func setupTestDB(t *testing.T) (*postgres.GormStudentRepo, *gorm.DB) {
	// Open a connection to an in-memory SQLite database.
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to connect to in-memory database: %v", err)
	}

	// Auto-migrate the database schema for the DBStudent model.
	if err := db.AutoMigrate(&postgres.DBStudent{}); err != nil {
		t.Fatalf("Failed to auto-migrate schema: %v", err)
	}

	// Create and return a new GormStudentRepo instance.
	repo := postgres.NewGormStudentRepo(db).(*postgres.GormStudentRepo)
	return repo, db
}

// TestGormStudentRepo_Create tests the Create method for successful student creation.
func TestGormStudentRepo_Create(t *testing.T) {
	// Setup the test database and get the repository instance.
	repo, db := setupTestDB(t)

	// Create a sample ValidatedStudent entity for the test.
	testUUID := uuid.New()
	now := time.Now()
	student := &entities.Student{
		StudentID:   testUUID,
		FirstName:   "John",
		LastName:    "Doe",
		DateOfBirth: &now,
		Email:       "john.doe@example.com",
		Major:       nil, // Optional fields can be nil
		EnrollmentDate: now,
		CreatedAt: now,
		UpdatedAt: now,
	}

	testStudent, err := entities.NewValidatedStudent(student)
	if err != nil {
		t.Errorf("Fail to create a valid student %v", err)
	}

	// Call the Create function and check for an error.
	createdStudent, err := repo.Create(testStudent)
	if err != nil {
		t.Errorf("Create returned an unexpected error: %v", err)
	}

	fmt.Println(createdStudent.StudentID)
	fmt.Println(createdStudent.Email)

	// Verify the returned student is not nil.
	if createdStudent == nil {
		t.Fatal("Create returned a nil student")
	}

	// Assert that the returned student's ID matches the one we provided.
	if createdStudent.StudentID != testUUID {
		t.Errorf("Expected student ID %s, but got %s", testUUID, createdStudent.StudentID)
	}

	// Verify the student exists in the database by fetching it directly.
	var dbStudent postgres.DBStudent
	if err := db.First(&dbStudent, testUUID).Error; err != nil {
		t.Fatalf("Could not find created student in DB: %v", err)
	}

	// Compare the fields to ensure they were saved correctly.
	if dbStudent.FirstName != testStudent.FirstName {
		t.Errorf("Expected first name %s, but got %s", testStudent.FirstName, dbStudent.FirstName)
	}

	


}

// TestGormStudentRepo_FindById tests the FindById method for both success and failure cases.
func TestGormStudentRepo_FindById(t *testing.T) {
	// Setup the test database and get the repository instance.
	repo, db := setupTestDB(t)

	// --- Test Case 1: Successful find ---
	t.Run("successful find", func(t *testing.T) {
		// First, create a student directly in the database to ensure it exists.
		testUUID := uuid.New()
		now := time.Now()
		dbStudentToCreate := &postgres.DBStudent{
			StudentID:   testUUID,
			FirstName:   "Jane",
			LastName:    "Smith",
			DateOfBirth: &now,
			Email:       "jane.smith@example.com",
			EnrollmentDate: now,
		}
		if err := db.Create(dbStudentToCreate).Error; err != nil {
			t.Fatalf("Failed to seed database for test: %v", err)
		}

		// Now, use the repository to find the student by their ID.
		foundStudent, err := repo.FindById(testUUID)
		if err != nil {
			t.Errorf("FindById returned an unexpected error: %v", err)
		}
		
		

		// Verify the found student is not nil and the ID matches.
		if foundStudent == nil {
			t.Fatal("FindById returned a nil student")
		}
		if foundStudent.StudentID != testUUID {
			t.Errorf("Expected student ID %s, but got %s", testUUID, foundStudent.StudentID)
		}

		// Also check that a key field matches to ensure the correct record was retrieved.
		if foundStudent.FirstName != "Jane" {
			t.Errorf("Expected first name 'Jane', but got '%s'", foundStudent.FirstName)
		}

		_ , err = entities.NewValidatedStudent(foundStudent)
		if err != nil {
			t.Errorf("Expect no err, got %v" , err)
		}
	})

	// --- Test Case 2: Student not found ---
	t.Run("not found", func(t *testing.T) {
		// Use a new, random UUID that is guaranteed not to exist in our fresh test DB.
		nonExistentUUID := uuid.New()
		
		// Call FindById with the non-existent UUID.
		foundStudent, err := repo.FindById(nonExistentUUID)

		// Assert that the function returned a gorm.ErrRecordNotFound error.
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			t.Errorf("Expected 'record not found' error, but got '%v'", err)
		}

		// Assert that the returned student is nil.
		if foundStudent != nil {
			t.Errorf("Expected nil student, but got %+v", foundStudent)
		}
	})
}

func TestGormStudentRepo_FindAll(t *testing.T) {
	repo, db := setupTestDB(t)


	t.Run("successful find" , func(t *testing.T){
		testUUID := uuid.New()
		now := time.Now()
		dbStudentToCreate := &postgres.DBStudent{
			StudentID:   testUUID,
			FirstName:   "Jane",
			LastName:    "Smith",
			DateOfBirth: &now,
			Email:       "jane.smith@example.com",
		}
		if err := db.Create(dbStudentToCreate).Error; err != nil {
			t.Fatalf("Failed to seed database for test (create student 1): %v", err)
		}


		testUUID1 := uuid.New()
		
		dbStudentToCreate2 := &postgres.DBStudent{
			StudentID:   testUUID1,
			FirstName:   "tran",
			LastName:    "vu",
			DateOfBirth: &now,
			Email:       "tranvu@example.com",
		}
		if err := db.Create(dbStudentToCreate2).Error; err != nil {
			t.Fatalf("Failed to seed database for test (create student 2): %v", err)
		}

		foundStudents, err := repo.FindAll()
		if err != nil {
			t.Errorf("FindAll returned an unexpected error: %v", err)
		}
		
		// Verify the found student is not nil and the ID matches.
		if foundStudents == nil {
			t.Fatal("FindById returned a nil student")
		}
		if foundStudents[0].StudentID != testUUID {
			t.Errorf("Expected student ID %s, but got %s", testUUID, foundStudents[0].StudentID)
		}

		if foundStudents[1].StudentID != testUUID1 {
			t.Errorf("Expected student ID %s, but got %s", testUUID1, foundStudents[1].StudentID)
		}

		// Also check that a key field matches to ensure the correct record was retrieved.
		if foundStudents[0].FirstName != "Jane" {
			t.Errorf("Expected first name 'Jane', but got '%s'", foundStudents[0].FirstName)
		}


	})
}

func TestGormProductRepository_Update(t *testing.T){
	
	repo, db := setupTestDB(t)
	cleanup := func() {
		db.Exec("DELETE FROM db_students")
	}
	defer cleanup()
	
	now := time.Now()
	student := entities.NewStudent(
		"John","Doe",&now,"john.doe@aloalo.com",nil,nil,now,
	)

	validStudent, err := entities.NewValidatedStudent(student)
	if err != nil {
		t.Fatalf("Invalid student test case")
	}

	_, err = repo.Create(validStudent)

	if err != nil {
		t.Fatal("Failed to create a new student " + err.Error())
	}

	new_dob :=time.Date(2003,time.March,11,0,0,0,0,time.Local)
	phone := "0932323232"
	major := "CNTT"
	validStudent.UpdateNewFields(&new_dob,&phone,&major)

	_, err = repo.Update(validStudent)
	if err != nil {
		t.Fatalf("UpdateName failed or fetched wrong product")
	}

}

func TestGormProductRepository_Delete(t *testing.T){
	repo, db := setupTestDB(t)

	defer func () {
		db.Exec("DELETE FROM db_students")
	}()

	now := time.Now()
	student := entities.NewStudent(
		"John","Doe",&now,"john.doe@aloalo.com",nil,nil,now,
	)
	validStudent, err := entities.NewValidatedStudent(student)
	if err != nil {
		t.Fatalf("Invalid student input test: %v" , err)
	}


	_,err = repo.Create(validStudent)
	if err != nil {
		t.Fatalf("Cannot create new student: %v",err)
	}

	err = repo.Delete(validStudent.StudentID)
	if err != nil {
		t.Fatalf("Failed to delete a student: %v" ,err)
	}

}