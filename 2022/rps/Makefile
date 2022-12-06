test:
	go fmt
	go test -v  -coverprofile cover.out -cover

cover: test
	go tool cover -html=cover.out
