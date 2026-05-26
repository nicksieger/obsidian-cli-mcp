package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerNoteTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("note_read",
		mcp.WithDescription("Read the content of an Obsidian note"),
		mcp.WithString("file",
			mcp.Description("Note name (partial match, like a wikilink)"),
		),
		mcp.WithString("path",
			mcp.Description("Exact note path (e.g. folder/note.md); use instead of file for precision"),
		),
		vaultOpt(),
	), noteReadHandler)

	s.AddTool(mcp.NewTool("note_create",
		mcp.WithDescription("Create a new Obsidian note. Provide name or path (not both)."),
		mcp.WithString("name",
			mcp.Description("File name without folder path"),
		),
		mcp.WithString("path",
			mcp.Description("Full path including folders (e.g. Projects/Meeting.md)"),
		),
		mcp.WithString("content",
			mcp.Description("Initial content; use \\n for newlines, \\t for tabs"),
		),
		mcp.WithString("template",
			mcp.Description("Name of an Obsidian template to apply"),
		),
		mcp.WithBoolean("overwrite",
			mcp.Description("Overwrite the note if it already exists"),
		),
		mcp.WithBoolean("open",
			mcp.Description("Open the note in Obsidian after creating"),
		),
		vaultOpt(),
	), noteCreateHandler)

	s.AddTool(mcp.NewTool("note_append",
		mcp.WithDescription("Append content to an existing Obsidian note"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content to append; use \\n for newlines"),
		),
		mcp.WithBoolean("inline",
			mcp.Description("Append without inserting a preceding newline"),
		),
		vaultOpt(),
	), noteAppendHandler)

	s.AddTool(mcp.NewTool("note_prepend",
		mcp.WithDescription("Prepend content to an existing Obsidian note"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		mcp.WithString("content",
			mcp.Required(),
			mcp.Description("Content to prepend; use \\n for newlines"),
		),
		mcp.WithBoolean("inline",
			mcp.Description("Prepend without inserting a trailing newline"),
		),
		vaultOpt(),
	), notePrependHandler)

	s.AddTool(mcp.NewTool("note_move",
		mcp.WithDescription("Move an Obsidian note to a different folder"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path to move"),
		),
		mcp.WithString("to",
			mcp.Required(),
			mcp.Description("Destination folder path"),
		),
		vaultOpt(),
	), noteMoveHandler)

	s.AddTool(mcp.NewTool("note_rename",
		mcp.WithDescription("Rename an Obsidian note"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Current note name or path"),
		),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("New name for the note (without path)"),
		),
		vaultOpt(),
	), noteRenameHandler)

	s.AddTool(mcp.NewTool("note_delete",
		mcp.WithDescription("Delete an Obsidian note (moves to system trash)"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path to delete"),
		),
		vaultOpt(),
	), noteDeleteHandler)

	s.AddTool(mcp.NewTool("note_info",
		mcp.WithDescription("Get file metadata (size, dates, path) for an Obsidian note"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), noteInfoHandler)
}

func noteReadHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file := req.GetString("file", "")
	path := req.GetString("path", "")

	var args []string
	switch {
	case file != "":
		args = append(args, "file="+file)
	case path != "":
		args = append(args, "path="+path)
	}

	out, err := runObsidian(vault, "read", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func noteCreateHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)

	var args []string
	if v := req.GetString("name", ""); v != "" {
		args = append(args, "name="+v)
	}
	if v := req.GetString("path", ""); v != "" {
		args = append(args, "path="+v)
	}
	if v := req.GetString("content", ""); v != "" {
		args = append(args, "content="+v)
	}
	if v := req.GetString("template", ""); v != "" {
		args = append(args, "template="+v)
	}
	if req.GetBool("overwrite", false) {
		args = append(args, "overwrite")
	}
	if req.GetBool("open", false) {
		args = append(args, "open")
	}

	out, err := runObsidian(vault, "create", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func noteAppendHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}
	content, err := req.RequireString("content")
	if err != nil {
		return toolError(err.Error())
	}

	args := []string{"file=" + file, "content=" + content}
	if req.GetBool("inline", false) {
		args = append(args, "inline")
	}

	out, err := runObsidian(vault, "append", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func notePrependHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}
	content, err := req.RequireString("content")
	if err != nil {
		return toolError(err.Error())
	}

	args := []string{"file=" + file, "content=" + content}
	if req.GetBool("inline", false) {
		args = append(args, "inline")
	}

	out, err := runObsidian(vault, "prepend", args...)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func noteMoveHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}
	to, err := req.RequireString("to")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "move", "file="+file, "to="+to)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func noteRenameHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}
	name, err := req.RequireString("name")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "rename", "file="+file, "name="+name)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func noteDeleteHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "delete", "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func noteInfoHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "file", "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
