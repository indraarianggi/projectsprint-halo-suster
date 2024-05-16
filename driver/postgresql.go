package driver

import (
	"log"

	"github.com/backend-magang/halo-suster/config"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func InitPostgres(config config.Config) *sqlx.DB {
	psqlInfo := config.GetDSN()

	db, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	log.Println("[Database] initialized...")

	err = db.Ping()
	if err != nil {
		log.Println("[Database] failed to connect to database: ", err)
		return nil
	}

	log.Println("[Database] successfully connected")
	return db
}
