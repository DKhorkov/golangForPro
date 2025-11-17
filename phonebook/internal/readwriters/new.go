package readwriters

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/filepath"
	"github.com/DKhorkov/golangForPro/phonebook/internal/interfaces"
	"strings"
)

func New(path string) (interfaces.ReadWriter, error) {
	switch {
	case strings.HasSuffix(path, filepath.CSVExtension):
		return NewCSVReadWriter(path), nil
	case strings.HasSuffix(path, filepath.JSONExtension):
		return NewJSONReadWriter(path), nil
	}

	return nil, fmt.Errorf(`invalid file extension "%s"`, path)
}
