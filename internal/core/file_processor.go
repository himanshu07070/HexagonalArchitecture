package core

import (
	"errors"
	"hexagonal-architecture/internal"
	"io"
	"path/filepath"
	"strings"
)

type FileProcessor struct {
	databasePort internal.DatabasePort
	loggerPort   internal.LoggerPort
}

func NewFileProcessor(databasePort internal.DatabasePort, loggerPort internal.LoggerPort) internal.FileProcessorPort {
	return &FileProcessor{
		databasePort: databasePort,
		loggerPort:   loggerPort,
	}
}

func (fp *FileProcessor) ProcessFile(fileName string, file io.Reader) (int64, error) {
	//check if file is already present in database
	fileSize, err := fp.databasePort.Retrieve(fileName)
	if err == nil {
		fp.loggerPort.Info("File already present in database")
		return fileSize, nil
	}
	fileInfo, err := io.ReadAll(file)
	if err != nil {
		return 0, err
	}

	if len(fileInfo) == 0 {
		return 0, errors.New("uploaded file is empty")
	}
	// Check if the file has a valid extension
	if !isValidFileExtension(fileName) {
		return 0, errors.New("invalid file extension")
	}
	fileSize = int64(len(fileInfo))

	err = fp.databasePort.Store(fileName, fileSize)
	if err != nil {
		return 0, err
	}

	fp.loggerPort.Info("File processed successfully")

	return fileSize, nil
}
func isValidFileExtension(filename string) bool {
	allowedExtensions := map[string]bool{
		".txt":  true,
		".pdf":  true,
		".png":  true,
		".jpg":  true,
		".jpeg": true,
		".gif":  true,
	}

	ext := strings.ToLower(filepath.Ext(filename))
	return allowedExtensions[ext]
}
