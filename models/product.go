package models

import (
	"fmt"
	"time"

	"github.com/garygause/go-api-app/db"
)

type Product struct {
	ID int64
	Title string `binding:"required"`
	Description string
	Price float64
	Status string
	CreatedAt time.Time
	StoreID int64
	Store *Store
}

func (p *Product) Save() error {
	query := `
	INSERT INTO products 
	(title, description, price, status, store_id, createdAt) 
	VALUES
	(?, ?, ?, ?, ?, ?)`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		panic("Prepare failed")
		//return err
	}
	fmt.Println("Prepare succeeded")
	defer stmt.Close()
	result, err := stmt.Exec(p.Title, p.Description, p.Price, p.Status, p.StoreID, time.Now())
	if err != nil {
		fmt.Println("Exec error")
		panic(err)
		//return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		panic("last id failed")
	}
	p.ID = id
	return err
}

func (p Product) Update() error {
	query := `
	UPDATE products 
	SET title=?, description=?, price=?, status=?, store_id=?
	WHERE
	id = ?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		//panic("Prepare failed")
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(p.Title, p.Description, p.Price, p.Status, p.StoreID, p.ID)
	if err != nil {
		panic(err)
		//return err
	}
	return err
}

func (p Product) Delete() error {
	query := "DELETE FROM products WHERE id = ?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return nil
	}
	_, err = stmt.Exec(p.ID)
	return err
}

func GetProductById(id int64) (*Product, error) {
	query := `
	SELECT p.*, s.user_id 
	FROM products AS p
	INNER JOIN stores AS s ON p.store_id = s.id
	WHERE id = ?
	`
	row := db.DB.QueryRow(query, id)
	p := &Product{}
	s := &Store{}
	err := row.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Status, &p.CreatedAt, &p.StoreID, &s.UserID)
	if err != nil {
		return nil, err
	}
	p.Store = s
	return p, nil
}

func GetAllProducts() ([]*Product, error) {
	query := `
	SELECT p.*, s.user_id 
	FROM products AS p
	INNER JOIN stores AS s ON p.store_id = s.id
	`
	rows, err := db.DB.Query(query)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
  defer rows.Close()
	var products []*Product
	for rows.Next() {
		p := &Product{}
		s := &Store{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Status, &p.CreatedAt, &p.StoreID, &s.UserID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		p.Store = s
		products = append(products, p)
	}
	return products, nil
}

func GetAllProductsByStore(storeId int64) ([]*Product, error) {
	query := `
	SELECT p.*, s.user_id 
	FROM products AS p
	INNER JOIN stores AS s ON p.store_id = s.id
	WHERE p.store_id = ?
	`

	rows, err := db.DB.Query(query, storeId)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
  defer rows.Close()

	var products []*Product
	for rows.Next() {
		p := &Product{}
		s := &Store{}
		err := rows.Scan(&p.ID, &p.Title, &p.Description, &p.Price, &p.Status, &p.CreatedAt, &p.StoreID, &s.UserID)
		if err != nil {
			fmt.Println(err)
			return nil, err
		}
		p.Store = s
		products = append(products, p)
	}
	return products, nil
}
