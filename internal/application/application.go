package application

import "database/sql"

type App struct {
	DbPool *sql.DB
}
