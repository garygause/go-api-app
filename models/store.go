package models

import (
	"fmt"
	"time"

	"github.com/garygause/go-api-app/db"
)

type Store struct {
	ID int64
	Title string `binding:"required"`
	Description string
	Status string
	CreatedAt time.Time
	UserID int64 `binding:"required"`
}

var stores = []Store{}

func (s *Store) Save() error {
	query := `
	INSERT INTO stores 
	(title, description, status, createdAt, user_id) 
	VALUES
	(?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Prepare failed")
		//return err
	}
	defer stmt.Close()
	result, err := stmt.Exec(s.Title, s.Description, s.Status, time.Now(), s.UserID)
	if err != nil {
		panic(err)
		//return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic("last id failed")
	}
	s.ID = id
	return err
}

func (s Store) Update() error {
	query := `
	UPDATE stores 
	SET title=?, description=?, status=?, user_id=?
	WHERE
	id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		//panic("Prepare failed")
		return err
	}

	defer stmt.Close()
	_, err = stmt.Exec(s.Title, s.Description, s.Status, s.UserID, s.ID)
	if err != nil {
		panic(err)
		//return err
	}
	return err
}

func (s Store) Delete() error {
	query := "DELETE FROM stores WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(s.ID)
	return err
}

func GetStoreById(id int64) (*Store, error) {
	query := "SELECT * FROM stores WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var s Store
	err := row.Scan(&s.ID, &s.Title, &s.Description, &s.Status, &s.Status, &s.CreatedAt, &s.UserID)
	if err != nil {
		return nil, err
	}
	return &s, nil
}

func GetAllStores() ([]Store, error) {
	query := "SELECT * FROM stores"
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
  defer rows.Close()
	var stores []Store
	for rows.Next() {
		var s Store
		err := rows.Scan(&s.ID, &s.Title, &s.Description, &s.Status, &s.Status, &s.CreatedAt, &s.UserID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		stores = append(stores, s)
	}
	return stores, nil
}

