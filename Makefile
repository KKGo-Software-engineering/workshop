dev:
	go run main.go

test:
	go test -v ./...

test-coverage:
	go test -cover ./...