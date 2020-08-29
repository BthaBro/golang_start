package main

import "testing"

var cs = NewContacts()

func TestContactsAdd(t *testing.T) {
	v := &Contact{FirstName: "John", LastName: "Doe", PhoneNumber: "+998983838383", ID: 0, Age: 21}
	added := cs.Add(v)
	if added.FirstName != v.FirstName {
		t.Error("Error while adding contact")
	}
}

func TestContactsUpdate(t *testing.T) {
	c := &cs.contacts[0]
	c.FirstName = "Abror"
	c.LastName = "Khamrokhujaev"
	c.Age = 20
	c.PhoneNumber = "+9989747747"
	upd := cs.Update(c)
	if upd != c {
		t.Error("Error while updating contact")
	}
}

func TestShowAllContacts(t *testing.T) {
	contacts := cs.ShowAllContacts()
	if contacts != len(cs.contacts) {
		t.Error("Error while showing all contacts")
	}
}

func TestDelete(t *testing.T) {
	cs.Delete(0)
	if len(cs.contacts) != 0 {
		t.Error("Error while editing contacts")
	}
}
