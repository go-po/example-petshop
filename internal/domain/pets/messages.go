package pets

import (
	"encoding/json"

	"github.com/go-po/po"
)

type AddPetCmd struct {
	Name string
	Tags []string
}

type Added struct {
	Name string
	Tags []string
}

type Deleted struct{}

func init() {
	po.RegisterMessages(
		func(b []byte) (interface{}, error) {
			msg := AddPetCmd{}
			err := json.Unmarshal(b, &msg)
			return msg, err
		},
		func(b []byte) (interface{}, error) {
			msg := Added{}
			err := json.Unmarshal(b, &msg)
			return msg, err
		},
		func(b []byte) (interface{}, error) {
			msg := Deleted{}
			err := json.Unmarshal(b, &msg)
			return msg, err
		},
	)
}
