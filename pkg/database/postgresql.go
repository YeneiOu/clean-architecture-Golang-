package database

import (
	"bank/configs"
	"bank/pkg/utils"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"log"
)

func NewPostgreSQLDBConnection(cfg *configs.Configs) (*sqlx.DB, error) {

	postgresUrl, err := utils.ConnectionUrlBuilder("postgresql", cfg)
	if err != nil {
		return nil, err
	}
	db, err := sqlx.Connect("pgx", postgresUrl)
	if err != nil {
		defer db.Close()
		log.Printf("error, can't connect to database, %s", err.Error())
		return nil, err
	}

	log.Println("postgreSQL database has been connected 🐘")
	return db, nil
}
