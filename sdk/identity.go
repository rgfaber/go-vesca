package sdk

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

type Identity struct {
	Prefix string `json:"prefix"`
	Value  string `json:"value"`
}

func (i *Identity) Id() (string string) {
	return fmt.Sprintf("%+v-%+v", i.Prefix, i.Value)
}

func NewIdentity(prefix string) *Identity {
	if prefix == "" {
		prefix = "id"
	}
	regex := regexp.MustCompile("^[a-z_]+$")
	match := regex.Match([]byte(prefix))
	if !match {
		err := fmt.Errorf("Prefix should only contain characters a-z (lowercase) and underscores")
		panic(err)
	}
	val, err := cleanUuid()
	if err != nil {
		panic(err)
	}
	return &Identity{Prefix: prefix, Value: val}
}

func newUuid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error in cleanUid:", err)
	}
	return uid.String(), err
}

func nullId() string {
	return uuid.Nil.String()
}

func cleanNullUuid() string {
	return strings.Replace(uuid.Nil.String(), "-", "", -1)
}

func cleanUuid() (string, error) {
	uid, err := uuid.NewUUID()
	if err != nil {
		fmt.Println("Error in cleanUid:", err)
	}
	return strings.Replace(uid.String(), "-", "", -1), err
}
