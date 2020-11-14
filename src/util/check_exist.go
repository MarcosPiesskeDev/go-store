package util

import "go-store-back/src/config"

//Check if exists an entity on table
func VerifySExists(id int, table string) bool {
	db, err := config.GetDb()

	if err != nil {
		return false
	}
	row, _ := db.Query("SELECT * FROM "+table+" WHERE id = ?", id)

	if row.Next() {
		return true
	}
	return false
}
