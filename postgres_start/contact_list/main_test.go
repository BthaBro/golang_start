package main

import "testing"

var cs, err = NewContacts()

func TestContactsAdd(t *testing.T) {
	v := &Contact{FirstName: "John", LastName: "Doe", PhoneNumber: "+998983838383", Age: 21}
	added, err := cs.Add(v)

	if err != nil || added.FirstName != v.FirstName {
		t.Error("Error while adding contact")
	}
}

func TestContactsUpdate(t *testing.T) {
	v := &Contact{FirstName: "Frenkie", LastName: "De Yong", PhoneNumber: "+99898234234", Age: 21, ID: 3}
	upd, err := cs.Update(v)

	if err != nil || upd.FirstName != v.FirstName {
		t.Error("Error while updating contact")
	}
}

func TestContactsDelete(t *testing.T) {
	id := 2
	err := cs.Delete(id)

	if err != nil {
		t.Error("Error while deleting contact")
	}
}

func TestContactsShowAllContacts(t *testing.T) {
	_, err := cs.ShowAllContacts()

	if err != nil {
		t.Error("Error while getting all contacts")
	}
}
