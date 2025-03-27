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
func GetContactByID(db *sql.DB, contactID int) (models.Contact, string) {
	// Consulta SQL para seleccionar un contacto por su ID
	query := "SELECT * FROM contact WHERE idcontact = ?"
	row := db.QueryRow(query, contactID)
	// Instacia del modelo contact
	contact := models.Contact{}

	var valueEmail sql.NullString
	var valuePhone sql.NullString
	var valueName sql.NullString
	var erro string
	// Escanear el resultado en el modelo contact
	err := row.Scan(&contact.Id, &valueName, &valueEmail, &valuePhone)
	if err != nil {
		if err == sql.ErrNoRows {
			erro = "El contacto no existe"
			return contact, erro
		}
	} else {
		contact.Name = validarString(valueName)
		contact.Email = validarString(valueEmail)
		contact.Phone = validarString(valuePhone)
		return contact, erro
	}
	return contact, erro
}

func validarString(valueString sql.NullString) string {
	if valueString.Valid {
		return valueString.String
	} else {
		return "Null"
	}
}

// CreateContact registra un contacto nuevo
func CreateContact(db *sql.DB, contact models.Contact) {
	query := "INSERT INTO contact (name, email, phone) VALUES (?, ?, ?)"
	// Ejecutar la sentencia sql
	_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Nuevo contacto registrado exitosamente")
}

// UpdateContact actualiza un contacto existente en la base de datos
func UpdateContact(db *sql.DB, contact models.Contact) {
	_, erro := GetContactByID(db, contact.Id)
	if erro != "El contacto no existe" {
		query := "UPDATE contact SET name = ?, email = ?, phone = ? WHERE idcontact = ?"
		_, err := db.Exec(query, contact.Name, contact.Email, contact.Phone, contact.Id)
		if err != nil {
			log.Fatal(err)
		}
		log.Println("Contacto modificado exitosamente")
	} else {
		fmt.Println(erro)
	}
}

// DeleteContact elimina un contacto existente de la base de datos
func DeleteContact(db *sql.DB, contactID int) {
	query := "DELETE FROM contact WHERE idcontact = ?"
	_, err := db.Exec(query, contactID)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Se elimino el contacto %d exitosamente", contactID)
}
