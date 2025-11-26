package usecases

import (
	"fmt"
	customErrors "github.com/DKhorkov/golangForPro/phonebook/server/internal/errors"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"sort"
)

type UseCases struct {
	entriesRepository interfaces.EntriesRepository
	record            models.PhoneBook
	indexes           map[string]int
}

func New(entriesRepository interfaces.EntriesRepository) (*UseCases, error) {
	u := new(UseCases)
	u.entriesRepository = entriesRepository

	entries, err := entriesRepository.Read()
	if err != nil {
		return nil, err
	}

	u.record = entries

	u.createIndex()

	return u, nil
}

func (u *UseCases) createIndex() {
	sort.Sort(u.record)

	u.indexes = make(map[string]int)
	for i, entry := range u.record {
		u.indexes[entry.Phone] = i
	}
}

func (u *UseCases) List(reverse bool) ([]models.Entry, error) {
	if reverse {
		sort.Sort(sort.Reverse(u.record))
	} else {
		sort.Sort(u.record)
	}

	return u.record, nil
}

func (u *UseCases) Search(key string) (*models.Entry, error) {
	index, ok := u.indexes[key]
	if !ok {
		return nil, fmt.Errorf("%w: %s", customErrors.ErrNotFound, key)
	}

	return &u.record[index], nil
}

func (u *UseCases) Insert(entry models.Entry) error {
	if _, ok := u.indexes[entry.Phone]; ok {
		return fmt.Errorf("%w: %s", customErrors.ErrAlreadyExists, entry.Phone)
	}

	u.record = append(u.record, entry)

	u.createIndex()

	entries := make([]models.Entry, 0, len(u.record))
	for _, entry = range u.record {
		entries = append(entries, entry)
	}

	return u.entriesRepository.Write(entries)
}

func (u *UseCases) Delete(key string) error {
	index, ok := u.indexes[key]
	if !ok {
		return fmt.Errorf("%w: %s", customErrors.ErrNotFound, key)
	}

	u.record = append(u.record[:index], u.record[index+1:]...)

	delete(u.indexes, key)

	entries := make([]models.Entry, 0, len(u.record))
	for _, entry := range u.record {
		entries = append(entries, entry)
	}

	return u.entriesRepository.Write(entries)
}
