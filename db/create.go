package db

import (
	"context"
	"database/sql"
	"log"
	"path/filepath"
	"server/utils"
	"time"
)

func CreateTable(db *sql.DB, path string) {
	filePath, _ := filepath.Abs(path)

	query := utils.ParseFile(filePath)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	res, err := db.ExecContext(ctx, query)

	if err != nil {
		panic(err)
	}

	rows, err := res.RowsAffected()

	if err != nil {
		panic(err)
	}

	log.Printf("Rows affected when creating table: %d", rows)
}
