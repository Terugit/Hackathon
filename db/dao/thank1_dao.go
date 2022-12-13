package dao

import (
	"database/sql"
	"db/model"
	"log"
)

// / sql statements
const (
	// Entries table
	sqlCreateThanksTable = ` CREATE TABLE IF NOT EXISTS thanks (
                      id char(26) NOT NULL primary key,
                      from_ varchar(50) NOT NULL,
                      to_ varchar(50) NOT NULL,
                      point int(3) NOT NULL,
                      message varchar(140) NOT NULL
);`
	sqlSelectThankCount = ` SELECT COUNT(id) FROM thanks`
	sqlSelectAllThanks  = ` SELECT id, from_, to_, point, message FROM thanks`
	sqlInsertThankItem  = ` INSERT INTO thanks (id, from_, to_, point, message) VALUES ($1, $2, $3, $4, $5)`
	sqlUpdateThankItem  = ` UPDATE thanks SET id = $2 , from_ = $3 , to_ = $4, message = $5 WHERE id = $1`
	sqlDeleteThankItem  = ` DELETE FROM thanks WHERE id = $1`
)

// DBThank embeds Thank and adds DB-specific methods
// https://golang.org/doc/effective_go.html#embedding
type DBThank struct {
	*model.Thank
}

func (thanks *DBThank) scan(rows *sql.Rows) {
	err := rows.Scan(
		&thanks.Id,
		&thanks.From_,
		&thanks.To_,
		&thanks.Point,
		&thanks.Message)
	if err != nil {
		log.Fatalln(err)
	}
}

// DBThanks embeds DBThank and adds DB-specific methods
type DBThanks []*DBThank

func (thanks *DBThanks) scan(rows *sql.Rows) error {
	for rows.Next() {
		thank := &DBThank{&model.Thank{}}
		thank.scan(rows)
		*thanks = append(*thanks, thank)
	}
	return rows.Err()
}

// ThankDB provides methods for accessing DB data
type ThankDB struct {
	db *sql.DB
}

// OpenDb opens a MySQL database as specified
func (gratitudeDb *ThankDB) OpenDb(connStr string) error {
	db, err := sql.Open("mysql", connStr)
	gratitudeDb.db = db
	return err
}

// CreateTablesIfNotExists creates any MySQL tables that do not exist
func (gratitudeDb *ThankDB) CreateTablesIfNotExists() error {
	_, err := gratitudeDb.db.Exec(sqlCreateThanksTable)
	return err
}

// SelectAllThanks returns all rows from the DB as DBThanks
func (gratitudeDb *ThankDB) SelectAllThanks() (items *DBThanks, err error) {
	thanks := &DBThanks{}
	rows, err := gratitudeDb.db.Query(sqlSelectAllThanks)
	if err != nil {
		return nil, err
	}
	thanks.scan(rows)
	return thanks, nil
}

// InsertThank inserts a single DBThank into the DB
func (gratitudeDb *ThankDB) InsertThank(item *DBThank) error {
	_, err := gratitudeDb.db.Exec(
		sqlInsertThankItem,
		item.Id,
		item.From_,
		item.To_,
		item.Point,
		item.Message)
	return err
}

// UpdateThankItem updates a single DBThank within the DB
func (gratitudeDb *ThankDB) UpdateThankItem(item *DBThank) error {
	_, err := gratitudeDb.db.Exec(
		sqlUpdateThankItem,
		item.From_,
		item.To_,
		item.Point,
		item.Message)
	return err
}

// DeleteThankItem deletes a single DBThank within the DB
func (gratitudeDb *ThankDB) DeleteThankItem(item *DBThank) error {
	_, err := gratitudeDb.db.Exec(
		sqlDeleteThankItem,
		item.Id)
	return err
}
