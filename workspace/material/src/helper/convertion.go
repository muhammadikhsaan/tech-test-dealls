package helper

import (
	"encoding/json"
	"net/http"

	"dealls.test/material/src/core"
)

func ToPointer[T comparable](value T) *T {
	return &value
}

func ToMap(v any) (map[string]any, *core.Error) {
	var val map[string]any

	j, err := json.Marshal(v)

	if err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "Failed to marshal json",
		}
	}

	if err := json.Unmarshal(j, &val); err != nil {
		return nil, &core.Error{
			StatusCode: http.StatusInternalServerError,
			Origin:     err,
			Message:    "Failed to unmarshal json",
		}
	}

	return val, nil
}
