package main

import (
	"fmt"

	"github.com/AhmedARmohamed/web-applications/models"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "ahmed"
	dbname = "lenslocked_dev"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	us, err := models.NewUserService(psqlInfo)
	if err != nil {
		panic(err)
	}
	defer us.Close()
	// us.DestructiveReset()
	user, err := us.ByID(2)
	if err != nil {
		panic(err)
	}
	fmt.Println(user)
}
