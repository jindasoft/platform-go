package xentities

import (
	"context"
	"time"

	"github.com/jindasoft/jinda-platform/xauth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type MongoBefore interface {
	BeforeCreate(ctx context.Context) error
	BeforeUpdate(ctx context.Context) error
	BeforeSoftDelete(ctx context.Context) error
}

type MongoBase struct {
	ID        primitive.ObjectID  `json:"id,omitzero" bson:"_id,omitempty"`
	Status    Status              `json:"status,omitzero" bson:"status,omitempty"`
	CreatedAt time.Time           `json:"created_at,omitzero" bson:"created_at,omitempty"`
	CreatedBy primitive.ObjectID  `json:"created_by,omitzero" bson:"created_by,omitempty"`
	UpdatedAt time.Time           `json:"updated_at,omitzero" bson:"updated_at,omitempty"`
	UpdatedBy primitive.ObjectID  `json:"updated_by,omitzero" bson:"updated_by,omitempty"`
	DeletedAt *time.Time          `json:"deleted_at,omitzero" bson:"deleted_at,omitempty"`
	DeletedBy *primitive.ObjectID `json:"deleted_by,omitzero" bson:"deleted_by,omitempty"`
}

func (b *MongoBase) BeforeCreate(ctx context.Context) error {
	accountID, err := xauth.GetAccountOID(ctx)
	if err != nil {
		return err
	}

	b.ID = primitive.NewObjectID()
	b.Status = StatusActive
	b.CreatedAt = time.Now()
	b.CreatedBy = accountID
	b.UpdatedAt = time.Now()
	b.UpdatedBy = accountID

	return nil
}

func (b *MongoBase) BeforeUpdate(ctx context.Context) error {
	accountID, err := xauth.GetAccountOID(ctx)
	if err != nil {
		return err
	}

	b.UpdatedAt = time.Now()
	b.UpdatedBy = accountID

	return nil
}

func (b *MongoBase) BeforeSoftDelete(ctx context.Context) error {
	accountID, err := xauth.GetAccountOID(ctx)
	if err != nil {
		return err
	}

	now := time.Now()
	b.Status = StatusDeleted
	b.DeletedAt = &now
	b.DeletedBy = &accountID

	return nil
}
