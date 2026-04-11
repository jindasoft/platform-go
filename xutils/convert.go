package xutils

import (
	"encoding/json"

	"github.com/google/uuid"
)

func JsonToStringOrDefault(data any) string {
	jsonData, err := json.Marshal(data)
	if err != nil {
		return ""
	}

	return string(jsonData)
}

func StringToUuidOrDefault(str string) *uuid.UUID {
	tmp, err := uuid.Parse(str)
	if err != nil {
		return nil
	}

	return &tmp
}
