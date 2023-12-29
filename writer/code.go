package writer

import (
	"go-server-gen/utils"
	"os"
)

type WriteCode struct {
	File     string
	Write    string
	Code     string
	Handlers map[string]string
}

func Write(codes map[string]WriteCode) error {
	for _, code := range codes {
		println(code.File)
	}

	return nil
}

func WriteGoCode(dir, file, src string, overwrite bool) error {
	path := dir + "/" + file + ".go"
	_, err := os.Stat(path)
	if err == nil && !overwrite {
		println(path + " is already exist, skip !!!!!!")
		return nil
	}

	if err != nil && !os.IsNotExist(err) {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}

	if _, err = f.Write([]byte(src)); err != nil {
		utils.DeferErr(f.Close)
		return err
	}

	utils.DeferErr(f.Close)
	println("generate " + path + " success!")

	return nil
}
