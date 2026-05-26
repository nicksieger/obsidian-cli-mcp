package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/mark3labs/mcp-go/server"
)

func main() {
	s := server.NewMCPServer(
		"obsidian-cli-mcp",
		"1.0.0",
		server.WithToolCapabilities(true),
	)

	registerNoteTools(s)
	registerDailyTools(s)
	registerSearchTools(s)
	registerPropertyTools(s)
	registerTagTools(s)
	registerTaskTools(s)
	registerFileTools(s)
	registerLinkTools(s)
	registerVaultTools(s)

	if err := server.ServeStdio(s); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func itoa(n int) string {
	return strconv.Itoa(n)
}
