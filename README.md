# obsidian-cli-mcp

An MCP server that exposes your [Obsidian](https://obsidian.md) vault to AI agents via the
official Obsidian CLI. Agents can read, create, search, and manage notes without any REST API
plugin — just the CLI that ships with Obsidian 1.12.4+.

## Prerequisites

- **Obsidian desktop** must be running
- **Obsidian CLI** enabled: Settings → General → Command line interface → Register CLI → toggle on
- **Go 1.22+** to build from source (or use a pre-built binary)

## Installation

```sh
go install github.com/nicksieger/obsidian-cli-mcp@latest
```

Or build from source:

```sh
git clone https://github.com/nicksieger/obsidian-cli-mcp
cd obsidian-cli-mcp
go build -o obsidian-cli-mcp .
```

## Configuration

### MCP client (Claude Desktop, etc.)

Add to your MCP client's config:

```json
{
  "mcpServers": {
    "obsidian": {
      "command": "obsidian-cli-mcp",
      "env": {
        "OBSIDIAN_VAULT": "My Vault"
      }
    }
  }
}
```

### Environment variables

| Variable | Description |
|----------|-------------|
| `OBSIDIAN_VAULT` | Default vault name. Can be overridden per-call with the `vault` parameter. |

If you only have one vault, you can omit `OBSIDIAN_VAULT` — the Obsidian CLI will use whichever
vault is active.

## Tools

Every tool accepts an optional `vault` parameter that overrides `OBSIDIAN_VAULT` for that call.

### Notes

| Tool | Description |
|------|-------------|
| `note_read` | Read note content by name or path |
| `note_create` | Create a note with optional content and template |
| `note_update` | Replace a note's full content (uses `create --overwrite`) |
| `note_append` | Append content to a note |
| `note_prepend` | Prepend content to a note |
| `note_move` | Move a note to a different folder |
| `note_rename` | Rename a note |
| `note_delete` | Delete a note (moves to system trash) |
| `note_info` | Get file metadata for a note |

### Daily notes

| Tool | Description |
|------|-------------|
| `daily_note_read` | Read today's daily note |
| `daily_note_append` | Append content to today's daily note |
| `daily_note_prepend` | Prepend content to today's daily note |
| `daily_note_path` | Get the file path of today's daily note |

### Search

| Tool | Description |
|------|-------------|
| `vault_search` | Full-text search across the vault |
| `vault_search_context` | Search with surrounding context snippets |

### Properties (frontmatter)

| Tool | Description |
|------|-------------|
| `property_read` | Read a frontmatter property value |
| `property_set` | Set a frontmatter property |
| `property_remove` | Remove a frontmatter property |
| `properties_list` | List all frontmatter properties of a note |

### Tags

| Tool | Description |
|------|-------------|
| `tags_list` | List tags in the vault or a specific note |
| `tag_files` | List files that have a specific tag |

### Tasks

| Tool | Description |
|------|-------------|
| `tasks_list` | List tasks (filter: `todo`, `done`, or all) |
| `task_toggle` | Toggle a task's completion status by line number |

### Files & folders

| Tool | Description |
|------|-------------|
| `files_list` | List files in the vault or a folder |
| `folders_list` | List all folders in the vault |

### Links

| Tool | Description |
|------|-------------|
| `links_list` | Outgoing links from a note |
| `backlinks_list` | Incoming links to a note |
| `orphans_list` | Notes with no incoming links |
| `deadends_list` | Notes with no outgoing links |
| `unresolved_list` | Broken wikilinks |

### Vault

| Tool | Description |
|------|-------------|
| `vault_info` | Information about the current vault |
| `vaults_list` | List all vaults known to Obsidian |

## Notes on usage

- **File vs path**: most tools accept either `file` (partial name match, like a wikilink) or
  `path` (exact path, e.g. `Projects/Meeting.md`). Use `path` when precision matters.
- **Newlines in content**: pass `\n` in the `content` parameter to insert line breaks.
- **Obsidian must be running**: the CLI communicates with the running Obsidian instance; tools
  will return an error if Obsidian is closed.
