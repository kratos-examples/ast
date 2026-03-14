// ast-analysis demonstrates astkratos package capabilities with Kratos demo projects
// Performs comprehensive gRPC component detection, struct extraction, and module metadata parsing
// Uses runpath to locate sibling demo1kratos and demo2kratos projects
// Generates analysis reports as markdown files for each demo project
//
// ast-analysis 演示 astkratos 包在 Kratos 演示项目中的功能
// 执行全面的 gRPC 组件检测、结构体提取和模块元数据解析
// 使用 runpath 定位同级的 demo1kratos 和 demo2kratos 项目
// 为每个演示项目生成 markdown 格式的分析报告
package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/yylego/kratos-ast/astkratos"
	"github.com/yylego/must"
	"github.com/yylego/neatjson/neatjsons"
	"github.com/yylego/osexistpath/osmustexist"
	"github.com/yylego/runpath"
	"github.com/yylego/zaplog"
)

const separatorWidth = 80 // Dividing mark width // 分隔标记宽度

// main performs comprehensive astkratos analysis on both demo projects
// main 对两个演示项目执行全面的 astkratos 分析
func main() {
	astkratos.SetDebugMode(true)

	// Locate demo project roots using runpath
	// 使用 runpath 定位演示项目根目录
	demo1Root := osmustexist.ROOT(runpath.PARENT.UpTo(1, "demo1kratos"))
	demo2Root := osmustexist.ROOT(runpath.PARENT.UpTo(1, "demo2kratos"))

	analyzeDemo("demo1kratos", demo1Root, filepath.Join(demo1Root, "api/student/student.pb.go"), runpath.PARENT.Join("demo1-report.md"))
	analyzeDemo("demo2kratos", demo2Root, filepath.Join(demo2Root, "api/article/article.pb.go"), runpath.PARENT.Join("demo2-report.md"))
}

// analyzeDemo performs full astkratos analysis on a single demo project and writes report
// analyzeDemo 对单个演示项目执行完整的 astkratos 分析并写入报告
func analyzeDemo(name string, projectRoot string, protoGoPath string, reportPath string) {
	fmt.Println(strings.Repeat("=", separatorWidth))
	fmt.Printf("Analyzing: %s\n", name)
	fmt.Println(strings.Repeat("=", separatorWidth))

	apiPath := osmustexist.ROOT(filepath.Join(projectRoot, "api"))

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf("# %s Analysis Report\n\n", name))

	// gRPC client detection
	// gRPC 客户端检测
	clients := astkratos.ListGrpcClients(apiPath)
	zaplog.SUG.Debugln("gRPC Clients:", neatjsons.S(clients))
	must.True(len(clients) > 0)
	sb.WriteString("## gRPC Clients\n\n")
	for _, c := range clients {
		sb.WriteString(fmt.Sprintf("- **%s** (package: %s)\n", c.Name, c.Package))
	}
	sb.WriteString("\n")

	// gRPC server detection
	// gRPC 服务端检测
	servers := astkratos.ListGrpcServers(apiPath)
	zaplog.SUG.Debugln("gRPC Servers:", neatjsons.S(servers))
	must.True(len(servers) > 0)
	sb.WriteString("## gRPC Servers\n\n")
	for _, s := range servers {
		sb.WriteString(fmt.Sprintf("- **%s** (package: %s)\n", s.Name, s.Package))
	}
	sb.WriteString("\n")

	// gRPC service detection
	// gRPC 服务检测
	services := astkratos.ListGrpcServices(apiPath)
	zaplog.SUG.Debugln("gRPC Services:", neatjsons.S(services))
	must.True(len(services) > 0)
	sb.WriteString("## gRPC Services\n\n")
	for _, s := range services {
		sb.WriteString(fmt.Sprintf("- **%s** (package: %s)\n", s.Name, s.Package))
	}
	sb.WriteString("\n")

	// Boolean check functions
	// 布尔检查函数
	hasClients := astkratos.HasGrpcClients(apiPath)
	hasServers := astkratos.HasGrpcServers(apiPath)
	serviceCount := astkratos.CountGrpcServices(apiPath)
	zaplog.SUG.Debugln("Has gRPC Clients:", hasClients)
	zaplog.SUG.Debugln("Has gRPC Servers:", hasServers)
	zaplog.SUG.Debugln("gRPC Service Count:", serviceCount)
	sb.WriteString("## Statistics\n\n")
	sb.WriteString(fmt.Sprintf("- Has gRPC Clients: %v\n", hasClients))
	sb.WriteString(fmt.Sprintf("- Has gRPC Servers: %v\n", hasServers))
	sb.WriteString(fmt.Sprintf("- gRPC Service Count: %d\n", serviceCount))
	sb.WriteString("\n")

	// Struct parsing from proto-generated files
	// 从 proto 生成文件中解析结构体
	structMap := astkratos.ParseStructs(protoGoPath)
	zaplog.SUG.Debugln("Struct count:", len(structMap))
	sb.WriteString("## Parsed Structs\n\n")
	for structName, definition := range structMap {
		zaplog.SUG.Debugln("Struct:", structName, "->", definition.Name)
		sb.WriteString(fmt.Sprintf("- **%s**\n", structName))
	}
	sb.WriteString("\n")

	// Module metadata extraction
	// 模块元数据提取
	moduleInfo, err := astkratos.GetModuleInfo(projectRoot)
	must.Done(err)
	zaplog.SUG.Debugln("Module Path:", moduleInfo.Module.Path)
	zaplog.SUG.Debugln("Go Version:", moduleInfo.Go)
	zaplog.SUG.Debugln("Toolchain Version:", moduleInfo.GetToolchainVersion())
	zaplog.SUG.Debugln("Dependencies count:", len(moduleInfo.Require))
	sb.WriteString("## Module Info\n\n")
	sb.WriteString(fmt.Sprintf("- Module Path: %s\n", moduleInfo.Module.Path))
	sb.WriteString(fmt.Sprintf("- Go Version: %s\n", moduleInfo.Go))
	sb.WriteString(fmt.Sprintf("- Toolchain Version: %s\n", moduleInfo.GetToolchainVersion()))
	sb.WriteString(fmt.Sprintf("- Dependencies: %d\n", len(moduleInfo.Require)))
	sb.WriteString("\n")

	// Comprehensive project analysis
	// 全面项目分析
	report := astkratos.AnalyzeProject(projectRoot)
	zaplog.SUG.Debugln("Project Report:", neatjsons.S(report))
	sb.WriteString("## Full Project Report\n\n")
	sb.WriteString("```json\n")
	sb.WriteString(neatjsons.S(report))
	sb.WriteString("\n```\n")

	// Write report to file
	// 写入报告文件
	must.Done(os.WriteFile(reportPath, []byte(sb.String()), 0644))
	zaplog.SUG.Debugln("Report written to:", reportPath)

	fmt.Println(strings.Repeat("-", separatorWidth))
	fmt.Printf("Analysis complete: %s -> %s\n\n", name, reportPath)
}
