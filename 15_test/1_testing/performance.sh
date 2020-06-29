# 性能测试

go test -v -bench="BenchmarkGenShortID$" --run=none ./utils/

# 查看CPU性能
go test -v -bench="BenchmarkGenShortID$" --run=none -cpuprofile cpu.out ./utils/



# 安装库: apt install graphviz
# 通过web查看CPU详细性能
go tool pprof -http=":" cpu.out