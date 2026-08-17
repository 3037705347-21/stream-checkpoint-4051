# 修复前故障复现（Docker）

## 项目与标准命令

项目为 Go 检查点协调组件。标准验证命令为 `go test -count=20 ./...`。

## 环境构建与编译

在当前 Windows 平台执行 `go build ./...` 可以完成编译；随后在容器工作目录执行标准验证命令。

## 故障触发步骤

执行 `go test -count=20 ./...`。

## 实际错误输出

`error = open rejected: stream rejected, want errors.Is(_, stream must not be empty)`

## 期望行为

调用方应能使用公开错误值分类无效输入和检查点状态，检查点编号应按创建顺序连续增长。
