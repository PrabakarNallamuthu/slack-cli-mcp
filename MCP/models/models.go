package models

import (
	"context"
	"github.com/mark3labs/mcp-go/mcp"
)

type Tool struct {
	Definition mcp.Tool
	Handler    func(ctx context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error)
}

// Result represents the Result schema from the OpenAPI specification
type Result struct {
	Message string `json:"message,omitempty"`
	Permalink string `json:"permalink,omitempty"`
}

// SearchRequest represents the SearchRequest schema from the OpenAPI specification
type SearchRequest struct {
	Query string `json:"query"` // Search query
}
