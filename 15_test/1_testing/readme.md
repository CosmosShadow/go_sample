### 测试

```shell
# 单元测试
go test -v .

# 测试覆盖率
go test -v -coverprofile=cover.out

# 生成html文件
go tool cover -html=cover.out -o cover.html

# 性能测试
go test -v -bench="BenchmarkGenShortID$" --run=none

# 性能测试: 查看CPU性能
go test -v -bench="BenchmarkGenShortID$" --run=none -cpuprofile cpu.out

# 通过web查看CPU详细性能
go tool pprof -http=":" cpu.out
```

