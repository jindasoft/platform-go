package xdb

import (
	"context"

	"gorm.io/gorm"
)

type PostgresOperations struct {
	svc PostgresService
}

func NewPostgresOperations(svc PostgresService) *PostgresOperations {
	return &PostgresOperations{svc: svc}
}

// Find retrieves multiple records based on conditions
func (op *PostgresOperations) Find(ctx context.Context, entity any, conditions ...interface{}) error {
	return op.svc.GetDB().WithContext(ctx).Where(conditions[0], conditions[1:]...).Find(entity).Error
}

// FindByID retrieves a single record by primary key
func (op *PostgresOperations) FindByID(ctx context.Context, id interface{}, entity any) error {
	return op.svc.GetDB().WithContext(ctx).First(entity, id).Error
}

// FindOne retrieves a single record based on conditions
func (op *PostgresOperations) FindOne(ctx context.Context, entity any, conditions ...interface{}) error {
	return op.svc.GetDB().WithContext(ctx).Where(conditions[0], conditions[1:]...).First(entity).Error
}

// FindPaging retrieves paginated records
func (op *PostgresOperations) FindPaging(ctx context.Context, entity any, offset, limit int, results interface{}, conditions ...interface{}) (*PaginationResult, error) {
	db := op.svc.GetDB().WithContext(ctx)

	if len(conditions) > 0 {
		db = db.Where(conditions[0], conditions[1:]...)
	}

	var total int64
	if err := db.Model(entity).Count(&total).Error; err != nil {
		return nil, err
	}

	if err := db.Offset(offset).Limit(limit).Find(results).Error; err != nil {
		return nil, err
	}

	return &PaginationResult{
		Total:  total,
		Offset: int64(offset),
		Limit:  int64(limit),
	}, nil
}

// Count returns the count of records matching conditions
func (op *PostgresOperations) Count(ctx context.Context, entity any, conditions ...interface{}) (int64, error) {
	var count int64
	db := op.svc.GetDB().WithContext(ctx)

	if len(conditions) > 0 {
		db = db.Where(conditions[0], conditions[1:]...)
	}

	if err := db.Model(entity).Count(&count).Error; err != nil {
		return 0, err
	}

	return count, nil
}

// Create inserts a new record
func (op *PostgresOperations) Create(ctx context.Context, entity any) error {
	return op.svc.GetDB().WithContext(ctx).Create(entity).Error
}

// CreateBatch inserts multiple records
func (op *PostgresOperations) CreateBatch(ctx context.Context, entities any, batchSize int) error {
	return op.svc.GetDB().WithContext(ctx).CreateInBatches(entities, batchSize).Error
}

// Update updates a record
func (op *PostgresOperations) Update(ctx context.Context, entity any, updates interface{}, conditions ...interface{}) error {
	db := op.svc.GetDB().WithContext(ctx)

	if len(conditions) > 0 {
		db = db.Where(conditions[0], conditions[1:]...)
	}

	return db.Model(entity).Updates(updates).Error
}

// UpdateByID updates a record by primary key
func (op *PostgresOperations) UpdateByID(ctx context.Context, id interface{}, updates interface{}, entity any) error {
	return op.svc.GetDB().WithContext(ctx).Model(entity).Where("id = ?", id).Updates(updates).Error
}

// Delete deletes a record (soft delete - marks deleted_at)
func (op *PostgresOperations) Delete(ctx context.Context, entity any, conditions ...interface{}) error {
	db := op.svc.GetDB().WithContext(ctx)

	if len(conditions) > 0 {
		db = db.Where(conditions[0], conditions[1:]...)
	}

	return db.Delete(entity).Error
}

// DeleteByID deletes a record by primary key (soft delete - marks deleted_at)
func (op *PostgresOperations) DeleteByID(ctx context.Context, id interface{}, entity any) error {
	return op.svc.GetDB().WithContext(ctx).Delete(entity, id).Error
}

// ForceDelete permanently deletes a record (hard delete)
func (op *PostgresOperations) ForceDelete(ctx context.Context, entity any, conditions ...interface{}) error {
	db := op.svc.GetDB().WithContext(ctx).Unscoped()

	if len(conditions) > 0 {
		db = db.Where(conditions[0], conditions[1:]...)
	}

	return db.Delete(entity).Error
}

// ForceDeleteByID permanently deletes a record by primary key (hard delete)
func (op *PostgresOperations) ForceDeleteByID(ctx context.Context, id interface{}, entity any) error {
	return op.svc.GetDB().WithContext(ctx).Unscoped().Delete(entity, id).Error
}

// Restore restores a soft-deleted record
func (op *PostgresOperations) Restore(ctx context.Context, entity any, conditions ...interface{}) error {
	db := op.svc.GetDB().WithContext(ctx).Unscoped()

	if len(conditions) > 0 {
		db = db.Where(conditions[0], conditions[1:]...)
	}

	return db.Model(entity).Update("deleted_at", nil).Error
}

// Transaction executes operations within a transaction
func (op *PostgresOperations) Transaction(ctx context.Context, fn func(*PostgresOperations) error) error {
	tx := op.svc.GetDB().WithContext(ctx).Begin()
	if tx.Error != nil {
		return tx.Error
	}

	txOp := NewPostgresOperations(&postgresService{db: tx})
	if err := fn(txOp); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Paginate is a helper for pagination in GORM
func (op *PostgresOperations) Paginate(offset, limit int) func(*gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(offset).Limit(limit)
	}
}

type PaginationResult struct {
	Total  int64
	Offset int64
	Limit  int64
}
