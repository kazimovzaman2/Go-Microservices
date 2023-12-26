package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

type DBInstance struct {
	Db *sql.DB
}

var Database DBInstance

func ConnectDB() {
	dbUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	orm.RegisterDriver("postgres", orm.DRPostgres)

	if err := orm.RegisterDataBase("default", "postgres", dbUrl); err != nil {
		log.Fatal(err.Error())
	}

	if db, err := sql.Open("postgres", dbUrl); err != nil {
		log.Fatal(err.Error())
	} else {
		if err := orm.RunSyncdb("default", false, true); err != nil {
			log.Fatal(err.Error())
		} else {
			Database = DBInstance{Db: db}
			log.Print("Connected Succcessfully")
		}
	}
}
