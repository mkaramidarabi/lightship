BUILD_FLAGS="-v"
.PHONY: build debug clean test install
build:
	@go build $(BUILD_FLAGS) -trimpath -o release/lightship
debug:
	@go build -race $(BUILD_FLAGS) -gcflags "-N -l" -o release/lightship
clean:
	@go clean --cache
test: 
	@go test -v ./... -count=1 -cover
install:
	@if ! [ "$(shell id -u)" = 0 ]; then \
		echo "You are not root, run this target as root please"; \
		exit 1; \
	fi
	@cp ./release/* /usr/bin/
