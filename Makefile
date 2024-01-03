# 设置变量
BINARY_NAME = ~/GolandProjects/bin/gsg
MAIN_PATH = main.go

# 默认目标
default: build

# 构建目标
build:
	go build -o $(BINARY_NAME) $(MAIN_PATH)

# 清理目标
clean:
	rm -f $(BINARY_NAME)

.PHONY: build clean
