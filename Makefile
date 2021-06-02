all: run

run:
	@go test

linux:
	@GOOS=linux go run examples/*.go

web:
	@GOOS=js GOARCH=wasm go build -o bin.wasm examples/*.go
	@mv bin.wasm examples/web/wasm/
	@cd examples/web && npx http-server -o

android:
	@cd examples && GO111MODULE=on gomobile build -target=android main.go