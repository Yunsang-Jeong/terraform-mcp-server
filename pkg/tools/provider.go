package tools

import (
	"context"
	"terraform-mcp-server/pkg/utils/registry"

	"github.com/mark3labs/mcp-go/mcp"
)

func getBlockDocument(providerNamespace, providerName, providerVersion, blockType, blockName string) (string, error) {
	var docId string

	if providerVersion == "" {
		provider, err := registry.GetProvider(providerNamespace, providerName)
		if err != nil {
			return "", err
		}

		for _, doc := range provider.Docs {
			if doc.Language == "hcl" && doc.Category == blockType && doc.Slug == blockName {
				docId = doc.ID
				break
			}
		}
	} else {
		versionId, err := registry.GetProviderVersionId(providerNamespace, providerName, providerVersion)
		if err != nil {
			return "", err
		}

		docId, err = registry.GetProviderDocsId(versionId, blockType, blockName)
		if err != nil {
			return "", err
		}
	}

	contents, err := registry.GetProviderDocsContent(docId)
	if err != nil {
		return "", err
	}

	return contents, nil
}

func GetResourceBlockDocument(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	providerNamespace := request.GetString("provider_namespace", "hashicorp")

	providerName, err := request.RequireString("provider_name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	providerVersion := request.GetString("provider_version", "")

	blockType := "resources"

	blockName, err := request.RequireString("block_name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	contents, err := getBlockDocument(providerNamespace, providerName, providerVersion, blockType, blockName)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(contents), nil
}

func GetDataBlockDocument(ctx context.Context, request mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	providerNamespace := request.GetString("provider_namespace", "hashicorp")

	providerName, err := request.RequireString("provider_name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	providerVersion := request.GetString("provider_version", "")

	blockType := "data-sources"

	blockName, err := request.RequireString("block_name")
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	contents, err := getBlockDocument(providerNamespace, providerName, providerVersion, blockType, blockName)
	if err != nil {
		return mcp.NewToolResultError(err.Error()), nil
	}

	return mcp.NewToolResultText(contents), nil
}
