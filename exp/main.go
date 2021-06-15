package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres
	_ "github.com/lib/pq"
)

const (
	host   = "localhost"
	port   = 5432
	user   = "ahmed"
	dbname = "lenslocked_dev"
)

type User struct {
	gorm.Model
	Name   string
	Email  string `gorm:"not null;unique_index"`
	Color  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserID      uint
	Amount      int
	Description string
}

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode=disable", host, port, user, dbname)
	db, err := gorm.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	defer db.Close()
	db.LogMode(true)
	db.AutoMigrate(&User{}, &Order{})

	var u User
	if err := db.Preload("orders").First(&u).Error; err != nil {
		panic(err)
	}
	// if err := db.Preload("orders").First(&u).Error; err != nil {
	// 	panic(err)
	// }

	fmt.Println(u)
	fmt.Println(u.Orders)

	// createOrder(db, u, 1001, "Fake Description #1")
	// createOrder(db, u, 999, "Fake Description #2")
	// createOrder(db, u, 100, "Fake Description #3")
}


func createOrder(db *gorm.DB, user User, amount int, desc string) {
	err := db.Create(&Order{
		UserID: user.ID,
		Amount: amount,
		Description: desc,
	}).Error
	if err != nil {
		panic(err)
	}
}