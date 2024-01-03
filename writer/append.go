package writer

import (
	"go-server-gen/utils"
	"os"
	"strings"
)

func FileAppendWriter(path string, src string, handlers map[string]string) error {
	if _, err := os.Stat(path); err != nil {
		return writeFile(path, []byte(src), false)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	defer utils.DeferErr(f.Close)
	for key, handler := range handlers {
		if !strings.Contains(src, key) {
			if _, err = f.WriteString(handler); err != nil {
				return err
			}
		}
	}

	return nil
}

// PointerAppendWriter 在文件指针下方插入代码
func PointerAppendWriter(path string, pointer string, src string, handlers map[string]string) error {
	if _, err := os.Stat(path); err != nil {
		return writeFile(path, []byte(src), false)
	}

	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	handlerSlice := mapStringToSlice(handlers)
	lines := strings.Split(string(content), "\n")
	for i, line := range lines {
		if strings.Contains(line, pointer) {
			lines = append(lines[:i+1], append(handlerSlice, lines[i+1:]...)...)
			break
		}
	}

	output := strings.Join(lines, "\n")
	if err = os.WriteFile(path, []byte(output), 0644); err != nil {
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
