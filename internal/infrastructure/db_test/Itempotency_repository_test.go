package db_test

import (
	"context"
	"database/sql"
	
	
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/tranvu1111/go-students-new/internal/domain/entities"
	postgres2 "github.com/tranvu1111/go-students-new/internal/infrastructure/db/postgres"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupTestItempotencyDB(t *testing.T) (*gorm.DB, sqlmock.Sqlmock) {
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Failed to create sqlmock: %v", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm DB: %v", err)
	}

	return gormDB, mock
}

func TestGormIdempotencyRepo_FindByKey(t *testing.T) {
	db, mock := setupTestItempotencyDB(t)
	repo := postgres2.NewGormIdempotencyRepository(db)
	ctx := context.Background()

	t.Run("Found Record", func(t *testing.T) {
		recordID := uuid.New()
		expected := &entities.IdempotencyRecord{
			ID:         recordID,
			Key:        "test-key",
			Request:    `{"data":"test"}`,
			Response:   `{"result":"success"}`,
			StatusCode: 200,
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
		}

		rows := sqlmock.NewRows([]string{"id", "key", "request", "response", "status_code", "created_at"}).
			AddRow(
				expected.ID,
				expected.Key,
				expected.Request,
				expected.Response,
				expected.StatusCode,
				expected.CreatedAt,
			)

		mock.ExpectQuery(`SELECT \* FROM "db_idempotency_records" WHERE key = \$1 ORDER BY "db_idempotency_records"."id" LIMIT \$2`).
			WithArgs("test-key", 1).
			WillReturnRows(rows)

		result, err := repo.FindByKey(ctx, "test-key")
		require.NoError(t, err)
		assert.Equal(t, expected, result)
	})

	t.Run("Record Not Found", func(t *testing.T) {
		mock.ExpectQuery(`SELECT .* FROM "db_idempotency_records" WHERE key = \$1 ORDER BY "db_idempotency_records"."id" LIMIT \$2`).
			WithArgs("non-existent-key", 1).
			WillReturnError(gorm.ErrRecordNotFound)

		result, err := repo.FindByKey(ctx, "non-existent-key")
		require.NoError(t, err)
		assert.Nil(t, result)
	})

	t.Run("Database Error", func(t *testing.T) {
		mock.ExpectQuery(`SELECT .* FROM "db_idempotency_records" WHERE key = \$1 ORDER BY "db_idempotency_records"."id" LIMIT \$2`).
			WithArgs("error-key", 1).
			WillReturnError(sql.ErrConnDone)

		_, err := repo.FindByKey(ctx, "error-key")
		assert.Error(t, err)
		assert.Equal(t, sql.ErrConnDone, err)
	})
}

