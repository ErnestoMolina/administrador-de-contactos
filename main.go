package main

import (
	"bufio"
	"fmt"
	"go-mysql/database"
	"go-mysql/handlers"
	"go-mysql/models"
	"log"
	"os"
	"strings"
)

func main() {
	// Establecer canexion a la base de datos
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	// updateContact := models.Contact{
	// 	Id:    8,
	// 	Name:  "Helver",
	// 	Email: "Helver@gmail.com",
	// 	Phone: "56789",
	// }
	// handlers.UpdateContact(db, updateContact)

	for {
		fmt.Println("\nMenú:")
		fmt.Println("1. Listar contactos")
		fmt.Println("2. Obtener contacto por ID")
		fmt.Println("3. Crear nuevo contacto")
		fmt.Println("4. Actualizar contacto por ID")
		fmt.Println("5. Eliminar contacto por ID")
		fmt.Println("6. Salir")
		fmt.Print("Seleccione una opción: ")

		// Leer la opcion seleccionada por el usuario
		var option int
		fmt.Scanln(&option)
		switch option {
		case 1:
			handlers.ListContact(db)
		case 2:
			var contactID int

			fmt.Print("Digite el ID del contacto: ")
			fmt.Scanln(&contactID)
			contact, err := handlers.GetContactByID(db, contactID)
			if err != "El contacto no existe" {
				fmt.Println("\nLista de un contacto")
				fmt.Println("--------------------------------------------------------")
				fmt.Printf("ID: %d, Nombre: %s, Email: %s, Telefono: %s\n",
					contact.Id, contact.Name, contact.Email, contact.Phone)
				fmt.Println("--------------------------------------------------------")
			} else {
				fmt.Println(err)
			}
		case 3:
			newContact := inputContactDetails(option)
			handlers.CreateContact(db, newContact)
		case 4:
			updateContact := inputContactDetails(option)
			handlers.UpdateContact(db, updateContact)
		case 5:
			var contactID int
			fmt.Print("Digite el ID del contacto que quiere eliminar: ")
			fmt.Scanln(&contactID)
			handlers.DeleteContact(db, contactID)
		case 6:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion no valida")
		}
	}
}

func inputContactDetails(option int) models.Contact {
	reader := bufio.NewReader(os.Stdin)

	var contact models.Contact
	if option == 4 {
		var idContact int
		fmt.Print("Ingrese el ID del contacto: ")
		fmt.Scanln(&idContact)

		contact.Id = idContact
	}
	fmt.Print("Digite el nombre de contacto: ")
	name, _ := reader.ReadString(('\n'))
	contact.Name = strings.TrimSpace(name)
	fmt.Print("Digite el correo electronico del contacto: ")
	email, _ := reader.ReadString(('\n'))
	contact.Email = strings.TrimSpace(email)
	fmt.Print("Digite el telefono del contacto: ")
	phone, _ := reader.ReadString(('\n'))
	contact.Phone = strings.TrimSpace(phone)
	return contact
}
