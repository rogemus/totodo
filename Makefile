.PHONY: test tui

test:
	@echo "Testing..."
	go test ./test/... -v

tui:
	@echo "Starting TUI..."
	DEBUG=true go run main.go

