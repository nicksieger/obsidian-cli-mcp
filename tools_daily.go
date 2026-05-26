package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerDailyTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("daily_note_read",
		mcp.WithDescription("Read the content of today's daily note"),
		vaultOpt(),
	), dailyNoteReadHandler)

	s.AddTool(mcp.NewTool("daily_note_append",
		mcp.WithDescription("Append content to today's daily note"),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content to append; use \\n for newlines"),
		),
		vaultOpt(),
	), dailyNoteAppendHandler)

	s.AddTool(mcp.NewTool("daily_note_prepend",
		mcp.WithDescription("Prepend content to today's daily note"),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content to prepend; use \\n for newlines"),
		),
		vaultOpt(),
	), dailyNotePrependHandler)

	s.AddTool(mcp.NewTool("daily_note_path",
		mcp.WithDescription("Get the file path of today's daily note"),
		vaultOpt(),
	), dailyNotePathHandler)
}

func dailyNoteReadHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "daily:read")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func dailyNoteAppendHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	content, err := req.RequireString("content")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vaultFromReq(req), "daily:append", "content="+content)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func dailyNotePrependHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	content, err := req.RequireString("content")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vaultFromReq(req), "daily:prepend", "content="+content)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func dailyNotePathHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "daily:path")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