func TestGormIdempotencyRepo_Create(t *testing.T) {
	db, mock := setupTestItempotencyDB(t)
	repo := postgres2.NewGormIdempotencyRepository(db)
	ctx := context.Background()

	t.Run("Successful Create", func(t *testing.T) {
		record := &entities.IdempotencyRecord{
			ID:         uuid.New(),
			Key:        "create-test-key",
			Request:    `{"data":"create-test"}`,
			Response:   `{"result":"created"}`,
			StatusCode: 201,
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "db_idempotency_records" \("id","key","request","response","status_code","created_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6\)`).
			WithArgs(
				record.ID,
				record.Key,
				record.Request,
				record.Response,
				record.StatusCode,
				record.CreatedAt,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rows := sqlmock.NewRows([]string{"id", "key", "request", "response", "status_code", "created_at"}).
			AddRow(
				record.ID,
				record.Key,
				record.Request,
				record.Response,
				record.StatusCode,
				record.CreatedAt,
			)

		mock.ExpectQuery(`SELECT .* FROM "db_idempotency_records" WHERE id = \$1 ORDER BY "db_idempotency_records"."id" LIMIT \$2`).
			WithArgs(record.ID, 1).
			WillReturnRows(rows)

		created, err := repo.Create(ctx, record)
		require.NoError(t, err)
		assert.Equal(t, record, created)
	})

	t.Run("Create Database Error", func(t *testing.T) {
		record := &entities.IdempotencyRecord{
			ID:         uuid.New(),
			Key:        "create-error-key",
			Request:    `{"data":"test"}`,
			Response:   `{"result":"test"}`,
			StatusCode: 200,
			CreatedAt:  time.Now().UTC(),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`INSERT INTO "db_idempotency_records" \("id","key","request","response","status_code","created_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5,\$6\)`).
			WithArgs(
				record.ID,
				record.Key,
				record.Request,
				record.Response,
				record.StatusCode,
				record.CreatedAt,
			).
			WillReturnError(sql.ErrConnDone)
		mock.ExpectRollback()

		_, err := repo.Create(ctx, record)
		assert.Error(t, err)
		assert.Equal(t, sql.ErrConnDone, err)
	})
}

func TestGormIdempotencyRepo_Update(t *testing.T) {
	db, mock := setupTestItempotencyDB(t)
	repo := postgres2.NewGormIdempotencyRepository(db)
	ctx := context.Background()

	t.Run("Successful Update", func(t *testing.T) {
		record := &entities.IdempotencyRecord{
			ID:         uuid.New(),
			Key:        "update-test-key",
			Request:    `{"data":"updated"}`,
			Response:   `{"result":"updated"}`,
			StatusCode: 201,
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "db_idempotency_records" SET "key"=\$1,"request"=\$2,"response"=\$3,"status_code"=\$4,"created_at"=\$5 WHERE "id" = \$6`).
			WithArgs(
				record.Key,
				record.Request,
				record.Response,
				record.StatusCode,
				record.CreatedAt,
				record.ID,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		rows := sqlmock.NewRows([]string{"id", "key", "request", "response", "status_code", "created_at"}).
			AddRow(
				record.ID,
				record.Key,
				record.Request,
				record.Response,
				record.StatusCode,
				record.CreatedAt,
			)

		mock.ExpectQuery(`SELECT .* FROM "db_idempotency_records" WHERE id = \$1 ORDER BY "db_idempotency_records"."id" LIMIT \$2`).
			WithArgs(record.ID, 1).
			WillReturnRows(rows)

		updated, err := repo.Update(ctx, record)
		require.NoError(t, err)
		assert.Equal(t, record, updated)
		
	})

	
}

func TestGormIdempotencyRepo_Update_non_exist_student(t *testing.T) {
	db, mock := setupTestItempotencyDB(t)
	repo := postgres2.NewGormIdempotencyRepository(db)
	ctx := context.Background()

	t.Run("Fail Update", func(t *testing.T) {
		record := &entities.IdempotencyRecord{
			ID:         uuid.New(),
			Key:        "update-test-key",
			Request:    `{"data":"updated"}`,
			Response:   `{"result":"updated"}`,
			StatusCode: 201,
			CreatedAt:  time.Now().UTC().Truncate(time.Millisecond),
		}

		mock.ExpectBegin()
		mock.ExpectExec(`UPDATE "db_idempotency_records" SET "key"=\$1,"request"=\$2,"response"=\$3,"status_code"=\$4,"created_at"=\$5 WHERE "id" = \$6`).
			WithArgs(
				record.Key,
				record.Request,
				record.Response,
				record.StatusCode,
				record.CreatedAt,
				record.ID,
			).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// rows := sqlmock.NewRows([]string{"id", "key", "request", "response", "status_code", "created_at"}).
		// 	AddRow(
		// 		record.ID,
		// 		record.Key,
		// 		record.Request,
		// 		record.Response,
		// 		record.StatusCode,
		// 		record.CreatedAt,
		// 	)

		mock.ExpectQuery(`SELECT .* FROM "db_idempotency_records" WHERE id = \$1 ORDER BY "db_idempotency_records"."id" LIMIT \$2`).
			WithArgs(
				record.ID, // new ID that non-exist => expect not found
				1,
			).
			WillReturnError(gorm.ErrRecordNotFound)

		updated, err := repo.Update(ctx, record)
		// require.Error(t, err, "Error should be gorm.ErrRecordNotFound")
		assert.ErrorIs(t, err, gorm.ErrRecordNotFound, "Error should be gorm.ErrRecordNotFound")
		assert.Nil(t, updated, "No record should be returned on error")
		

		require.NoError(t, mock.ExpectationsWereMet(), "All mock expectations should be fulfilled")
	})

	
}