package server

import (
	"context"
	"fmt"
	"os"

	"github.com/Yunsang-Jeong/terraform-mcp-server/pkg/tools"
	"github.com/Yunsang-Jeong/terraform-mcp-server/version"

	"github.com/charmbracelet/lipgloss/v2"
	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

// Styles for pretty output
var (
	titleStyle = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("4"))   // Blue
	modeStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("6"))              // Cyan
	keyStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("8"))              // Gray
	valueStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("2"))              // Green
	protoStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("3"))              // Yellow
	urlStyle   = lipgloss.NewStyle().Foreground(lipgloss.Color("5"))              // Magenta
	toolStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("7")).Faint(true) // Dim
	okStyle    = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("2"))   // Green bold
	errStyle   = lipgloss.NewStyle().Bold(true).Foreground(lipgloss.Color("1"))   // Red bold
)

// createMCPServer creates and configures the MCP server with all tools
func createMCPServer() *server.MCPServer {
	s := server.NewMCPServer(
		"Terraform MCP Server",
		version.Version,
		server.WithToolCapabilities(false),
	)

	s.AddTool(mcp.NewTool("search_resource_block_document",
		mcp.WithDescription("특정 버전의 resource block 설명을 가져옵니다."),
		mcp.WithString("provider_namespace",
			mcp.Description("provider의 namespace 입니다. 기본값은 'hashicorp' 입니다."),
		),
		mcp.WithString("provider_name",
			mcp.Description("provider의 name 입니다. 예: 'aws', 'azurerm'."),
			mcp.Required(),
		),
		mcp.WithString("provider_version",
			mcp.Description("provider의 version 입니다. 최신버전은 생략하거나 공백을 입력합니다."),
		),
		mcp.WithString("block_name",
			mcp.Description("확인하려는 block의 name 입니다. 예: 's3_bucket'."),
			mcp.Required(),
		),
	), tools.GetResourceBlockDocument)

	s.AddTool(mcp.NewTool("search_data_block_document",
		mcp.WithDescription("특정 버전의 data block 설명을 가져옵니다."),
		mcp.WithString("provider_namespace",
			mcp.Description("provider의 namespace 입니다. 기본값은 'hashicorp' 입니다."),
		),
		mcp.WithString("provider_name",
			mcp.Description("provider의 name 입니다. 예: 'aws', 'azurerm'."),
			mcp.Required(),
		),
		mcp.WithString("provider_version",
			mcp.Description("provider의 version 입니다. 최신버전은 생략하거나 공백을 입력합니다."),
		),
		mcp.WithString("block_name",
			mcp.Description("확인하려는 block의 name 입니다. 예: 's3_bucket'."),
			mcp.Required(),
		),
	), tools.GetDataBlockDocument)

	s.AddTool(mcp.NewTool("get_module",
		mcp.WithDescription("Git 저장소(GitLab/GitHub)에서 Terraform 모듈 정보를 가져옵니다."),
		mcp.WithString("url",
			mcp.Description("Git 저장소 URL입니다. SSH 또는 HTTPS 형식 모두 지원합니다."),
			mcp.Required(),
		),
		mcp.WithString("branch",
			mcp.Description("사용할 브랜치 또는 태그입니다. 기본값은 main/master 브랜치입니다."),
		),
		mcp.WithString("subdir",
			mcp.Description("저장소 내 하위 디렉토리 경로입니다. 루트가 아닌 위치에 모듈이 있는 경우 사용합니다."),
		),
	), tools.GetModule)

	return s
}

// RunHttp starts the MCP server over HTTP
func RunHttp(port uint16) {
	s := createMCPServer()
	addr := fmt.Sprintf(":%d", port)

	fmt.Printf("Starting %s (%s)\n", 
		titleStyle.Render("Terraform MCP Server"), 
		modeStyle.Render("HTTP mode"))
	
	fmt.Printf("%s: %s\n", 
		keyStyle.Render("Server"), 
		valueStyle.Render(fmt.Sprintf("Terraform MCP Server v%s", version.Version)))
	
	fmt.Printf("%s: %s\n", 
		keyStyle.Render("Protocol"), 
		protoStyle.Render("HTTP"))
	
	fmt.Printf("%s: %s\n", 
		keyStyle.Render("Listening on"), 
		urlStyle.Render(fmt.Sprintf("http://localhost%s", addr)))
	
	fmt.Printf("%s: %s\n", 
		keyStyle.Render("Available tools"), 
		valueStyle.Render("3 tools"))
	
	fmt.Printf("  %s\n", toolStyle.Render("search_resource_block_document"))
	fmt.Printf("  %s\n", toolStyle.Render("search_data_block_document"))
	fmt.Printf("  %s\n", toolStyle.Render("get_module"))
	
	fmt.Printf("%s\n\n", okStyle.Render("Server ready to accept connections!"))

	if err := server.NewStreamableHTTPServer(s).Start(addr); err != nil {
		fmt.Printf("%s: %s\n", errStyle.Render("Server error"), errStyle.Render(err.Error()))
	}
}

// RunStdio starts the MCP server over stdio
func RunStdio() {
	s := createMCPServer()
	stdioServer := server.NewStdioServer(s)
	ctx := context.Background()

	// stderr로 로깅 (stdout은 MCP 통신용)
	fmt.Fprintf(os.Stderr, "Starting %s (%s)\n", 
		titleStyle.Render("Terraform MCP Server"), 
		modeStyle.Render("stdio mode"))
	
	fmt.Fprintf(os.Stderr, "%s: %s\n", 
		keyStyle.Render("Server"), 
		valueStyle.Render(fmt.Sprintf("Terraform MCP Server v%s", version.Version)))
	
	fmt.Fprintf(os.Stderr, "%s: %s\n", 
		keyStyle.Render("Protocol"), 
		protoStyle.Render("JSON-RPC over stdio"))
	
	fmt.Fprintf(os.Stderr, "%s: %s\n", 
		keyStyle.Render("Communication"), 
		urlStyle.Render("stdin/stdout"))
	
	fmt.Fprintf(os.Stderr, "%s: %s\n", 
		keyStyle.Render("Available tools"), 
		valueStyle.Render("3 tools"))
	
	fmt.Fprintf(os.Stderr, "  %s\n", toolStyle.Render("search_resource_block_document"))
	fmt.Fprintf(os.Stderr, "  %s\n", toolStyle.Render("search_data_block_document"))
	fmt.Fprintf(os.Stderr, "  %s\n", toolStyle.Render("get_module"))
	
	fmt.Fprintf(os.Stderr, "%s\n\n", okStyle.Render("Server ready to accept JSON-RPC messages!"))

	if err := stdioServer.Listen(ctx, os.Stdin, os.Stdout); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s\n", errStyle.Render("Server error"), errStyle.Render(err.Error()))
	}
}
