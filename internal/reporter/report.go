package reporter

import (
	"encoding/json"
	"gowatcher_g3/internal/checker"
	"os"
)

func ExportResultsToJsonfile(filePath string, results []checker.ReportEntry) error {
	data, err := json.MarshalIndent(results, "", "  ")
	if err != nil {
		return err
	}
	if err := os.WriteFile(filePath, data, 0644); err != nil {
		return err
	}
	return nil
}
