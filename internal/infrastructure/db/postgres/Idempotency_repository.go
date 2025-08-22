package postgres

import (
	"context"

	"github.com/tranvu1111/go-students-new/internal/domain/entities"
	"github.com/tranvu1111/go-students-new/internal/domain/repositories"
	// "github.com/tranvu1111/go-students-new/internal/domain/entities"
	"gorm.io/gorm"
)

type GormIdempotencyRepo struct {
	db *gorm.DB
}


func NewGormIdempotencyRepository(db *gorm.DB) repositories.IdempotencyRepository {
	return &GormIdempotencyRepo{db:db}
}

func (repo *GormIdempotencyRepo) FindByKey(ctx context.Context, key string) (*entities.IdempotencyRecord, error) {
	var dbRecord DBIdempotencyRecord
	result := repo.db.WithContext(ctx).Where("key = ?", key).First(&dbRecord)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}

		return nil, result.Error
	}

	return &entities.IdempotencyRecord{
		ID:         dbRecord.ID,
		Key:        dbRecord.Key,
		Request:    dbRecord.Request,
		Response:   dbRecord.Response,
		StatusCode: dbRecord.StatusCode,
		CreatedAt:  dbRecord.CreatedAt,
	}, nil
}

func (repo *GormIdempotencyRepo) Create(ctx context.Context, record *entities.IdempotencyRecord) (*entities.IdempotencyRecord, error) {
	dbRecord := DBIdempotencyRecord{
		ID:         record.ID,
		Key:        record.Key,
		Request:    record.Request,
		Response:   record.Response,
		StatusCode: record.StatusCode,
		CreatedAt:  record.CreatedAt,
	}

	result := repo.db.WithContext(ctx).Create(&dbRecord)
	if result.Error != nil {
		return nil,result.Error
	}

	var createdRecord DBIdempotencyRecord
	if err := repo.db.WithContext(ctx).Where("id = ?", dbRecord.ID).First(&createdRecord).Error;err != nil {
		return nil, err
	}

	return &entities.IdempotencyRecord{
		ID:	createdRecord.ID,
		Key: createdRecord.Key,
		Request: createdRecord.Request,
		Response: createdRecord.Response,
		StatusCode: createdRecord.StatusCode,
		CreatedAt: createdRecord.CreatedAt,
	},nil
}

func (repo *GormIdempotencyRepo) Update(ctx context.Context, record *entities.IdempotencyRecord) (*entities.IdempotencyRecord, error){
	dbRecord := DBIdempotencyRecord{
		ID:         record.ID,
		Key:        record.Key,
		Request:    record.Request,
		Response:   record.Response,
		StatusCode: record.StatusCode,
		CreatedAt:  record.CreatedAt,
	}

	result := repo.db.WithContext(ctx).Save(&dbRecord)
	if result.Error != nil {
		return nil, result.Error
	}

	// Read back the updated record
	var updatedRecord DBIdempotencyRecord
	if err := repo.db.WithContext(ctx).Where("id = ?", dbRecord.ID).First(&updatedRecord).Error; err != nil {
		return nil, err
	}

	return &entities.IdempotencyRecord{
		ID:         updatedRecord.ID,
		Key:        updatedRecord.Key,
		Request:    updatedRecord.Request,
		Response:   updatedRecord.Response,
		StatusCode: updatedRecord.StatusCode,
		CreatedAt:  updatedRecord.CreatedAt,
	}, nil
}

