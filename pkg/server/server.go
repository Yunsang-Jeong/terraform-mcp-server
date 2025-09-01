package server

import (
	"context"
	"fmt"
	"os"

	"github.com/Yunsang-Jeong/terraform-mcp-server/pkg/tools"
	"github.com/Yunsang-Jeong/terraform-mcp-server/version"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
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
func RunHttp(port uint16) error {
	s := createMCPServer()
	addr := fmt.Sprintf(":%d", port)

	if err := server.NewStreamableHTTPServer(s).Start(addr); err != nil {
		return err
	}

	return nil
}

// RunStdio starts the MCP server over stdio
func RunStdio() error {
	s := createMCPServer()
	stdioServer := server.NewStdioServer(s)
	ctx := context.Background()

	if err := stdioServer.Listen(ctx, os.Stdin, os.Stdout); err != nil {
		return err
	}

	return nil
}
