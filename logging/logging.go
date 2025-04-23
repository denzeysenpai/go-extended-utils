package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

func WriteErrorLog(path string, file_name string, algo_code string, details string) error {
	now := time.Now()
	fullPath := filepath.Join(path, file_name)
	dirPath := filepath.Dir(fullPath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", dirPath, err)
	}

	logEntry := fmt.Sprintf(
		"<error-log>\n    <time=\"%s\">\n	<algo code=\"%s\">\n	</algo>\n    <details>\n        %s\n    </details>\n</error-log>\n\n",
		now.Format("15:04:05"), algo_code, details,
	)

	existingContent, err := os.ReadFile(fullPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Failed to read file %s: %w", fullPath, err)
	}

	newContent := append([]byte(logEntry), existingContent...)

	if err := os.WriteFile(fullPath, newContent, 0644); err != nil {
		return fmt.Errorf("Failed to write file %s: %w", fullPath, err)
	}
	fmt.Println("\033[31mError:", details, "\033[0m")
	return nil
}
