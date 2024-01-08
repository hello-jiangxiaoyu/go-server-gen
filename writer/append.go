package writer

import (
	"go-server-gen/utils"
	"io"
	"os"
	"strings"
)

func FileAppendWriter(path string, src string, handlers map[string]string) error {
	if _, err := os.Stat(path); err != nil {
		return writeFile(path, []byte(src), false)
	}

	f, err := os.OpenFile(path, os.O_APPEND|os.O_RDWR, 0644)
	if err != nil {
		return err
	}

	defer utils.DeferErr(f.Close)
	oldCode, err := io.ReadAll(f)
	if err != nil {
		return err
	}
	for key, handler := range handlers {
		if !strings.Contains(string(oldCode), key) {
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

	handlerSlice := filterPointerHandlerSlice(handlers, string(content))
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

func filterPointerHandlerSlice(obj map[string]string, src string) []string {
	res := make([]string, 0, len(obj))
	for k, v := range obj {
		if !strings.Contains(src, k) {
			res = append(res, v)
		}
	}
	return res
}
