package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerTaskTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("tasks_list",
		mcp.WithDescription("List tasks in the vault or a specific note"),
		mcp.WithString("status",
			mcp.Description("Filter by status: \"todo\" (incomplete), \"done\" (completed), or omit for all"),
		),
		mcp.WithString("file",
			mcp.Description("Note name or path; if omitted, lists tasks across the whole vault"),
		),
		vaultOpt(),
	), tasksListHandler)

	s.AddTool(mcp.NewTool("task_toggle",
		mcp.WithDescription("Toggle the completion status of a task at a specific line"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path containing the task"),
		),
		mcp.WithInteger("line",
			mcp.Required(),
			mcp.Description("Line number of the task to toggle (1-based)"),
		),
		vaultOpt(),
	), taskToggleHandler)
}

func tasksListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)

	var args []string
	if v := req.GetString("file", ""); v != "" {
		args = append(args, "file="+v)
	}

	status := req.GetString("status", "")
	switch status {
	case "todo", "done":
		args = append(args, status)
	}

	out, err := runObsidian(vault, "tasks", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func taskToggleHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}
	line, err := req.RequireInt("line")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "task", "file="+file, "line="+itoa(line), "toggle")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
