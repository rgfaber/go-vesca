package sdk

import (
	"fmt"
	"github.com/google/uuid"
	"regexp"
	"strings"
)

type Identity struct {
	prefix string `json:"prefix"`
	value  string `json:"value"`
}

func (i *Identity) Value() string {
	return i.value
}

func (i *Identity) Prefix() string {
	return i.prefix
}

func (i *Identity) Id() string {
	return fmt.Sprintf("%+v-%+v", i.prefix, i.value)
}

func checkPrefix(prefix string) string {
	if prefix == "" {
		prefix = DEFAULT_PREFIX
	}
	regex := regexp.MustCompile("^[a-z]+$")
	match := regex.Match([]byte(prefix))
	if !match {
		err := fmt.Errorf("prefix should only contain characters a-z (lowercase)")
		panic(err)
	}
	return prefix
}

func NewIdentityFrom(prefix string, id string) *Identity {
	uid, err := uuid.Parse(id)
	if err != nil {
		err := fmt.Errorf("'%+v' cannot be used as a value for an Identity", id)
		panic(err)
	}
	prefix = checkPrefix(prefix)
	cleanId := strings.Replace(uid.String(), "-", "", -1)
	return &Identity{prefix: prefix, value: cleanId}
}

func NewIdentity(prefix string) *Identity {
	if prefix == "" {
		prefix = "id"
	}
	regex := regexp.MustCompile("^[a-z_]+$")
	match := regex.Match([]byte(prefix))
	if !match {
		err := fmt.Errorf("prefix should only contain characters a-z (lowercase) and underscores")
		panic(err)
	}
	val, err := CleanUuid()
	if err != nil {
		panic(err)
	}
	return &Identity{prefix: prefix, value: val}
}
