# 设置变量
BINARY_NAME = ~/GolandProjects/bin/gsg

# 默认目标
default: build

# 构建目标
build:
	go build -o $(BINARY_NAME) .

# 清理目标
clean:
	rm -f $(BINARY_NAME)

.PHONY: build clean
