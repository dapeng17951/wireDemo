package wireCleanUp

import (
	"github.com/google/wire"
)

type File string

func NewFile() (string, func(), error) {
	return "file", func() {
		println("file clean up func")
	}, nil
}

var WireSet = wire.NewSet(NewFile)
