package utils

import (
	"fmt"

	"github.com/google/uuid"
)

func UUIDParser(input string) (uuid.UUID, error) {
	id, err := uuid.Parse(input)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("error parsing UUID: %v", err)
	}
	return id, nil
}
