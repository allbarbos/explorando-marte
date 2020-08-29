fmt:
	go fmt ./...

test: fmt
	go test ./... -cover -coverprofile=cover.out

cov: test
	go tool cover -html=cover.out

run:
	cd cmd && go run main.go