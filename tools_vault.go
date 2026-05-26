package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerVaultTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("vault_info",
		mcp.WithDescription("Get information about the current Obsidian vault"),
		vaultOpt(),
	), vaultInfoHandler)

	s.AddTool(mcp.NewTool("vaults_list",
		mcp.WithDescription("List all Obsidian vaults known to the desktop app"),
	), vaultsListHandler)
}

func vaultInfoHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "vault")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func vaultsListHandler(_ context.Context, _ mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian("", "vaults")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
