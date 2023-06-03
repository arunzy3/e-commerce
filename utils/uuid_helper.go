package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateUUID(prefix string) string {
	uuidNew := uuid.New().String() + uuid.New().String()
	uuidNew = prefix + "_" + strings.ReplaceAll(uuidNew, "-", "")
	return uuidNew[:36]
}
