package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func main() {
	// Use the thumbnailHandler function
	http.HandleFunc("/data", setData)
	http.HandleFunc("/get", getData)
	http.HandleFunc("/delete", deleteData)
	http.HandleFunc("/editar", editarUsuario)

	// Serve static files from the frontend/dist directory.
	fs := http.FileServer(http.Dir("../dist"))
	http.Handle("/", fs)

	// Start the server.
	fmt.Println("Server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

}

//go get github.com/go-sql-driver/mysql

var _ = godotenv.Load(".env") // Cargar del archivo llamado ".env"
var (
	ConnectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("user"),
		os.Getenv("pass"),
		os.Getenv("host"),
		os.Getenv("port"),
		os.Getenv("db_name"))
)

const AllowedCORSDomain = "http://localhost"

// Credenciales de conexion y dominios permitidos para hacer peticiones

func getDB() (*sql.DB, error) {
	return sql.Open("mysql", ConnectionString)
}

func conectDB() {
	db, err := getDB()
	if err != nil {
		fmt.Println("Error with database" + err.Error())
		return
	} else {
		err = db.Ping()
		if err != nil {
			fmt.Println("Error making connection to DB. Please check credentials. The error is: " + err.Error())
			return
		}
	}
	defer db.Close()

	/*
		results, err := db.Query("select * from product")
		if err != nil {
			log.Fatal("Error when fetching product table rows:", err)
		}
		defer results.Close()
	*/

}

type person struct {
	ID       int    `json:"id"`
	Nombre   string `json:"nombre"`
	Apellido string `json:"apellido"`
	Email    string `json:"email"`
}

//Recibe datos en formato JSON y los transforma a variables go
func setData(w http.ResponseWriter, request *http.Request) {
	decoder := json.NewDecoder(request.Body)

	var data person
	decoder.Decode(&data)

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	insertPerson(data)

}

func getData(w http.ResponseWriter, request *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	getInfoTableNamed("personas", w)

}

func deleteData(w http.ResponseWriter, r *http.Request) {
	idEmpleado := r.URL.Query().Get("id")
	fmt.Println(idEmpleado)
	deleteByID(idEmpleado)
	http.Redirect(w, r, "/", 301)
}

func editarUsuario(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		decoder := json.NewDecoder(r.Body)
		var data person
		decoder.Decode(&data)
		//fmt.Println(data)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.WriteHeader(http.StatusOK)

		fmt.Println(data.ID)
		getByID(data.ID, w)
	}
	if r.Method == "UPDATE" {

	}

}
