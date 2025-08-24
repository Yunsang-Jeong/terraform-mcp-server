package server

import (
	"fmt"
	"terraform-mcp-server/pkg/tools"
	"terraform-mcp-server/version"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func Run(port uint16) {
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

	addr := fmt.Sprintf(":%d", port)
	if err := server.NewStreamableHTTPServer(s).Start(addr); err != nil {
		fmt.Printf("Server error: %v\n", err)
	}
}
