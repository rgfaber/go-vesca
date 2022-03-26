package sdk

import (
	"github.com/google/uuid"
	"strings"
)

func cleanUuid() string {
	uid, _ := uuid.NewUUID()
	return strings.Replace(uid.String(), "-", "", -1)
}
