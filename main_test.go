package main

import (
	"testing"
)

// データベース接続
func TestConnectDB(t *testing.T) {
	db := connectDB()
	if db == nil {
		t.Error("Failed to connect to DB")
	}
	defer db.Close()
}

// SQLのテスト
func TestInsertData(t *testing.T) {
	db := connectDB()
	defer db.Close()

	sql := "INSERT INTO cats(name, breed, age) VALUES(?, ?, ?);"
	result, err := db.Exec(sql, "abutato", "hoshiimo", 4)
	if err != nil {
		t.Error(err.Error())
	}

	// 実際に影響を与えた行数を返す
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected != 1 {
		t.Error("Insertion of data failed")
	}
}
