package main

import (
	"context"

	"github.com/mark3labs/mcp-go/mcp"
	"github.com/mark3labs/mcp-go/server"
)

func registerPropertyTools(s *server.MCPServer) {
	s.AddTool(mcp.NewTool("property_read",
		mcp.WithDescription("Read the value of a frontmatter property from a note"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Property key name"),
		),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), propertyReadHandler)

	s.AddTool(mcp.NewTool("property_set",
		mcp.WithDescription("Set a frontmatter property value on a note"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Property key name"),
		),
		mcp.WithString("value",
			mcp.Required(),
			mcp.Description("Value to set"),
		),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), propertySetHandler)

	s.AddTool(mcp.NewTool("property_remove",
		mcp.WithDescription("Remove a frontmatter property from a note"),
		mcp.WithString("name",
			mcp.Required(),
			mcp.Description("Property key name to remove"),
		),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), propertyRemoveHandler)

	s.AddTool(mcp.NewTool("properties_list",
		mcp.WithDescription("List all frontmatter properties of a note"),
		mcp.WithString("file",
			mcp.Required(),
			mcp.Description("Note name or path"),
		),
		vaultOpt(),
	), propertiesListHandler)
}

func propertyReadHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	name, err := req.RequireString("name")
	if err != nil {
		return toolError(err.Error())
	}
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "property:read", "name="+name, "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func propertySetHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	name, err := req.RequireString("name")
	if err != nil {
		return toolError(err.Error())
	}
	value, err := req.RequireString("value")
	if err != nil {
		return toolError(err.Error())
	}
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "property:set", "name="+name, "value="+value, "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func propertyRemoveHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	name, err := req.RequireString("name")
	if err != nil {
		return toolError(err.Error())
	}
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "property:remove", "name="+name, "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}

func propertiesListHandler(_ context.Context, req mcp.CallToolRequest) (*mcp.CallToolResult, error) {
	vault := vaultFromReq(req)
	file, err := req.RequireString("file")
	if err != nil {
		return toolError(err.Error())
	}

	out, err := runObsidian(vault, "properties", "file="+file)
	if err != nil {
		return toolError(err.Error())
	}
	return toolOK(out)
}
