package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTagTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("tags_list",
		mcp.WithDescription("List all tags in the vault, or tags in a specific note"),
		mcp.WithString("file",
			mcp.Description("Note name or path; if omitted, lists tags across the whole vault"),
		),
		vaultOpt(),
	), tagsListHandler)

	s.AddTool(mcp.NewTool("tag_files",
		mcp.WithDescription("List all files that have a specific tag"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Tag name (without the # prefix)"),
		),
		mcp.WithBoolean("verbose",
			mcp.Description("Include additional file details in output"),
		),
		vaultOpt(),
	), tagFilesHandler)
}

func tagsListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)

	var args []string
	if v := req.GetString("file", ""); v != "" {
		args = append(args, "file="+v)
	}

	out, err := runObsidian(vault, "tags", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func tagFilesHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	name, err := req.RequireString("name")
	if err != nil {
		return toolError(err.Error())
	}

	args := []string{"name=" + name}
	if req.GetBool("verbose", false) {
		args = append(args, "verbose")
	}

	out, err := runObsidian(vault, "tag", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
