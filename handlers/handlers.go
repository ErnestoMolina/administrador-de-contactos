package handlers

import (
	"database/sql"
	"fmt"
	"go-mysql/models"
	"log"
)

// ListContacts lista de todos los contactos desde la base de datos
func ListContact(db *sql.DB) {
	// Consulta SQL para seleccionar todos los contactos
	query := "SELECT * FROM contact"

	// Ejecutar la consulta
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Iterrar sobre los resultados y mostrarlos
	fmt.Println("\nLista de contactos")
	fmt.Println("--------------------------------------------------------")
	for rows.Next() {
		contact := models.Contact{}

		var valueEmail sql.NullString
		var valuePhone sql.NullString
		var valueName sql.NullString

		err := rows.Scan(&contact.Id, &valueName, &valueEmail, &valuePhone)
		if err != nil {
			log.Fatal(err)
		}
		// Validar datos
		contact.Name = validarString(valueName)
		contact.Email = validarString(valueEmail)
		contact.Phone = validarString(valuePhone)

		fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
			contact.Id, contact.Name, contact.Email, contact.Phone)
		fmt.Println("--------------------------------------------------------")
	}
}

// GetContactById obtiene un contacto de la base de datos mediante su ID
func GetContactByID(db *sql.DB, contactID int) {
	// Consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE idcontact = ?"
	row := db.QueryRow(query, contactID)
	// Instacia del modelo contact
	contact := models.Contact{}

	var valueEmail sql.NullString
	var valuePhone sql.NullString
	var valueName sql.NullString

	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &valueName, &valueEmail, &valuePhone)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Fatalf("El contacto %d no existe", contactID)
		}
	}
	contact.Name = validarString(valueName)
	contact.Email = validarString(valueEmail)
	contact.Phone = validarString(valuePhone)

	fmt.Println("\nLista de un contacto")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
		contact.Id, contact.Name, contact.Email, contact.Phone)
	fmt.Println("--------------------------------------------------------")
}

func validarString(valueString sql.NullString) string {
	if valueString.Valid {
		return valueString.String
	} else {
		return "Null"
	}
}
