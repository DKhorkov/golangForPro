package readwriters

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"os"
)

type JSONReadWriter struct {
	filepath string
}

func NewJSONReadWriter(filepath string) *JSONReadWriter {
	return &JSONReadWriter{filepath: filepath}
}

func (rw *JSONReadWriter) Read() ([]models.Entry, error) {
	fileInfo, err := os.Stat(rw.filepath)
	switch {
	case errors.Is(err, os.ErrNotExist):
		if err = rw.Write(defaultEntries); err != nil {
			return nil, err
		} // наполняем книгу дефолтными записями:
	case err != nil:
		return nil, err
	default:
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			return nil, fmt.Errorf("%s not a regular file", rw.filepath)
		}
	}

	f, err := os.Open(rw.filepath)
	if err != nil {
		return nil, err
	}

	defer f.Close()

	var entries []models.Entry
	decoder := json.NewDecoder(f)
	if err = decoder.Decode(&entries); err != nil {
		return nil, err
	}

	return entries, nil
}

func (rw *JSONReadWriter) Write(entries []models.Entry) error {
	f, err := os.Create(rw.filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	return encoder.Encode(entries)
}
