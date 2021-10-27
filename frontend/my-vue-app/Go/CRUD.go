//CRUD MySQL
package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
)

//Mandas un tipo de estructura en particular
func insertPerson(persona person) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("INSERT INTO Personas (Persona, Apellido, Email) VALUES (?, ?, ?)", persona.Nombre, persona.Apellido, persona.Email)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func getInfoTableNamed(nombre string, w http.ResponseWriter) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	results, err := db.Query("select * from " + nombre)
	if err != nil {
		fmt.Println("Error when fetching product table rows:", err)
	}
	defer results.Close()

	fmt.Println(reflect.TypeOf(results))

	auxPers := person{}
	arrPers := []person{}

	for results.Next() {
		var (
			id     int
			name   string
			ape    string
			correo string
		)
		err = results.Scan(&id, &name, &ape, &correo)
		if err != nil {
			panic(err.Error())
		}
		auxPers.ID = id
		auxPers.Nombre = name
		auxPers.Apellido = ape
		auxPers.Email = correo

		arrPers = append(arrPers, auxPers)

	}
	defer db.Close()

	//Mandar en JSON los resultados al navegador
	if err := json.NewEncoder(w).Encode(arrPers); err != nil {
		panic(err)
	}
}

func deleteByID(id string) {
	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Query("DELETE FROM personas WHERE codPersona=? ", id)
	if err != nil {
		fmt.Println("Error:", err)
	}
}

func getByID(idSearch int, w http.ResponseWriter) {
	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	results, err := db.Query("select * from personas where CodPersona=?", idSearch)
	if err != nil {
		fmt.Println("Error when fetching product table rows:", err)
	}
	defer results.Close()

	auxPers := person{}
	arrPers := []person{}

	for results.Next() {
		var (
			id     int
			name   string
			ape    string
			correo string
		)
		err = results.Scan(&id, &name, &ape, &correo)
		if err != nil {
			panic(err.Error())
		}
		auxPers.ID = id
		auxPers.Nombre = name
		auxPers.Apellido = ape
		auxPers.Email = correo

		arrPers = append(arrPers, auxPers)

		/*
			if err != nil {
				fmt.Println("Unable to parse row:", err)
			}
			fmt.Printf("ID: %d, Name: '%s', Apellido: %s\n", id, name, ape)
		*/
	}
	defer db.Close()

	//Mandar en JSON los resultados al navegador
	if err := json.NewEncoder(w).Encode(arrPers); err != nil {
		panic(err)
	}

}

func updateByID(persona person) {
	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("UPDATE Personas SET Persona=?, Apellido=?, Email=? WHERE CodPersona=?", persona.Nombre, persona.Apellido, persona.Email, persona.ID)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}
