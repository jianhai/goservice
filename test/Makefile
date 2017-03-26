.PHONY: default

default: build

build:
	@go build -o LiveTopics LiveTopics.go
	@go build -o LiveGivings LiveGivings.go
	@go build -o UserRegister UserRegister.go
	@go build -o UserSignin UserSignin.go
	@go build -o UserModify UserModify.go

clean:
	@rm -rf LiveTopics
	@rm -rf LiveGivings
	@rm -rf UserRegister
	@rm -rf UserSignin
	@rm -rf UserModify

test: build
	@./runtest.sh
