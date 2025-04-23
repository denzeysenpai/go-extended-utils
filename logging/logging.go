package logging

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"time"
)

type Logger struct {
	Name            string
	BasePath        string
	UseDate         bool // Set to true if log file name uses current date
	DefaultFileName string
	FileType        string
}

var SUPPORTED_TYPES = []string{
	".txt",
	".xml",
	".json",
	".html",
}


func NewLogger(name string, base_path string, file_type string) Logger {
	final_path := ".txt"
	if slices.Contains(SUPPORTED_TYPES, file_type) {
		ft := strings.Trim(strings.Trim(file_type, " "), ".")
		for _, val := range SUPPORTED_TYPES {
			if strings.Contains(val, ft) {
				final_path = val
				break
			}
		}
	}
	return Logger{
		Name:            name,
		BasePath:        base_path,
		UseDate:         false,
		DefaultFileName: "",
		FileType:        final_path,
	}
}

func (logger *Logger) Write(header string, details string) error {
	now := time.Now()
	var fullPath string
	if logger.UseDate {
		currentDate := now
		formattedDate := currentDate.Format("2006-01-02")
		fullPath = filepath.Join(logger.BasePath, formattedDate+".xml")
	} else {
		fullPath = filepath.Join(logger.BasePath, logger.DefaultFileName+".xml")
	}
	dirPath := filepath.Dir(fullPath)
	if err := os.MkdirAll(dirPath, os.ModePerm); err != nil {
		return fmt.Errorf("Failed to create directory %s: %w", dirPath, err)
	}

	logEntry := fmt.Sprintf(
		"\n<Log name='%s' time='%s'>\n\t<Header>\n\t\t%s\n\t</Header>\n\t<Details>\n\t\t%s\n\t</Details>\n</Log>\n",
		logger.Name,
		now.Format("15:04:05"),
		header,
		details,
	)

	existingContent, err := os.ReadFile(fullPath)
	if err != nil && !os.IsNotExist(err) {
		return fmt.Errorf("Failed to read file %s: %w", fullPath, err)
	}

	newContent := append([]byte(logEntry), existingContent...)

	if err := os.WriteFile(fullPath, newContent, 0644); err != nil {
		return fmt.Errorf("Failed to write file %s: %w", fullPath, err)
	}
	return nil
}

func (logger *Logger) UseDateAsFileName() {
	logger.UseDate = true
}

func (logger *Logger) UseCustomFileName(name string) {
	logger.UseDate = false
}
