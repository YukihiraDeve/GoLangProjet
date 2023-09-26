package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

// Crée un program qui permet de se connecter a une base de donnée sur Docker et d'enregistrer des données de l'API avec un timestamp
func main() {

	resp, err := http.Get("https://api.nuki.io/#/Smartlock/smartLockState")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	connect(string(body))
}

func connect(value string) {
	value = "value"
	db, err := sql.Open("postgres", "user=postgres password=postgrespw host=host.docker.internal port=55001 dbname=postgres sslmode=disable")
	if err != nil {
		log.Fatalf("Error : Unable to connect to database : %v", err)
	}
	sqlStatement := `INSERT INTO return_api VALUES ($1)`
	result, err := db.Exec(sqlStatement, value)
	if err != nil {
		log.Fatalf("Error : Unable to execute update : %v", err)
	}
	affectedRows, _ := result.RowsAffected()
	fmt.Printf("Updated %d rows\n", affectedRows)
	defer db.Close()
}
