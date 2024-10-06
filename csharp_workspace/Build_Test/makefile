# 不支持
loongarch64:
	dotnet publish Build_Test/Build_Test.csproj --os linux-loongarch64 -p:PublishSingleFile=true -o release/linux_loongarch64 &\
	dotnet publish Build_Test/Build_Test.csproj --os linux-riscv64 -p:PublishSingleFile=true -o release/linux_riscv64

clean:
	-rm -rf release