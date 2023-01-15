package main

import (
	"database/sql"
	"os"
	"testing"
)

var testDB *sql.DB

func setup() error {
	var err error

	testDB = connectDB()
	if testDB == nil {
		return err
	}
	return nil
}

func teardown() {
	testDB.Close()
}

func TestMain(m *testing.M) {
	err := setup()
	if err != nil {
		os.Exit(1)
	}

	m.Run()

	teardown()
}

// データベース接続
// func TestConnectDB(t *testing.T) {
// 	db := connectDB()
// 	if db == nil {
// 		t.Error("Failed to connect to DB")
// 	}
// 	defer db.Close()
// }

// SQLのテスト
func TestInsertData(t *testing.T) {
	var name string = "testcat"
	var breed string = "testbreed"
	var age int = 4

	sql := "INSERT INTO cats(name, breed, age) VALUES(?, ?, ?);"
	result, err := testDB.Exec(sql, name, breed, age)
	if err != nil {
		t.Error(err.Error())
	}

	// 実際に影響を与えた行数を返す
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		t.Error("Insertion of data failed")
	}

	// 後処理
	t.Cleanup(func() {
		sql := "DELETE FROM cats WHERE name = ? AND breed = ? AND age = ? ;"
		testDB.Exec(sql, name, breed, age)
	})
}
