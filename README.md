# DoubleAgent
Sample android agent written in Go

### Compile environment variables:
- export NDK=$ANDROID_NDK_HOME
- export TARGET=aarch64-linux-android
- export API=34
- export CC=$NDK/toolchains/llvm/prebuilt/darwin-x86_64/bin/$TARGET$API-clang
- export CXX=$NDK/toolchains/llvm/prebuilt/darwin-x86_64/bin/$TARGET$API-clang++

### Compile command for Android:
- env GOOS=android GOARCH=arm64 CGO_ENABLED=1 go build -o myapp main.go

