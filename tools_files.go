package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerFileTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("files_list",
		mcp.WithDescription("List files in the vault or a specific folder"),
		mcp.WithString("folder",
			mcp.Description("Folder path to list; if omitted, lists all files in the vault"),
		),
		vaultOpt(),
	), filesListHandler)

	s.AddTool(mcp.NewTool("folders_list",
		mcp.WithDescription("List all folders in the vault"),
		vaultOpt(),
	), foldersListHandler)
}

func filesListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)

	var args []string
	if v := req.GetString("folder", ""); v != "" {
		args = append(args, "folder="+v)
	}

	out, err := runObsidian(vault, "files", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func foldersListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "folders")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
