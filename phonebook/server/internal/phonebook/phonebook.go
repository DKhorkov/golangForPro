package phonebook

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/server/internal/models"
	"sort"
)

type PhoneBook struct {
	rw      interfaces.ReadWriter
	record  models.PhoneBook
	indexes map[string]int
}

func New(rw interfaces.ReadWriter) (*PhoneBook, error) {
	pb := new(PhoneBook)
	pb.rw = rw

	entries, err := rw.Read()
	if err != nil {
		return nil, err
	}

	pb.record = entries

	pb.createIndex()

	return pb, nil
}

func (pb *PhoneBook) createIndex() {
	sort.Sort(pb.record)

	pb.indexes = make(map[string]int)
	for i, entry := range pb.record {
		pb.indexes[entry.Phone] = i
	}
}

func (pb *PhoneBook) List(reverse bool) ([]models.Entry, error) {
	if reverse {
		sort.Sort(sort.Reverse(pb.record))
	} else {
		sort.Sort(pb.record)
	}

	return pb.record, nil
}

func (pb *PhoneBook) Search(key string) (*models.Entry, error) {
	index, ok := pb.indexes[key]
	if !ok {
		return nil, fmt.Errorf("no entry found: %s", key)
	}

	return &pb.record[index], nil
}

func (pb *PhoneBook) Insert(entry models.Entry) error {
	if _, ok := pb.indexes[entry.Phone]; ok {
		return fmt.Errorf("entry already exists: %s", entry.Phone)
	}

	pb.record = append(pb.record, entry)

	pb.createIndex()

	entries := make([]models.Entry, 0, len(pb.record))
	for _, entry = range pb.record {
		entries = append(entries, entry)
	}

	return pb.rw.Write(entries)
}

func (pb *PhoneBook) Delete(key string) error {
	index, ok := pb.indexes[key]
	if !ok {
		return fmt.Errorf("no entry found: %s", key)
	}

	pb.record = append(pb.record[:index], pb.record[index+1:]...)

	delete(pb.indexes, key)

	entries := make([]models.Entry, 0, len(pb.record))
	for _, entry := range pb.record {
		entries = append(entries, entry)
	}

	return pb.rw.Write(entries)
}
