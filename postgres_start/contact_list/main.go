package main

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Contact struct {
	FirstName, LastName, PhoneNumber string
	ID, Age                          int
}

type Contacts struct {
	db *sqlx.DB
}

type Manager interface {
	Add(*Contact) (*Contact, error)
	Update(*Contact) (*Contact, error)
	ShowAllContacts() (*Contacts, error)
	Delete(int) error
}

func NewContacts() (*Contacts, error) {
	cs := Contacts{}
	var err error
	cs.db, err = sqlx.Connect("postgres", "user=postgres dbname=mydb password=123456 host=127.0.0.1")
	if err != nil {
		return nil, err
	}

	return &cs, err
}

func (cs *Contacts) Add(c *Contact) (*Contact, error) {
	query := `INSERT INTO contacts (firstname, lastname, phonenumber, age) values ($1, $2, $3, $4)`

	cs.db.MustExec(query, c.FirstName, c.LastName, c.PhoneNumber, c.Age)

	rows, err := cs.db.Query(`SELECT * FROM contacts ORDER BY id DESC LIMIT 1`)

	if err != nil {
		return nil, err
	}

	var fn, ln, pn string
	var id, age int

	for rows.Next() {
		err = rows.Scan(&id, &fn, &ln, &pn, &age)
		if err != nil {
			return nil, err
		}
		c.ID = id
		c.FirstName = fn
		c.LastName = ln
		c.PhoneNumber = pn
		c.Age = age
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *Contacts) Update(c *Contact) (*Contact, error) {
	id := c.ID

	query := `UPDATE contacts SET firstname = $1, lastname = $2, phonenumber = $3, age = $4 WHERE id = $5`

	cs.db.Exec(query, c.FirstName, c.LastName, c.PhoneNumber, c.Age, c.ID)

	rows, err := cs.db.Query(`SELECT * FROM contacts WHERE id = $1`, id)
	if err != nil {
		return nil, err
	}

	var fn, ln, pn string
	var age int

	for rows.Next() {
		err = rows.Scan(&id, &fn, &ln, &pn, &age)
		if err != nil {
			return nil, err
		}
		c.ID = id
		c.FirstName = fn
		c.LastName = ln
		c.PhoneNumber = pn
		c.Age = age
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}
	return c, nil
}

func (cs *Contacts) Delete(id int) error {
	query := `DELETE FROM contacts WHERE id = $1`

	_, err := cs.db.Exec(query, id)

	if err != nil {
		return err
	}

	return nil
}

func (cs *Contacts) ShowAllContacts() ([]Contact, error) {
	query := `SELECT * FROM contacts`

	rows, err := cs.db.Query(query)

	if err != nil {
		return nil, err
	}

	var fn, ln, pn string
	var id, age int
	contacts := []Contact{}
	c := &Contact{}

	for rows.Next() {
		err := rows.Scan(&id, &fn, &ln, &pn, &age)
		if err != nil {
			return nil, err
		}

		c.FirstName = fn
		c.LastName = ln
		c.PhoneNumber = pn
		c.ID = id
		c.Age = age
		contacts = append(contacts, *c)
	}

	return contacts, nil
}

func main() {
	// db, err := sqlx.Connect("postgres", "user=postgres dbname=mydb password=123456 host=127.0.0.1")
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// // tx := db.MustBegin()
	// // tx.MustExec("INSERT INTO contacts (id, firstname, lastname, phonenumebr, age) values ($1, $2, $3, $4, $5)", 1, "John", "Doe", "+99897777777", 20)
	// // tx.Commit()

	// cs := &Contacts{}
	// rows, err := db.Query(`SELECT * FROM contacts`)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// defer rows.Close()

	// for rows.Next() {
	// 	err := rows.Scan(&cs.contacts)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}

	// 	fmt.Println(cs.contacts)
	// }

	// err = rows.Err()
	// if err != nil {
	// 	log.Fatal(err)
	// }
}
