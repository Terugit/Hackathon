package main

//
//// no sql or db driver imports
//import (
//	_ "db/dao"
//	"log"
//)
//
//const (
//	dbname = "todoapp"
//	dbuser = "todoapp"
//)
//
//// note error handling is omitted for brevity - handle err returns!
//func main() {
//	// declare data-access struct
//	db := &data.GratitudeDB{}
//
//	// open PG database
//	db.OpenDb(dbname, dbuser)
//
//	// create DB schema (if needed)
//	db.CreateTablesIfNotExists()
//
//	// fetch DB items as a typed slice of DBTodoItem
//	items, _ := db.SelectAllTodoItems()
//	log.Println(len(*items), "items in the database")
//
//	// declare a new DBTodoItem (with embedded TodoItem)
//	item := &data.DBTodoItem{&models.TodoItem{}}
//	item.Title = "test"
//
//	// insert into the DB
//	db.InsertTodoItem(item)
//	log.Printf("%s (done: %t) inserted into the database", item.Title, item.Done)
//
//	// demonstrate successful insert
//	items, _ = db.SelectAllTodoItems()
//	log.Println(len(*items), "items in the database")
//	item = (*items)[0]
//	log.Printf("%s (done: %t) found in database", item.Title, item.Done)
//
//	// flag item as Done
//	item.Done = true
//
//	// update an item in DB
//	db.UpdateTodoItem(item)
//	log.Printf("%s (done: %t) updated in the database", item.Title, item.Done)
//
//	// demonstrate successful update
//	items, _ = db.SelectAllTodoItems()
//	log.Println(len(*items), "items in the database")
//	item = (*items)[0]
//	log.Printf("%s (done: %t) found in database", item.Title, item.Done)
//
//	// delete item
//	db.DeleteTodoItem(item)
//
//	// demonstrate successful deletion 	items, _ = db.SelectAllTodoItems()
//	log.Println(len(*items), "items in the database")
//}
