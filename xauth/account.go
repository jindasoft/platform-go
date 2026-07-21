package xauth

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/jindasoft/jinda-platform/xconst"
	"github.com/jindasoft/jinda-platform/xutils"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetAccountOID(ctx context.Context) (primitive.ObjectID, error) {
	oid, ok := ctx.Value(xconst.ContextAccountOID).(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("GetAccountOID: AccountID not found in context")
	}

	return oid, nil
}

func GetAccountUUID(ctx context.Context) (uuid.UUID, error) {
	uid, ok := ctx.Value(xconst.ContextAccountUUID).(uuid.UUID)
	if !ok {
		return uuid.Nil, fmt.Errorf("GetAccountUUID: AccountID not found in context")
	}

	return uid, nil
}

func GetAccountUUIDOrDefault(ctx context.Context) *uuid.UUID {
	str, ok := ctx.Value(xconst.ContextAccountUUID).(string)
	if !ok {
		return nil
	}

	return xutils.StringToUuidOrDefault(str)
}

func CheckOwnership(ctx context.Context, accountOID primitive.ObjectID) (bool, error) {
	oid, err := GetAccountOID(ctx)
	if err != nil {
		return false, fmt.Errorf("unauthorized")
	}

	return oid == accountOID, nil
}
