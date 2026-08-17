# 修复前故障复现（Docker）

## 项目与标准命令

项目为 Go 检查点协调组件。标准验证命令为 `go test -count=20 ./...`。

## 环境构建与编译

在当前 Windows 平台执行 `go build ./...` 可以完成编译；随后在容器工作目录执行标准验证命令。

## 故障触发步骤

执行 `go test -count=20 ./...`。

## 实际错误输出

`panic: runtime error: invalid memory address or nil pointer dereference`

## 期望行为

新建协调器后应能立即打开检查点并开始记录分片进度。
