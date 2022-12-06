run: build
	./2022

build: test
	go build

test:
	go fmt
	go test

cover: test
	go test -v  -coverprofile cover.out -cover
	go tool cover -html=cover.out
