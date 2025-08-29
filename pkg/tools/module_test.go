package tools

import (
	"context"
	"testing"

	"github.com/mark3labs/mcp-go/mcp"
)

func TestNormalizeGitURL(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		wantErr  bool
	}{
		{
			name:     "SSH URL with git@ prefix",
			input:    "git@github.com:user/repo.git",
			expected: "https://github.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "SSH URL without .git suffix",
			input:    "git@github.com:user/repo",
			expected: "https://github.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "HTTPS URL with .git suffix",
			input:    "https://github.com/user/repo.git",
			expected: "https://github.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "HTTPS URL without .git suffix",
			input:    "https://github.com/user/repo",
			expected: "https://github.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "HTTP URL converted to HTTPS",
			input:    "http://github.com/user/repo",
			expected: "https://github.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "URL without scheme",
			input:    "github.com/user/repo",
			expected: "https://github.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "GitLab URL",
			input:    "https://gitlab.com/user/repo.git",
			expected: "https://gitlab.com/user/repo.git",
			wantErr:  false,
		},
		{
			name:     "Invalid SSH format",
			input:    "git@github.com",
			expected: "",
			wantErr:  true,
		},
		{
			name:     "Invalid URL format",
			input:    "not-a-valid-url",
			expected: "https://not-a-valid-url.git",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := normalizeGitURL(tt.input)
			
			if tt.wantErr {
				if err == nil {
					t.Errorf("normalizeGitURL() expected error but got none")
				}
				return
			}
			
			if err != nil {
				t.Errorf("normalizeGitURL() unexpected error: %v", err)
				return
			}
			
			if result != tt.expected {
				t.Errorf("normalizeGitURL() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestGetModule_MissingURL(t *testing.T) {
	ctx := context.Background()
	
	// Mock request without URL parameter
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"branch": "main",
			},
		},
	}
	
	result, err := GetModule(ctx, request)
	
	if err != nil {
		t.Errorf("GetModule() unexpected error: %v", err)
	}
	
	if result == nil {
		t.Fatal("GetModule() returned nil result")
	}
	
	// Should return an error result for missing URL
	if result.IsError {
		t.Logf("GetModule() correctly returned error for missing URL: %+v", result)
	} else {
		t.Error("GetModule() should return error for missing URL parameter")
	}
}

func TestGetModule_InvalidURL(t *testing.T) {
	ctx := context.Background()
	
	// Mock request with invalid URL
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"url": "invalid-url-format",
			},
		},
	}
	
	result, err := GetModule(ctx, request)
	
	if err != nil {
		t.Errorf("GetModule() unexpected error: %v", err)
	}
	
	if result == nil {
		t.Fatal("GetModule() returned nil result")
	}
	
	// Should return an error result for repository fetching failure
	if result.IsError {
		t.Logf("GetModule() correctly returned error for invalid repository: %+v", result)
	} else {
		t.Error("GetModule() should return error for invalid repository")
	}
}

func TestGetModule_ValidParameters(t *testing.T) {
	ctx := context.Background()
	
	// Mock request with valid parameters
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"url":    "https://github.com/terraform-aws-modules/terraform-aws-vpc.git",
				"branch": "v5.0.0",
				"subdir": "",
			},
		},
	}
	
	result, err := GetModule(ctx, request)
	
	if err != nil {
		t.Errorf("GetModule() unexpected error: %v", err)
	}
	
	if result == nil {
		t.Fatal("GetModule() returned nil result")
	}
	
	// Note: This test might fail if network is unavailable or repository doesn't exist
	// In a production environment, you might want to mock the Git operations
	t.Logf("GetModule() result: %+v", result)
}

func TestGetModule_WithSubdirectory(t *testing.T) {
	ctx := context.Background()
	
	// Mock request with subdirectory parameter
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"url":    "https://github.com/user/mono-repo.git",
				"branch": "main",
				"subdir": "modules/vpc",
			},
		},
	}
	
	result, err := GetModule(ctx, request)
	
	if err != nil {
		t.Errorf("GetModule() unexpected error: %v", err)
	}
	
	if result == nil {
		t.Fatal("GetModule() returned nil result")
	}
	
	// This will likely fail due to non-existent repository, but tests parameter handling
	t.Logf("GetModule() with subdir result: %+v", result)
}

// Benchmark tests
func BenchmarkNormalizeGitURL(b *testing.B) {
	testURL := "git@github.com:user/repo"
	
	for i := 0; i < b.N; i++ {
		_, _ = normalizeGitURL(testURL)
	}
}

func BenchmarkGetModule_ParameterExtraction(b *testing.B) {
	ctx := context.Background()
	request := mcp.CallToolRequest{
		Params: mcp.CallToolParams{
			Arguments: map[string]interface{}{
				"url":    "invalid-repo-for-benchmark",
				"branch": "main",
				"subdir": "test",
			},
		},
	}
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = GetModule(ctx, request)
	}
}
