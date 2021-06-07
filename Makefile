all: run

run:
	@go test

linux:
	@GOOS=linux go run examples/desktop/*.go

web:
	@GOOS=js GOARCH=wasm go build -o bin.wasm examples/desktop/*.go
	@mv bin.wasm examples/desktop/web/wasm/
	@cd examples/desktop/web && npx http-server -o

android: build-android install-android stop-android run-android debug-android

build-android:
	@cd examples/mobile && GO111MODULE=on gomobile build -androidapi 26 -target=android/arm --tags android .

install-android:
	@adb install -r examples/mobile/mobile.apk

stop-android:
	@adb shell am force-stop dev.fragoso.mustard

run-android:
	@adb shell am start -n dev.fragoso.mustard/org.golang.app.GoNativeActivity

debug-android:
	@adb logcat --pid=`bash examples/mobile/get_pid.sh`
