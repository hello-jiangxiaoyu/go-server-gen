package writer

import "errors"

type WriteCode struct {
	File     string
	Write    string // overwrite, skip, append, pointer
	Code     string
	Handlers map[string]string
}

func Write(codes map[string]WriteCode) error {
	for _, code := range codes {
		switch code.Write {
		case "overwrite":
			println(code.File + " overwrite\n" + code.Code)
		case "skip":
			println(code.File + " skip\n" + code.Code)
		case "append":
			println(code.File + " append\n" + code.Code)
		case "pointer":
			println(code.File + " pointer\n" + code.Code)
		default:
			return errors.New("no such writer")
		}
	}

	return nil
}
