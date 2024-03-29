package models

import (
	"errors"
	"fmt"
	"time"

	"github.com/garygause/go-api-app/db"
	"github.com/garygause/go-api-app/utils"
)

type User struct {
	ID int64
	Name string 
	Email string `binding:"required"`
	Password string `binding:"required"`
	Status string
	CreatedAt time.Time
}

var users = []User{}

func (u User) Save() error {
	query := `
	INSERT INTO users 
	(name, email, password, status, createdAt) 
	VALUES
	(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Prepare failed")
		//return err
	}
	defer stmt.Close()

	password, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		//return err
	}
	result, err := stmt.Exec(u.Name, u.Email, password, u.Status, time.Now())
	if err != nil {
		panic(err)
		//return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic("last id failed")
	}
	u.ID = id
	return err
}

func (u User) Update() error {
	query := `
	UPDATE users 
	SET name=?, email=?, password=?, status=?
	WHERE
	id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		//panic("Prepare failed")
		return err
	}

	defer stmt.Close()
	password, err := utils.HashPassword(u.Password)
	if err != nil {
		fmt.Println(err)
		// return err
	}
	_, err = stmt.Exec(u.Name, u.Email, password, u.Status, u.ID)
	if err != nil {
		panic(err)
		//return err
	}
	return err
}

/*
	Delete 

	delete the user and all associated stores, products
*/
func (u User) Delete() error {
	query := "DELETE FROM users WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(u.ID)
	return err
}

func (u *User) AuthenticateUser() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var password string
	err := row.Scan(&u.ID, &password)
	if err != nil {
		return errors.New("invalid credentials")
	}

	isValid := utils.CheckPassword(password, u.Password)
	if !isValid {
		return errors.New("invalid credentials")
	}

	return nil
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var u User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Status, &u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func GetAllUsers() ([]User, error) {
	query := "SELECT * FROM users"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
  defer rows.Close()
	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password, &u.Status, &u.CreatedAt)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

