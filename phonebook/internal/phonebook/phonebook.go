package phonebook

import (
	"fmt"
	"github.com/DKhorkov/golangForPro/phonebook/internal/commands"
	"github.com/DKhorkov/golangForPro/phonebook/internal/interfaces"
	"github.com/DKhorkov/golangForPro/phonebook/internal/models"
	"time"
)

type PhoneBook struct {
	rw      interfaces.ReadWriter
	storage map[string]models.Entry
}

func New(rw interfaces.ReadWriter) (*PhoneBook, error) {
	pb := new(PhoneBook)
	pb.rw = rw

	entries, err := rw.Read()
	if err != nil {
		return nil, err
	}

	pb.storage = make(map[string]models.Entry, len(entries))
	for _, entry := range entries {
		pb.storage[entry.Phone] = entry
	}

	return pb, nil
}

func (pb *PhoneBook) Execute(command models.Command) error {
	switch command.Name {
	case commands.CommandList:
		return pb.list()
	case commands.CommandSearch:
		return pb.search(command.Params[0])
	case commands.CommandInsert:
		return pb.insert(command.Params[0], command.Params[1], command.Params[2])
	case commands.CommandDelete:
		return pb.delete(command.Params[0])
	}

	return nil
}

func (pb *PhoneBook) list() error {
	for _, entry := range pb.storage {
		fmt.Println(entry.View())
	}

	return nil
}

func (pb *PhoneBook) search(phone string) error {
	if entry, ok := pb.storage[phone]; ok {
		fmt.Println(entry.View())

		return nil
	}

	return fmt.Errorf("no entry found: %s", phone)
}

func (pb *PhoneBook) insert(name, surname, phone string) error {
	if _, ok := pb.storage[phone]; ok {
		return fmt.Errorf("entry already exists: %s", phone)
	}

	pb.storage[phone] = models.Entry{
		Name:       name,
		Surname:    surname,
		Phone:      phone,
		LastAccess: time.Now(),
	}

	entries := make([]models.Entry, 0, len(pb.storage))
	for _, entry := range pb.storage {
		entries = append(entries, entry)
	}

	return pb.rw.Write(entries)
}

func (pb *PhoneBook) delete(key string) error {
	if _, ok := pb.storage[key]; !ok {
		return fmt.Errorf("no entry found: %s", key)
	}

	delete(pb.storage, key)

	entries := make([]models.Entry, 0, len(pb.storage))
	for _, entry := range pb.storage {
		entries = append(entries, entry)
	}

	return pb.rw.Write(entries)
}
