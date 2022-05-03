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

func (i *Identity) Id() string {
	return fmt.Sprintf("%+v-%+v", i.Prefix, i.Value)
}

func checkPrefix(prefix string) string {
	if prefix == "" {
		prefix = DEFAULT_PREFIX
	}
	regex := regexp.MustCompile("^[a-z]+$")
	match := regex.Match([]byte(prefix))
	if !match {
		err := fmt.Errorf("Prefix should only contain characters a-z (lowercase)")
		panic(err)
	}
	return prefix
}

func NewIdentityFrom(prefix string, id string) *Identity {
	uid, err := uuid.Parse(id)
	if err != nil {
		err := fmt.Errorf("'%+v' cannot be used as a Value for an Identity", id)
		panic(err)
	}
	prefix = checkPrefix(prefix)
	cleanId := strings.Replace(uid.String(), "-", "", -1)
	return &Identity{Prefix: prefix, Value: cleanId}
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
	val, err := CleanUuid()
	if err != nil {
		panic(err)
	}
	return &Identity{Prefix: prefix, Value: val}
}
