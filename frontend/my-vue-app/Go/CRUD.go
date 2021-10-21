//CRUD MySQL
package main

import "fmt"

func insertPerson(persona person) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	_, err = db.Exec("INSERT INTO Personas (Persona, Apellido) VALUES (?, ?)", persona.Nombre, persona.Apellido)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
}

func getInfoTableNamed(nombre string) {

	db, err := getDB()
	if err != nil {
		fmt.Println(err)
	}
	results, err := db.Query("select * from " + nombre)
	if err != nil {
		fmt.Println("Error when fetching product table rows:", err)
	}
	defer results.Close()
	for results.Next() {
		var (
			id   int
			name string
			ape  string
		)
		err = results.Scan(&id, &name, &ape)

		if err != nil {
			fmt.Println("Unable to parse row:", err)
		}
		fmt.Printf("ID: %d, Name: '%s', Apellido: %s\n", id, name, ape)
	}

	defer db.Close()

}
