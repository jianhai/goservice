.PHONY: default

default: build

build:
	@go build -o UserLogin UserLogin.go
	@go build -o UserModify UserModify.go

clean:
	@rm -rf UserLogin	

test: build
	@./runtest.sh
