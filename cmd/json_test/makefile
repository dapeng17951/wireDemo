build:riscv_build loongarch_build

# riscv
riscv_build:
	-GOOS=linux GOARCH=riscv64 go build -o main_riscv64 main.go

# 龙芯
loongarch_build:
	-GOOS=linux GOARCH=loong64 go build -o main_loong64 main.go


# .NET 芯片的支持
# https://github.com/dotnet/runtime/tree/main/src/coreclr/vm

clean:
	-rm -rf main_riscv64 main_loong64