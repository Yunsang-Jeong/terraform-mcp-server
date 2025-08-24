package tools

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
)

// terraform module에 대한 정보를 획득하는 Tool을 만드려고함
// hashicorp의 public registry가 아니라, gitlab 혹은 github에 있는 모듈을 기반으로 보고 싶어
func GetModule(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	return nil, nil
}
