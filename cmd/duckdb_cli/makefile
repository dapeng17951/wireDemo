.DEFAULT_GOAL = windows-amd64-debug

debug:windows-386-debug windows-amd64-debug windows-arm-debug windows-arm64-debug\
linux-386-debug linux-amd64-debug linux-arm64-debug linux-arm64-debug\
linux-loong64-debug linux-riscv64-debug


release:windows-386-release windows-amd64-release windows-arm-release windows-arm64-release\
linux-386-release linux-amd64-release windows-arm64-release linux-arm64-release\
linux-loong64-release linux-riscv64-release


# 误报病毒
windows-386-debug:
	-CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o debug/server_windows-386.exe -gcflags "-N -l" *.go
windows-386-release:
	-CGO_ENABLED=0 GOOS=windows GOARCH=386 go build -o release/server_windows-386.exe -ldflags="-s -w" *.go

windows-amd64-debug:
	-CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o debug/server_windows-amd64.exe -gcflags "-N -l" *.go
windows-amd64-release:
	-CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/server_windows-amd64.exe -ldflags="-s -w" *.go

windows-arm-debug:
	-CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -o debug/server_windows-arm.exe -gcflags "-N -l" *.go
windows-arm-release:
	-CGO_ENABLED=0 GOOS=windows GOARCH=arm go build -o release/server_windows-arm.exe -ldflags="-s -w" *.go

windows-arm64-debug:
	-CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o debug/server_windows-arm64.exe -gcflags "-N -l" *.go
windows-arm64-release:
	-CGO_ENABLED=0 GOOS=windows GOARCH=arm64 go build -o release/server_windows-arm64.exe -ldflags="-s -w" *.go

linux-386-debug:
	-CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o debug/server_linux-386 -gcflags "-N -l" *.go
linux-386-release:
	-CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -o release/server_linux-386 -ldflags="-s -w" *.go

linux-amd64-debug:
	-CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o debug/server_linux-amd64 -gcflags "-N -l" *.go
linux-amd64-release:
	-CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o release/server_linux-amd64 -ldflags="-s -w" *.go

linux-arm-debug:
	-CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o debug/server_linux-arm -gcflags "-N -l" *.go
linux-arm-release:
	-CGO_ENABLED=0 GOOS=linux GOARCH=arm go build -o release/server_linux-arm -ldflags="-s -w" *.go

linux-arm64-debug:
	-CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o debug/server_linux-arm64 -gcflags "-N -l" *.go
linux-arm64-release:
	-CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -o release/server_linux-arm64 -ldflags="-s -w" *.go

linux-loong64-debug:
	-CGO_ENABLED=0 GOOS=linux GOARCH=loong64 go build -o debug/server_linux-loong64 -gcflags "-N -l" *.go
linux-loong64-release:
	-CGO_ENABLED=0 GOOS=linux GOARCH=loong64 go build -o release/server_linux-loong64 -gcflags "-N -l" *.go

linux-riscv64-debug:
	-CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -o debug/server_linux-riscv64 -gcflags "-N -l" *.go
linux-riscv64-release:
	-CGO_ENABLED=0 GOOS=linux GOARCH=riscv64 go build -o release/server_linux-riscv64 -gcflags "-N -l" *.go

clean:
	-rm -rf debug/server*
	-rm -rf release/server*
config:
	-cp -r ../../config ../debug/config
	-cp -r ../../config ../release/config