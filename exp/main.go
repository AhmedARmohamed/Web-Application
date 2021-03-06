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
	us.DestructiveReset()
	user := models.User{
		Name:  "Michael Scott",
		Email: "michael@dundermifflin.com",
	}
	if err := us.Create(&user); err != nil {
		panic(err)
	}

	if err := us.Delete(user.ID); err != nil {
		panic(err)
	}
	// user.Email = "michael@michaelscottpaperco.com"
	// if err := us.Update(&user); err != nil {
	// 	panic(err)
	// }
	// userByEmail, err := us.ByEmail("michael@michaelscottpaperco.com")
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(userByEmail)

	userByID, err := us.ByID(user.ID)
	if err != nil {
		panic(err)
	}
	fmt.Println(userByID)
}
