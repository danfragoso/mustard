all: run

run:
	@go test

linux:
	@GOOS=linux go run examples/desktop/*.go

web:
	@GOOS=js GOARCH=wasm go build -o bin.wasm examples/desktop/*.go
	@mv bin.wasm examples/desktop/web/wasm/
	@cd examples/desktop/web && npx http-server -o

android:
	@cd examples/mobile && GO111MODULE=on gomobile build -target=android --tags android .