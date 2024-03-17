package validate

import (
	"github.com/google/uuid"
)

func ValidateGuid(guid string) error {
	// GUID is actually the same 128-bit identifier as UUID but in context of Windows OS
	return uuid.Validate(guid)
}
