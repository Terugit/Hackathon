package dao

import (
	"database/sql"
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

type HackathonDB struct {
	*sql.DB
}

func (todoItem *HackathonDB) scan(rows *sql.Rows) {
	err := rows.Scan(
		&todoItem.Id,
		&todoItem.Title,
		&todoItem.Due,
		&todoItem.Done)
	if err != nil {
		log.Fatalln(err)
	}
}

// DBTodoItems embeds DBTodoItem and adds DB-specific methods
type DBTodoItems []*DBTodoItem

func (todoItems *DBTodoItems) scan(rows *sql.Rows) error {
	for rows.Next() {
		todoItem := &DBTodoItem{&models.TodoItem{}}
		todoItem.scan(rows)
		*todoItems = append(*todoItems, todoItem)
	}
	return rows.Err()
}
