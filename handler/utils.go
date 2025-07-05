package handler

import (
	"database/sql"
	"log"
)

func CloseCursor(cursor *sql.Rows) {
	if err := cursor.Close(); err != nil {
		log.Fatal("Ошибка закрытия курсора БД:", err)
	}
}
