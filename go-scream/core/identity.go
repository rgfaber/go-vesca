package core

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/rgfaber/go-vesca/sdk/core/test"
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
		prefix = test.DEFAULT_PREFIX
	}
	regex := regexp.MustCompile("^[a-z]+$")
	match := regex.Match([]byte(prefix))
	if !match {
		err := fmt.Errorf("Prefix should only contain characters a-z (lowercase)")
		panic(err)
	}
	return prefix
}

func checkValue(value string) string {
	if value == "" {
		value = test.CLEAN_NULL_UUID
	}
	value = strings.Replace(value, "-", "", -1)
	value = strings.ToLower(value)
	regex := regexp.MustCompile("^[a-z\\d]+$")
	match := regex.Match([]byte(value))
	if !match {
		err := fmt.Errorf("Value should only contain characters 0-9 and a-z (lowercase)")
		panic(err)
	}
	return value
}

func NewIdentityFrom(prefix string, id string) *Identity {
	uid, err := uuid.Parse(id)
	if err != nil {
		err := fmt.Errorf("'%+v' cannot be used as a Value for an Identity", id)
		panic(err)
	}
	prefix = checkPrefix(prefix)
	value := checkValue(uid.String())
	return &Identity{Prefix: prefix, Value: value}
}

func NewIdentity(prefix string) *Identity {
	prefix = checkPrefix(prefix)
	val, err := CleanUuid()
	if err != nil {
		panic(err)
	}
	return &Identity{Prefix: prefix, Value: val}
}

func NewIdentityFromAggregateId(aggregateId string) *Identity {
	parts := strings.Split(aggregateId, "-")
	prefix := checkPrefix(parts[0])
	value := checkValue(parts[1])
	return &Identity{
		Prefix: prefix,
		Value:  value,
	}
}
