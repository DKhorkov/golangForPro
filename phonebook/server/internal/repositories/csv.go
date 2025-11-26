package repositories

import (
	"encoding/csv"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
)

type CSVEntriesRepository struct {
	filepath string
}

func NewCSVEntriesRepository(filepath string) *CSVEntriesRepository {
	return &CSVEntriesRepository{filepath: filepath}
}

func (r *CSVEntriesRepository) Read() ([]models.Entry, error) {
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

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	entries := make([]models.Entry, 0, len(lines))
	for _, line := range lines {
		lastAccess, err := time.Parse(time.RFC1123, line[3])
		if err != nil {
			return nil, err
		}

		entries = append(
			entries,
			models.Entry{
				Name:       line[0],
				Surname:    line[1],
				Phone:      line[2],
				LastAccess: lastAccess,
			},
		)
	}

	return entries, nil
}

func (r *CSVEntriesRepository) Write(entries []models.Entry) error {
	f, err := os.Create(r.filepath)
	if err != nil {
		return err
	}

	defer f.Close()

	w := csv.NewWriter(f)
	for _, entry := range entries {
		temp := []string{entry.Name, entry.Surname, entry.Phone, entry.LastAccess.Format(time.RFC1123)}
		_ = w.Write(temp)
	}

	w.Flush()

	return nil
}
