package exo4

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"database/sql"

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

	fmt.Println(string(body))
	connect()
}

func connect() {
	db, err := sql.Open("postgres", "user=postgres password=postgrespw host=host.docker.internal dbname=test sslmode=disable")
	if err != nil {
		log.Fatalf("Error : Unable to connect to database : %v", err)
	}
	defer db.Close()
}
