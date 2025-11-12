package readwriters

import (
	"encoding/csv"
	"errors"
	"os"
	"time"

	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
)

type CSVReadWriter struct {
	filepath string
}

func NewCSVReadWriter(filepath string) *CSVReadWriter {
	return &CSVReadWriter{filepath: filepath}
}

func (rw *CSVReadWriter) Read() ([]models.Entry, error) {
	_, err := os.Stat(rw.filepath)
	switch {
	case errors.Is(err, os.ErrNotExist):
		if err = rw.Write(defaultEntries); err != nil {
			return nil, err
		} // наполняем книгу дефолтными записями:
	case err != nil:
		return nil, err
	}

	f, err := os.Open(rw.filepath)
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

func (rw *CSVReadWriter) Write(entries []models.Entry) error {
	f, err := os.Create(rw.filepath)
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
