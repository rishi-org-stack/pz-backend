package model

import "fmt"

type User struct {
	ID          int    `pg:"id,pk"`
	Name        string `pg:"name"`
	Dob         string `pg:"dob"`
	Email       string `pg:"email"`
	PhoneNumber string `pg:"phone_number"`
}
func (u *User) Save(db *DB) error {
	_,err := db.Db.Model(u).Insert(u)
	return err
}


func (u *User) GetallUser(db *DB) []User {
	var users []User
	err := db.Db.Model(&users).Select()
	if err != nil {
		fmt.Println(err)
	}
	return users
}
