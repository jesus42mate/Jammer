build:
	@echo "Building a binary..."
	@go build -o ./target/runner.go update_apt_linux_22.04.go actions.go helpers.go

run:
	@echo "Turning on the gears..."
	@cd src/
	@go run src/update_apt_linux_22.04.go src/actions.go src/helpers.go
