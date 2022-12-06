run: build
	./2022

build: cover
	go build

test:
	go fmt
	go test -v  -coverprofile cover.out -cover

cover: test
	go tool cover -html=cover.out
