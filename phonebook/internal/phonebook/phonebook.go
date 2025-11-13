package phonebook

import (
	"fmt"
	"sort"
	"strconv"
	"time"

	"github.com/DKhorkov/golangForPro/phonebook/internal/commands"
	"github.com/DKhorkov/golangForPro/phonebook/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
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

func (pb *PhoneBook) Execute(command models.Command) error {
	switch command.Name {
	case commands.CommandList:
		reverse, err := strconv.ParseBool(command.Params[0])
		if err != nil {
			return err
		}

		return pb.list(reverse)
	case commands.CommandSearch:
		return pb.search(command.Params[0])
	case commands.CommandInsert:
		return pb.insert(command.Params[0], command.Params[1], command.Params[2])
	case commands.CommandDelete:
		return pb.delete(command.Params[0])
	}

	return nil
}

func (pb *PhoneBook) list(reverse bool) error {
	if reverse {
		sort.Sort(sort.Reverse(pb.record))
	}

	for _, entry := range pb.record {
		fmt.Println(entry.View())
	}

	return nil
}

func (pb *PhoneBook) search(key string) error {
	index, ok := pb.indexes[key]
	if !ok {
		return fmt.Errorf("no entry found: %s", key)
	}

	fmt.Println(pb.record[index].View())

	return nil
}

func (pb *PhoneBook) insert(name, surname, phone string) error {
	if _, ok := pb.indexes[phone]; ok {
		return fmt.Errorf("entry already exists: %s", phone)
	}

	entry := models.Entry{
		Name:       name,
		Surname:    surname,
		Phone:      phone,
		LastAccess: time.Now(),
	}

	pb.record = append(pb.record, entry)

	pb.createIndex()

	entries := make([]models.Entry, 0, len(pb.record))
	for _, entry = range pb.record {
		entries = append(entries, entry)
	}

	return pb.rw.Write(entries)
}

func (pb *PhoneBook) delete(key string) error {
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
