package repositories

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/filepath"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"strings"
)

func New(path string) (interfaces.EntriesRepository, error) {
	switch {
	case strings.HasSuffix(path, filepath.CSVExtension):
		return NewCSVEntriesRepository(path), nil
	case strings.HasSuffix(path, filepath.JSONExtension):
		return NewJSONEntriesRepository(path), nil
	}

	return nil, fmt.Errorf(`invalid file extension "%s"`, path)
}
