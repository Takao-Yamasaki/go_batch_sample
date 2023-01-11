package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func loadEnv() {
	err := godotenv.Load(".env")
	if err != nil {
		panic("could not read .env.file")
	}
}

func connectDB() *sql.DB {
	loadEnv()

	DBMS := "mysql"
	USER := os.Getenv("MYSQL_USER")
	PASS := os.Getenv("MYSQL_PASSWORD")
	HOST := os.Getenv("MYSQL_HOST")
	DBNAME := "pet_shop"
	CONNECT := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", USER, PASS, HOST, DBNAME)

	db, err := sql.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	db := connectDB()
	defer db.Close()

	fmt.Println("---start---")

	sql := "INSERT INTO cats (name, breed, age) VALUES (?, ?, ?);"
	_, err := db.Exec(sql, "abutaro", "hoshiimo", 4)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("---end---")
}
