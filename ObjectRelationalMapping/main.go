package main

import (
	"GoExercises/ObjectRelationalMapping/domain"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
)

func main() {
	log.Printf("Object Relational Mapping Example Application")
	db, err := gorm.Open("sqlite3", "gorm.db")
	defer db.Close()

	if err != nil {
		log.Fatal("Can not connect to database")
	}

	log.Printf("Migrate")
	db.AutoMigrate(&domain.User{})

	log.Printf("Creating users...")
	db.Create(&domain.User{Name: "User", Surname: "Fortest"})
	db.Create(&domain.User{Name: "Second", Surname: "Alsotest"})

	var user domain.User

	log.Printf("User with first id: ")
	t := db.First(&user)
	log.Println(t.Value)

	log.Printf("User with name eq Second")
	t2 := db.Where(&user, "name = ?", "Second")
	log.Println(t2.Value)

	db.Model(&user).Update("Surname", "Justtest")
}
