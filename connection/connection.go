package connection

import (
	"database/sql"
	"fmt"
	"log"
	"os"
    _"github.com/lib/pq"
	"github.com/joho/godotenv"
)
type Connection struct{}
func (con *Connection) CreateConection() *sql.DB {

    if err := godotenv.Load(".env"); err != nil{
		log.Fatal("ErrorLoading in .env File")
	}

	Db , err := sql.Open("postgres" , os.Getenv("POSTGRES_URL"))
	if err != nil{
		panic(err)
	}

	if err = Db.Ping(); err != nil{
		panic(err)
	}

	fmt.Println("Successfully Conneted to DataBase..!")
	return Db
}