package xmodels

import "go.mongodb.org/mongo-driver/bson/primitive"

type BasicInfo struct {
	ID   primitive.ObjectID `json:"id"`
	Code string             `json:"code"`
	Name string             `json:"name"`
}
