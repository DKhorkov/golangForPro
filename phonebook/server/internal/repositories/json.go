package repositories

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"os"
)

type JSONEntriesRepository struct {
	filepath string
}

func NewJSONEntriesRepository(filepath string) *JSONEntriesRepository {
	return &JSONEntriesRepository{filepath: filepath}
}

func (r *JSONEntriesRepository) Read() ([]models.Entry, error) {
	fileInfo, err := os.Stat(r.filepath)
	switch {
	case errors.Is(err, os.ErrNotExist):
		if err = r.Write(defaultEntries); err != nil {
			return nil, err
		} // наполняем книгу дефолтными записями:
	case err != nil:
		return nil, err
	default:
		mode := fileInfo.Mode()
		if !mode.IsRegular() {
			return nil, fmt.Errorf("%s not a regular file", r.filepath)
		}
	}

	f, err := os.Open(r.filepath)
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

func (r *JSONEntriesRepository) Write(entries []models.Entry) error {
	f, err := os.Create(r.filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	return encoder.Encode(entries)
}
