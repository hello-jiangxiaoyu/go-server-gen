package writer

import (
	"errors"
	"go-server-gen/utils"
)

const (
	Overwrite = "overwrite"
	Skip      = "skip"
	Append    = "append"
	Pointer   = "pointer"
)

type WriteCode struct {
	File     string
	Write    string // overwrite, skip, append, pointer
	Code     string
	Handlers map[string]string
}

func Write(codes map[string]WriteCode) error {
	var err error
	for _, code := range codes {
		switch code.Write {
		case Overwrite:
			err = writeFile(code.File, []byte(code.Code), true)
		case Skip:
			err = writeFile(code.File, []byte(code.Code), false)
		case Append:
			err = FileAppendWriter(code.File, code.Code, code.Handlers)
		case Pointer:
			err = PointerAppendWriter(code.File, "//INSERT_POINT", code.Code, code.Handlers)
		default:
			return errors.New("no such writer")
		}
		if err != nil {
			return utils.WithMessage(err, "write "+code.File+" err")
		}
	}

	return nil
}
