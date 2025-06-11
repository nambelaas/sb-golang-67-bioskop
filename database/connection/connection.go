package connection

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/spf13/viper"
)

var (
	DbConn *sql.DB
	err    error
)

func Init() {
	dbData := fmt.Sprintf("host=%s user=%s port=%d password=%s dbname=%s sslmode=disable",
		viper.GetString("database.host"),
		viper.GetString("database.username"), viper.GetInt("database.port"), viper.GetString("database.password"), viper.GetString("database.name"))

	DbConn, err = sql.Open("postgres", dbData)

	err = DbConn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
