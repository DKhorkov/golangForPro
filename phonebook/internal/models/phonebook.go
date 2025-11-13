package models

type PhoneBook []Entry

func (pb PhoneBook) Len() int {
	return len(pb)
}

func (pb PhoneBook) Less(i, j int) bool {
	if pb[i].Surname == pb[j].Surname {
		return pb[i].Name < pb[j].Name
	}

	return pb[i].Surname < pb[j].Surname
}

func (pb PhoneBook) Swap(i, j int) {
	pb[i], pb[j] = pb[j], pb[i]
}
