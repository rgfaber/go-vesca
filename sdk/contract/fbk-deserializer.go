package contract

import (
	"github.com/rgfaber/go-vesca/sdk/interfaces"
)

type FbkDeserializer struct {
	Handler func([]byte) interfaces.IFbk
}
