package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerSearchTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("vault_search",
		mcp.WithDescription("Search the Obsidian vault for notes matching a query"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search text"),
		),
		mcp.WithString("path",
			mcp.Description("Limit search to notes within this folder"),
		),
		mcp.WithInteger("limit",
			mcp.Description("Maximum number of results to return"),
		),
		mcp.WithBoolean("case_sensitive",
			mcp.Description("Use case-sensitive matching"),
		),
		vaultOpt(),
	), vaultSearchHandler)

	s.AddTool(mcp.NewTool("vault_search_context",
		mcp.WithDescription("Search the vault and return results with surrounding context snippets"),
		mcp.WithString("query",
			mcp.Required(),
			mcp.Description("Search text"),
		),
		mcp.WithInteger("limit",
			mcp.Description("Maximum number of results to return"),
		),
		vaultOpt(),
	), vaultSearchContextHandler)
}

func vaultSearchHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	query, err := req.RequireString("query")
	if err != nil {
		return toolError(err.Error())
	}

	args := []string{"query=" + query}
	if v := req.GetString("path", ""); v != "" {
		args = append(args, "path="+v)
	}
	if v := req.GetInt("limit", 0); v > 0 {
		args = append(args, "limit="+itoa(v))
	}
	if req.GetBool("case_sensitive", false) {
		args = append(args, "case")
	}

	out, err := runObsidian(vault, "search", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func vaultSearchContextHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	query, err := req.RequireString("query")
	if err != nil {
		return toolError(err.Error())
	}

	args := []string{"query=" + query}
	if v := req.GetInt("limit", 0); v > 0 {
		args = append(args, "limit="+itoa(v))
	}

	out, err := runObsidian(vault, "search:context", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
