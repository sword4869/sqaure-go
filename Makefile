# 自定义变量
APP = hello

run:
	go run 2.go
	@echo "\t√ 2.go"
build:
	go build -o 2 2.go
clean:
	@echo "\t√ default"
	-rm he*
	-rm output/*
all: build1 windows linux darwin

windows:
	@echo "\t√ windows"
	@GOOS=windows go build -o ${APP}-windows mytest/2.go
linux:
	@echo "\t√ linux"
	@GOOS=linux go build -o ${APP}-linux mytest/2.go
darwin:
	@echo "\t√ mac"
	@GOOS=darwin GOARCH=arm64 go build -o ${APP}-darwin mytest/2.go