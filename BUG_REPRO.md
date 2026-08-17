# 修复前故障复现（Docker）

## 项目与标准命令

项目为 Go 检查点协调组件。标准验证命令为 `go test -count=20 ./...`。

## 环境构建与编译

在当前 Windows 平台执行 `go build ./...` 可以完成编译；随后在容器工作目录执行标准验证命令。

## 故障触发步骤

执行 `go test -count=20 ./...`。

## 实际错误输出

`markers = []checkpoint.Marker{checkpoint.Marker{Shard:"edge-a", Sequence:9}, checkpoint.Marker{Shard:"edge-b", Sequence:13}}, want []checkpoint.Marker{checkpoint.Marker{Shard:"edge-a", Sequence:7}, checkpoint.Marker{Shard:"edge-b", Sequence:11}}`

## 期望行为

提交批次中的每个分片标记应与确认时的序号一致，并以稳定顺序交付。
