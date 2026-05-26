package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerLinkTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("links_list",
		mcp.WithDescription("List outgoing links from a note"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), linksListHandler)

	s.AddTool(mcp.NewTool("backlinks_list",
		mcp.WithDescription("List notes that link to a given note (incoming links)"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), backlinksListHandler)

	s.AddTool(mcp.NewTool("orphans_list",
		mcp.WithDescription("List notes that have no incoming links from other notes"),
		vaultOpt(),
	), orphansListHandler)

	s.AddTool(mcp.NewTool("deadends_list",
		mcp.WithDescription("List notes that have no outgoing links to other notes"),
		vaultOpt(),
	), deadendsListHandler)

	s.AddTool(mcp.NewTool("unresolved_list",
		mcp.WithDescription("List broken links (wikilinks that point to non-existent notes)"),
		vaultOpt(),
	), unresolvedListHandler)
}

func linksListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "links", "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func backlinksListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "backlinks", "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func orphansListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "orphans")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func deadendsListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "deadends")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func unresolvedListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	out, err := runObsidian(vaultFromReq(req), "unresolved")
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
