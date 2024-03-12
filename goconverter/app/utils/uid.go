package utils

import (
	"github.com/google/uuid"
)

func GetUid() uuid.UUID {
	return uuid.New()
}
