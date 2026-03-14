# ast-analysis

Command-line app to demonstrate astkratos package capabilities with Kratos demo projects.

---

## CHINESE README

[中文说明](README.zh.md)

## Overview

**ast-analysis** is a command-line app that demonstrates the [kratos-ast](https://github.com/yylego/kratos-ast) package capabilities. It performs comprehensive code structure analysis on sibling demo1kratos and demo2kratos projects, including gRPC component detection, struct extraction, and module metadata parsing.

## Features

- **gRPC Detection** - Scans proto-generated files to detect clients, servers, and services
- **Struct Extraction** - Parses Go source with AST to extract struct definitions
- **Module Analysis** - Extracts go.mod metadata including dependencies and toolchain info
- **Project Scanning** - Comprehensive project analysis combining all detection capabilities

## Usage

### Build and Run

```bash
go build -o ast-analysis .
./ast-analysis
```

The app analyzes both demo1kratos and demo2kratos projects, prints detailed results to console, and generates report files.

### Generated Reports

The app generates two markdown report files in the current DIR:

- **demo1-report.md** - Analysis report for demo1kratos (Student service)
- **demo2-report.md** - Analysis report for demo2kratos (Article service)

Each report contains:

- **gRPC Clients** - Detected client interfaces (e.g., StudentServiceClient, ArticleServiceClient)
- **gRPC Servers** - Detected server interfaces (e.g., StudentServiceServer, ArticleServiceServer)
- **gRPC Services** - Detected service names (e.g., StudentService, ArticleService)
- **Struct Definitions** - Parsed proto message structs (e.g., CreateStudentRequest, ArticleInfo)
- **Module Info** - Module path, Go version, toolchain, dependencies
- **Project Report** - Full aggregated analysis in JSON format

## Dependencies

- [github.com/yylego/kratos-ast](https://github.com/yylego/kratos-ast) - Kratos project AST analysis engine
- [github.com/yylego/runpath](https://github.com/yylego/runpath) - Runtime path operations
- [github.com/yylego/zaplog](https://github.com/yylego/zaplog) - Logging support

## About

This is a subproject of [ast](https://github.com/kratos-examples/ast), demonstrating how to use the astkratos package to analyze Kratos project code structures.

See the main project docs to get more info.
