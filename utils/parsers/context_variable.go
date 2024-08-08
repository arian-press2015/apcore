package parsers

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ParseUUIDFromContext retrieves a UUID variable from the Gin context
func ParseUUIDFromContext(c *gin.Context, variableName string) (uuid.UUID, error) {
	variable, exists := c.Get(variableName)

	if !exists {
		return uuid.Nil, errors.New("variable not found")
	}

	parsedUUID, ok := variable.(uuid.UUID)
	if !ok {
		return uuid.Nil, errors.New("invalid variable type, expected uuid")
	}

	return parsedUUID, nil
}
