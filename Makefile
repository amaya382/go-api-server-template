export GO111MODULE=on

example: clean
	go build ./cmd/example

migrate: clean
	go build ./cmd/migration
	./migration

.PHONY: update
update:
	go mod tidy

.PHONY: init
init:
	go get -u github.com/swaggo/swag/cmd/swag

.PHONY: clean
clean:
	rm -f example migration || :
