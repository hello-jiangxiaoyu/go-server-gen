# 设置变量
BINARY_NAME = gsg

# 构建目标
build:
	go build -ldflags="-s -w" -o $(BINARY_NAME) .

# 清理目标
clean:
	rm -f $(BINARY_NAME)

.PHONY: build clean
