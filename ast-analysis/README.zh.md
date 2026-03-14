# ast-analysis

用于演示 astkratos 包功能的命令行应用，配合 Kratos 演示项目使用。

---

## 英文文档

[ENGLISH README](README.md)

## 项目简介

**ast-analysis** 是一个命令行应用，用于演示 [kratos-ast](https://github.com/yylego/kratos-ast) 包的功能。它对同级的 demo1kratos 和 demo2kratos 项目执行全面的代码结构分析，包括 gRPC 组件检测、结构体提取和模块元数据解析。

## 功能特性

- **gRPC 检测** - 扫描 proto 生成的文件以检测客户端、服务端和服务
- **结构体提取** - 使用 AST 解析 Go 源码并提取结构体定义
- **模块分析** - 提取 go.mod 元数据，包括依赖和工具链信息
- **项目扫描** - 综合所有检测功能的全面项目分析

## 使用方法

### 构建和运行

```bash
go build -o ast-analysis .
./ast-analysis
```

应用会分析 demo1kratos 和 demo2kratos 两个项目，将详细结果打印到控制台，并生成报告文件。

### 生成的报告

应用在当前目录生成两个 markdown 报告文件：

- **demo1-report.md** - demo1kratos 的分析报告（Student 服务）
- **demo2-report.md** - demo2kratos 的分析报告（Article 服务）

每个报告包含：

- **gRPC 客户端** - 检测到的客户端接口（如 StudentServiceClient、ArticleServiceClient）
- **gRPC 服务端** - 检测到的服务端接口（如 StudentServiceServer、ArticleServiceServer）
- **gRPC 服务** - 检测到的服务名称（如 StudentService、ArticleService）
- **结构体定义** - 解析的 proto 消息结构体（如 CreateStudentRequest、ArticleInfo）
- **模块信息** - 模块路径、Go 版本、工具链、依赖
- **项目报告** - JSON 格式的完整聚合分析结果

## 依赖项

- [github.com/yylego/kratos-ast](https://github.com/yylego/kratos-ast) - Kratos 项目 AST 分析引擎
- [github.com/yylego/runpath](https://github.com/yylego/runpath) - 运行时路径操作
- [github.com/yylego/zaplog](https://github.com/yylego/zaplog) - 日志支持

## 项目说明

本项目是 [ast](https://github.com/kratos-examples/ast) 的子项目，用于演示如何使用 astkratos 包分析 Kratos 项目代码结构。

更多信息请参考主项目文档。
