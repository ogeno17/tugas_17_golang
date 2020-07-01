package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// Menu is struct
type Menu struct {
	IDMenu   int
	Nama     string
	Kategori string
	Harga    int
}

func connect() (*sql.DB, error) {
	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/restaurant")

	if err != nil {
		return nil, err
	}

	return db, nil
}

func tampilkanMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	db, err := connect()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer db.Close()

	rows, err := db.Query("SELECT idMenu, nama, kategori, harga FROM menu")
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	defer rows.Close()

	var result []Menu

	for rows.Next() {
		var each = Menu{}

		err = rows.Scan(&each.IDMenu, &each.Nama, &each.Kategori, &each.Harga)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		result = append(result, each)
	}

	if err = rows.Err(); err != nil {
		fmt.Println(err.Error())
		return
	}

	for _, each := range result {
		fmt.Println("Menu: ", each.Nama, each.Harga)
	}

	message, _ := json.Marshal(result)

	w.WriteHeader(http.StatusOK)
	w.Write(message)
}

func main() {
	http.HandleFunc("/menu", tampilkanMenu)
	fmt.Println("Start Web")
	http.ListenAndServe(":8080", nil)
}
