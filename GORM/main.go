package main

import (
	"fmt"
	"log"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Find all the SQL DB driver to this docs :- https://go.dev/wiki/SQLDrivers
// GORM Docs :- https://gorm.io/docs/connecting_to_the_database.html

// Models are defined using normal structs. These structs can contain fields with basic Go types, pointers or aliases of these types, or even custom types, as long as they implement the Scanner and Valuer interfaces from the database/sql package
type GORM struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"<-:create"` // <- means to read and write both or -> it is for only read purposes
	Email string `gorm:"<-:create"`
}

var db *gorm.DB
var err error

func init() {
	// By-Default port of mysql
	dbCredentials := `root:Hasan@tcp(127.0.0.1:3306)/gorm_crud?charset=utf8mb4&parseTime=True&loc=Local`

	// Gorm make the connection open with the MySQL DB
	db, err = gorm.Open(mysql.Open(dbCredentials), &gorm.Config{
		// This will log the operation info at every task in DB
		Logger: logger.Default.LogMode(logger.Info),
	})

	fmt.Println(db, err)
	if err != nil {
		log.Fatalf("Error in connecting to MySQL DB :- %v", err)
	}

	// We are saying to GORM that takes the GORM struct and pass the object ; Also if db is available then OK otheriwse creates a DB
	db.AutoMigrate(&GORM{})
	fmt.Println("DB connected successfully !")
}

func main() {

	// createDBRecord(db)
	fetchWithConditionRecord(db)
	// fetchAllRecords(db)
	// updateRecord(db)
	// deleteRecord(db)
}

func deleteRecord(db *gorm.DB) {
	resp := db.Table("gorms").Where("name=? OR email=?", "Umer", "umerhello@gmail.com").Delete(&GORM{})

	fmt.Println("The updated value is :-", resp, resp.RowsAffected, resp.Error)
}

func updateRecord(db *gorm.DB) {
	// Find the first user with id 1 and email not equal to
	resp := db.Table("gorms").Where("id=? AND email <> ?", 1, "umerhello@gmail.com").Updates(map[string]interface{}{
		"name":  "Muhammad Hasan Ali",
		"email": "xyz@gmail.com",
	})

	fmt.Println("The updated value is :-", resp, resp.RowsAffected, resp.Error)
}

func fetchAllRecords(db *gorm.DB) {
	var userArray []GORM
	db.Find(&userArray)
	fmt.Println("All Record Data :- ", userArray)
}

func fetchWithConditionRecord(db *gorm.DB) { // you can learn more like this type of conidition :- https://gorm.io/docs/query.html
	var findWithEmail GORM
	var findWithId GORM

	// To prevent the SQL injection like we direclty get the id from params and pass it to the DB query so it's a wrong method
	id := "1"
	idInt, err := strconv.Atoi(id) // Here we are preventing the sql injection
	if err != nil {
		log.Println(err)
	}

	db.Where("email=?", "umerhello@gmail.com").Find(&findWithEmail)
	fmt.Println("The Only FIND Data :- \n \n", findWithEmail)

	// Find the first user with name jinzhu
	db.Where("id=? AND name=?", idInt, "Muhammad Hasan Ali").Find(&findWithId)
	fmt.Println("The Only Condition Data :- \n \n", findWithId)

}

func createDBRecord(db *gorm.DB) {
	u := GORM{Name: "Umer", Email: "umerhello@gmail.com"}
	resp := db.Create(&u)
	fmt.Println(resp.Error, resp.RowsAffected)

}

// Create Table intially in MySQL
// mysql -u root -p
// CREATE DATABASE test_gorm;
// SHOW DATABASES;

// mysql> USE test_gorm;
// Database changed
// mysql> SELECT DATABASE();

// SHOW TABLES;  -->> This will show you the tables present in DB

// SELECT * FROM gorms;
