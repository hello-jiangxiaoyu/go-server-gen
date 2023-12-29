package writer

import (
	"fmt"
	"os"
)

func makeDir(folder string) error {
	_, err := os.Stat(folder)
	if err == nil {
		return nil // 目录已存在
	}
	if !os.IsNotExist(err) {
		return err
	}

	fmt.Printf("### %s is not exist, creating! ", folder)
	if err = os.Mkdir(folder, os.ModePerm); err != nil {
		return err
	}

	return nil
}

func mapStringToSlice(obj map[string]string) []string {
	res := make([]string, 0, len(obj))
	for _, v := range obj {
		res = append(res, v)
	}
	return res
}
