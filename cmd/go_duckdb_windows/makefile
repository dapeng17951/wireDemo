# 不支持 386
mingw-build:mingw-build-windows-amd64 mingw-build-windows-arm mingw-build-windows-arm64\
mingw-build-linux-amd64 mingw-build-linux-arm mingw-build-linux-arm64 mingw-build-linux-loong64 mingw-build-linux-riscv64

mingw-build-windows-amd64:
	-CGO_ENABLED=1 GOOS=windows GOARCH=amd64 CC="gcc" CXX="g++" go build -o main_windows-amd64.exe *.go
mingw-build-windows-arm:
	-CGO_ENABLED=1 GOOS=windows GOARCH=arm CC="gcc" CXX="g++" go build -o main_windows-arm.exe *.go
mingw-build-windows-arm64:
	-CGO_ENABLED=1 GOOS=windows GOARCH=arm64 CC="gcc" CXX="g++" go build -o main_windows-arm64.exe *.go
mingw-build-linux-amd64:
	-CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="gcc" CXX="g++" go build -o main_linux-amd64 *.go
mingw-build-linux-arm:
	-CGO_ENABLED=1 GOOS=linux GOARCH=arm CC="gcc" CXX="g++" go build -o main_linux-arm *.go
mingw-build-linux-arm64:
    -CGO_ENABLED=1 GOOS=linux GOARCH=arm64 CC="gcc" CXX="g++" go build -o main_linux-arm64 *.go
mingw-build-linux-loong64:
    -CGO_ENABLED=1 GOOS=linux GOARCH=loong64 CC="gcc" CXX="g++" go build-o main_linux-loong64 *.go
mingw-build-linux-riscv64:
	-CGO_ENABLED=1 GOOS=linux GOARCH=riscv64 CC="gcc" CXX="g++" go build -o main_linux-riscv64 *.go

# 编译出错
zig-build:
	CGO_ENABLED=1 GOOS=linux GOARCH=amd64 CC="zig cc" CXX="zig c++" go build *.go

clean:
	-rm -rf *.exe