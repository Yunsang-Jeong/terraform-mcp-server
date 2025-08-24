package tools_test

import (
	"context"
	"testing"

	"terraform-mcp-server/pkg/tools"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestGetResourceBlockDocument_Success(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_resource_block_document"
	request.Params.Arguments = map[string]any{
		"provider_namespace": "hashicorp",
		"provider_name":      "aws",
		"block_name":         "s3_bucket",
	}
	result, err := tools.GetResourceBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.Content == nil {
		t.Fatal("Expected content, got nil")
	}

	if len(result.Content) == 0 {
		t.Error("Expected non-empty content")
	}
}

func TestGetDataBlockDocument_Success(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_data_block_document"
	request.Params.Arguments = map[string]any{
		"provider_namespace": "hashicorp",
		"provider_name":      "aws",
		"block_name":         "s3_bucket",
	}
	result, err := tools.GetDataBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.Content == nil {
		t.Fatal("Expected content, got nil")
	}

	if len(result.Content) == 0 {
		t.Error("Expected non-empty content")
	}
}

func TestGetResourceBlockDocument_WithCustomNamespace(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_resource_block_document"
	request.Params.Arguments = map[string]any{
		"provider_namespace": "custom",
		"provider_name":      "aws",
		"block_name":         "s3_bucket",
	}
	result, err := tools.GetResourceBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}
}

func TestGetResourceBlockDocument_WithVersion(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_resource_block_document"
	request.Params.Arguments = map[string]any{
		"provider_name":    "aws",
		"provider_version": "5.0.0",
		"block_name":       "s3_bucket",
	}
	result, err := tools.GetResourceBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.Content == nil {
		t.Fatal("Expected content, got nil")
	}

	if len(result.Content) == 0 {
		t.Error("Expected non-empty content")
	}
}

func TestGetResourceBlockDocument_MissingProviderName(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_resource_block_document"
	request.Params.Arguments = map[string]any{
		"block_name": "s3_bucket",
	}
	result, err := tools.GetResourceBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error from function, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if !result.IsError {
		t.Error("Expected error result for missing provider_name")
	}
}

func TestGetResourceBlockDocument_MissingBlockName(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_resource_block_document"
	request.Params.Arguments = map[string]any{
		"provider_name": "aws",
	}
	result, err := tools.GetResourceBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error from function, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if !result.IsError {
		t.Error("Expected error result for missing block_name")
	}
}

func TestGetDataBlockDocument_WithCustomNamespace(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_data_block_document"
	request.Params.Arguments = map[string]any{
		"provider_namespace": "custom",
		"provider_name":      "aws",
		"block_name":         "s3_bucket",
	}
	result, err := tools.GetDataBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}
}

func TestGetDataBlockDocument_WithVersion(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_data_block_document"
	request.Params.Arguments = map[string]any{
		"provider_name":    "aws",
		"provider_version": "5.0.0",
		"block_name":       "s3_bucket",
	}
	result, err := tools.GetDataBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if result.Content == nil {
		t.Fatal("Expected content, got nil")
	}

	if len(result.Content) == 0 {
		t.Error("Expected non-empty content")
	}
}

func TestGetDataBlockDocument_MissingProviderName(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_data_block_document"
	request.Params.Arguments = map[string]any{
		"block_name": "s3_bucket",
	}
	result, err := tools.GetDataBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error from function, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if !result.IsError {
		t.Error("Expected error result for missing provider_name")
	}
}

func TestGetDataBlockDocument_MissingBlockName(t *testing.T) {
	ctx := context.Background()

	request := mcp.CallToolRequest{}
	request.Params.Name = "search_data_block_document"
	request.Params.Arguments = map[string]any{
		"provider_name": "aws",
	}
	result, err := tools.GetDataBlockDocument(ctx, request)

	if err != nil {
		t.Fatalf("Expected no error from function, got %v", err)
	}

	if result == nil {
		t.Fatal("Expected result, got nil")
	}

	if !result.IsError {
		t.Error("Expected error result for missing block_name")
	}
}
