package xentities

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

// PostgresBaseModel is a base struct for all PostgreSQL models with common fields
type PostgresBaseModel struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null;column:created_at" json:"created_at"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null;column:updated_at" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
}

// PostgresBaseModelWithTimestamp includes timestamps for created_by and updated_by
type PostgresBaseModelWithTimestamp struct {
	ID        uuid.UUID      `gorm:"type:uuid;primary_key;default:gen_random_uuid()" json:"id"`
	CreatedAt time.Time      `gorm:"autoCreateTime;not null;column:created_at" json:"created_at"`
	CreatedBy sql.NullString `json:"created_by,omitempty"`
	UpdatedAt time.Time      `gorm:"autoUpdateTime;not null;column:updated_at" json:"updated_at"`
	UpdatedBy sql.NullString `json:"updated_by,omitempty"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at,omitempty"`
	DeletedBy sql.NullString `json:"deleted_by,omitempty"`
}

func (u *PostgresBaseModel) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	u.ID = uuid

	return
}

func (u *PostgresBaseModelWithTimestamp) BeforeCreate(tx *gorm.DB) (err error) {
	uuid, err := uuid.NewRandom()
	if err != nil {
		return err
	}

	u.ID = uuid

	return
}
