# Terraform MCP Server

A Model Context Protocol (MCP) server that provides tools for searching and retrieving Terraform stuff.

## Available Tools

### `search_resource_block_document`

특정 버전의 resource block 설명을 가져옵니다.

**Parameters:**
- `provider_name` (required): Provider name (e.g., 'aws', 'azurerm')
- `block_name` (required): Block name to search for (e.g., 's3_bucket')
- `provider_namespace` (optional): Provider namespace (default: 'hashicorp')
- `provider_version` (optional): Provider version (leave empty for latest)

### `search_data_block_document`

특정 버전의 data block 설명을 가져옵니다.

**Parameters:**
- `provider_name` (required): Provider name (e.g., 'aws', 'azurerm')
- `block_name` (required): Block name to search for (e.g., 's3_bucket')
- `provider_namespace` (optional): Provider namespace (default: 'hashicorp')
- `provider_version` (optional): Provider version (leave empty for latest)
