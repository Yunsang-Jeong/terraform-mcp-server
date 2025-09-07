package tools

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"

	"github.com/Yunsang-Jeong/terraform-config-parser/pkg/parser"
	"github.com/Yunsang-Jeong/terraform-config-parser/pkg/source"
	"github.com/mark3labs/mcp-go/mcp"
)

// GetModule retrieves terraform module information from Git repositories (GitLab/GitHub)
// instead of HashiCorp's public registry
func GetModule(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	// Extract module URL and optional parameters
	moduleURL, err := request.RequireString("url")
	if err != nil {
		return mcp.NewToolResultError("'url' parameter is required"), nil
	}

	// Optional parameters
	ref := request.GetString("ref", "")
	subDir := request.GetString("subdir", "")

	// Validate and normalize Git URL
	gitURL, err := normalizeGitURL(moduleURL)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Invalid Git URL: %v", err)), nil
	}

	// Create Git source
	config := source.SourceConfig{
		Ref:    ref,
		SubDir: subDir,
	}
	gitSource := source.NewGitSource(gitURL, config)

	// Fetch repository
	fs, rootPath, err := gitSource.Fetch()
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error fetching repository: %v", err)), nil
	}

	// Cleanup when done
	defer gitSource.Cleanup()

	// Parse Terraform configuration
	terraformParser := parser.NewParser(fs, parser.Simple)
	tfConfig, err := terraformParser.ParseTerraformWorkspace(rootPath)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error parsing Terraform configuration: %v", err)), nil
	}

	// Generate summary
	summary, err := tfConfig.Summary(true)
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error generating summary: %v", err)), nil
	}

	// Create response with module information
	moduleInfo := map[string]interface{}{
		"url":    gitURL,
		"ref":    ref,
		"subdir": subDir,
		"config": json.RawMessage(summary),
	}

	moduleInfoJSON, err := json.MarshalIndent(moduleInfo, "", "  ")
	if err != nil {
		return mcp.NewToolResultError(fmt.Sprintf("Error marshaling module info: %v", err)), nil
	}

	return mcp.NewToolResultText(string(moduleInfoJSON)), nil
}

// normalizeGitURL converts various Git URL formats to a standardized format
func normalizeGitURL(rawURL string) (string, error) {
	// Handle different URL formats
	if strings.HasPrefix(rawURL, "git@") {
		// SSH format: git@github.com:user/repo.git -> https://github.com/user/repo.git
		parts := strings.Split(rawURL, ":")
		if len(parts) != 2 {
			return "", fmt.Errorf("invalid SSH URL format")
		}
		host := strings.TrimPrefix(parts[0], "git@")
		path := parts[1]

		// Ensure .git suffix for SSH URLs
		if !strings.HasSuffix(path, ".git") {
			path = path + ".git"
		}

		return fmt.Sprintf("https://%s/%s", host, path), nil
	}

	// Parse as URL
	u, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("failed to parse URL: %w", err)
	}

	// Ensure HTTPS scheme
	if u.Scheme == "" || u.Scheme == "http" {
		u.Scheme = "https"
	}

	// Ensure .git suffix for proper Git operations
	if !strings.HasSuffix(u.Path, ".git") {
		u.Path = u.Path + ".git"
	}

	return u.String(), nil
}
