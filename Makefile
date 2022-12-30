dev:
	go run main.go

test-unit:
	go test -tags=unit -v ./...

test-coverage:
	go test -cover ./...

test-integration:
	docker-compose -f docker-compose.it-test.yaml down && \
	docker-compose -f docker-compose.it-test.yaml up --build --force-recreate --abort-on-container-exit --exit-code-from it_tests
