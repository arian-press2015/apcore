package parsers

import (
	"errors"

	"github.com/google/uuid"
)

func ParseUUID(uuidString string) (uuid.UUID, error) {
	parsedUUID, err := uuid.Parse(uuidString)
	if err != nil {
		return uuid.Nil, errors.New("invalid UUID format")
	}

	return parsedUUID, nil
}
