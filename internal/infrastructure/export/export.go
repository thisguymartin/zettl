package export

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	domain "thisguymartin/zettl/internal/ui"
)

// ExportFormat defines the export format type
type ExportFormat string

const (
	FormatJSON     ExportFormat = "json"
	FormatMarkdown ExportFormat = "markdown"
)

// ExportData represents the exported data structure
type ExportData struct {
	ExportDate time.Time     `json:"export_date"`
	Notes      []domain.Note `json:"notes"`
	Version    string        `json:"version"`
}

// ExportToJSON exports notes to a JSON file
func ExportToJSON(notes []domain.Note, filepath string) error {
	data := ExportData{
		ExportDate: time.Now(),
		Notes:      notes,
		Version:    "1.0",
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return fmt.Errorf("failed to marshal notes: %w", err)
	}

	if err := os.WriteFile(filepath, jsonData, 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}

// ImportFromJSON imports notes from a JSON file
func ImportFromJSON(filepath string) ([]domain.Note, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("failed to read file: %w", err)
	}

	var exportData ExportData
	if err := json.Unmarshal(data, &exportData); err != nil {
		return nil, fmt.Errorf("failed to unmarshal data: %w", err)
	}

	return exportData.Notes, nil
}

// ExportToMarkdown exports notes to individual markdown files
func ExportToMarkdown(notes []domain.Note, dirPath string) error {
	// Create directory if it doesn't exist
	if err := os.MkdirAll(dirPath, 0755); err != nil {
		return fmt.Errorf("failed to create directory: %w", err)
	}

	for _, note := range notes {
		// Create safe filename from title
		filename := sanitizeFilename(note.Title) + ".md"
		filepath := filepath.Join(dirPath, filename)

		// Build markdown content with frontmatter
		var content strings.Builder
		content.WriteString("---\n")
		content.WriteString(fmt.Sprintf("title: %s\n", note.Title))
		content.WriteString(fmt.Sprintf("created: %s\n", note.CreatedAt.Format(time.RFC3339)))
		content.WriteString(fmt.Sprintf("updated: %s\n", note.UpdatedAt.Format(time.RFC3339)))
		content.WriteString(fmt.Sprintf("tags: %s\n", note.Tags))
		content.WriteString("---\n\n")
		content.WriteString(note.Content)

		if err := os.WriteFile(filepath, []byte(content.String()), 0644); err != nil {
			return fmt.Errorf("failed to write file %s: %w", filepath, err)
		}
	}

	return nil
}

// sanitizeFilename removes invalid characters from filename
func sanitizeFilename(name string) string {
	// Replace invalid characters with underscore
	invalid := []string{"/", "\\", ":", "*", "?", "\"", "<", ">", "|"}
	result := name
	for _, char := range invalid {
		result = strings.ReplaceAll(result, char, "_")
	}

	// Limit length
	if len(result) > 200 {
		result = result[:200]
	}

	return result
}
