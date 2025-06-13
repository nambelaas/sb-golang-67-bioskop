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
	dbData := fmt.Sprintf("host=%s user=%s port=%s password=%s dbname=%s sslmode=disable",
		viper.GetString("database.railway.pghost"),
		viper.GetString("database.railway.pguser"), 
		viper.GetString("database.railway.pgport"), 
		viper.GetString("database.railway.pgpassword"), 
		viper.GetString("database.railway.pgdatabase"))

	DbConn, err = sql.Open("postgres", dbData)

	err = DbConn.Ping()
	if err != nil {
		panic(err)
	}

	fmt.Println("Successfully connected to database")
}
