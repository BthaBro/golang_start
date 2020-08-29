package main

import "fmt"

type Contact struct {
	FirstName, LastName, PhoneNumber string
	ID, Age                          int
}

type Contacts struct {
	contacts []Contact
}

type Manager interface {
	Add()
	Update()
	ShowAllContacts()
	Delete()
}

func NewContacts() Contacts {
	cs := Contacts{}
	cs.contacts = []Contact{}

	return cs
}

func (cs *Contacts) Add(c *Contact) *Contact {
	c.ID = len(cs.contacts)
	cs.contacts = append(cs.contacts, *c)
	fmt.Println("Contact have been successfully added.")
	return &cs.contacts[c.ID]
}

func (cs *Contacts) Update(c *Contact) *Contact {
	contact := &cs.contacts[c.ID]
	contact.FirstName = c.FirstName
	contact.LastName = c.LastName
	contact.Age = c.Age
	contact.PhoneNumber = c.PhoneNumber
	fmt.Println("Contact has been successfully updated")
	return contact
}

func (cs *Contacts) Delete(id int) {
	cs.contacts = append(cs.contacts[:id], cs.contacts[id+1:]...)
	fmt.Println("Contact has been successfully deleted")
}

func (cs *Contacts) ShowAllContacts() int {
	if len(cs.contacts) == 0 {
		fmt.Println("Contact list is empty")
		return 0
	}
	for _, v := range cs.contacts {
		fmt.Printf("ID: %v\n", v.ID)
		fmt.Printf("First Name: %v\n", v.FirstName)
		fmt.Printf("Last Name: %v\n", v.LastName)
		fmt.Printf("Age: %v\n", v.Age)
		fmt.Printf("Phone Number: %v\n", v.PhoneNumber)
		fmt.Println("---------------------------")
	}
	return len(cs.contacts)
}

// func Menu() {
// 	fmt.Println("********************")
// 	fmt.Println("Contact list Program")
// 	fmt.Println("********************")
// 	fmt.Println("1. Add new contact")
// 	fmt.Println("2. Update existing contact")
// 	fmt.Println("3. Delete existing contact")
// 	fmt.Println("4. Show list of all existing contacts")
// 	fmt.Println("0. Exit")
// 	fmt.Println("\nEnter your choice: ")
// }

// func EnterDetails(c *Contact) {
// 	var fName, lName, ph string
// 	var age int

// 	fmt.Println("Please, enter following details:")
// 	fmt.Println("First Name: ")
// 	fmt.Scanln(&fName)
// 	fmt.Println("Last Name: ")
// 	fmt.Scanln(&lName)
// 	fmt.Println("Age: ")
// 	fmt.Scanln(&age)
// 	fmt.Println("Phone Number:")
// 	fmt.Scanln(&ph)

// 	c.FirstName = fName
// 	c.LastName = lName
// 	c.Age = age
// 	c.PhoneNumber = ph
// }

func main() {

}
