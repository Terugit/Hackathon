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
	sqlUpdateTodoItem   = ` UPDATE thanks SET  = $2 , from_ = $3 , to_ = $4, message = $5 WHERE id = $1`
	sqlDeleteTodoItem   = ` DELETE FROM thanks WHERE id = $1`
)

// DBThank embeds Gratitude and adds DB-specific methods
// https://golang.org/doc/effective_go.html#embedding
type DBThank struct {
	*model.Gratitude
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
		thank := &DBThank{&model.Gratitude{}}
		thank.scan(rows)
		*thanks = append(*thanks, thank)
	}
	return rows.Err()
}
