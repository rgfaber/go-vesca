package contract

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type FbkSerializer struct {
	Handler func(fbk interfaces.IFbk) ([]byte, error)
}

func (f *FbkSerializer) Serialize(fbk interfaces.IFbk) ([]byte, error) {
	return f.Handler(fbk)
}

func NewFbkSerializer(handler func(fbk interfaces.IFbk) ([]byte, error)) *FbkSerializer {
	return &FbkSerializer{
		Handler: handler,
	}
}
