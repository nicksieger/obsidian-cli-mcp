package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/mark3labs/mcp-go/mcp"
)

// vaultOpt returns a reusable optional vault parameter for tool definitions.
func vaultOpt() mcp.ToolOption {
	return mcp.WithString("vault",
		mcp.Description("Vault name to use (overrides OBSIDIAN_VAULT env var)"),
	)
}

// vaultFromReq extracts the vault name from a request, falling back to OBSIDIAN_VAULT.
func vaultFromReq(req mcp.CallToolRequest) string {
	if v := req.GetString("vault", ""); v != "" {
		return v
	}
	return os.Getenv("OBSIDIAN_VAULT")
}

// runObsidian executes the obsidian CLI. vault is optional; if non-empty it is
// prepended as vault=<name> before the command. args may be key=value pairs or
// bare flags.
func runObsidian(vault, command string, args ...string) (string, error) {
	cmdArgs := make([]string, 0, len(args)+2)
	if vault != "" {
		cmdArgs = append(cmdArgs, "vault="+vault)
	}
	cmdArgs = append(cmdArgs, command)
	cmdArgs = append(cmdArgs, args...)

	cmd := exec.Command("obsidian", cmdArgs...)
	out, err := cmd.CombinedOutput()
	output := strings.TrimSpace(string(out))
	if err != nil {
		if output != "" {
			return "", fmt.Errorf("%s", output)
		}
		return "", fmt.Errorf("obsidian: %v", err)
	}
	return output, nil
}

func toolError(msg string) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultError(msg), nil
}

func toolOK(text string) (*mcp.CallToolResult, error) {
	return mcp.NewToolResultText(text), nil
}
