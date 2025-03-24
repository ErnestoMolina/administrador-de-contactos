package main

import (
	"go-mysql/database"
	"go-mysql/handlers"
	"log"
)

func main() {
	// Establecer canexion a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	handlers.ListContact(db)
	handlers.GetContactByID(db, 3)
}
