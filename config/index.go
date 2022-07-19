package config

import (
	"log"
	"os"

	"github.com/crownss/dark_market/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB()*gorm.DB{
	DBUSER := os.Getenv("DBUSER")
	DBPASSWORD := os.Getenv("DBPASSWORD")
	DBNAME := os.Getenv("DBNAME")
	DBHOST := os.Getenv("DBHOST")
	DBPORT := os.Getenv("DBPORT")
	SSLMODE := os.Getenv("SSLMODE")
	DBTIMEZONE := os.Getenv("DBTIMEZONE")

	conn := "host=" + DBHOST +
		" user=" + DBUSER +
		" password=" + DBPASSWORD +
		" dbname=" + DBNAME +
		" sslmode=" + SSLMODE +
		" port=" + DBPORT +
		" TimeZone=" + DBTIMEZONE

	var err error
	DB, err = gorm.Open(postgres.Open(conn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	} else {
		log.Println("Connection Success")
		Migration()
	}
	return DB
}

func Migration() {
	DB.AutoMigrate(
		&models.Tx{},
		&models.Users{},
		&models.Stuff{},
	)
}