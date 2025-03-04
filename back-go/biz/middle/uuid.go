package middle

import (
	"strings"

	"github.com/google/uuid"
)

func NewUUID() string {
	newUUID := uuid.New()
	return strings.ReplaceAll(newUUID.String(), "-", "")
}
