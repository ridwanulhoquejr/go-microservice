server:
	go run main.go

mod:
	go mod tidy

.PHONY: server, mod