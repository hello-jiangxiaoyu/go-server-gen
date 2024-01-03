package writer

import (
	"go-server-gen/utils"
	"os"
	"path/filepath"
)

func writeFile(path string, src []byte, overwrite bool) error {
	_, err := os.Stat(path)
	if err != nil && !os.IsNotExist(err) {
		return err
	}

	if os.IsNotExist(err) {
		if err = createFile(path); err != nil {
			return err
		}
		if err = os.WriteFile(path, src, 0644); err != nil {
			return err
		}
		return nil
	}

	if !overwrite {
		return nil // 文件已存在且不覆盖
	}
	f, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer utils.DeferErr(f.Close)
	if _, err = f.Write(src); err != nil {
		return err
	}

	return nil
}

func createFile(file string) error {
	if err := os.MkdirAll(filepath.Dir(file), os.ModePerm); err != nil {
		return err
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	return f.Close()
}
