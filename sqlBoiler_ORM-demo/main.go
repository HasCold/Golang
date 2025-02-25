package main

import (
	// Import qm so you can use its functions directly (like Limit, Where, etc.)
	"context"
	"database/sql" // Standard library package to work with SQL databases.
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql" // Blank import to register the MySQL driver.
	"github.com/volatiletech/sqlboiler/v4/boil"
	. "github.com/volatiletech/sqlboiler/v4/queries/qm"
	"sqlboiler.com/models"
)

var DB *sql.DB

func init() {
	var err error
	// Open a connection to the MySQL database.
	// Adjust connection string with your MySQL credentials and database name.
	DB, err = sql.Open("mysql", "root:Hasan@tcp(localhost:3306)/institute?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}

	// Set global DB
	boil.SetDB(DB)

}

func main() {
	defer DB.Close()
	ctx := context.Background()

	// Example: Read operation using the qm package for query modifiers.
	// Let's say we want to fetch test courses with fees greater than 4000,
	// limiting the result to 2 records, and ordering them by fee.
	courses, err := models.TestCourses(
		Where("fees > ?", 4000),
		OrderBy("fees DESC"),
		Limit(2),
	).All(ctx, DB)

	if err != nil {
		log.Fatalf("Error fetching courses: %v", err)
	}

	// Print courses
	fmt.Println("Courses Available:")
	for _, course := range courses {
		fmt.Printf("ID: %d | Name: %s | Fees: %d\n", course.ID, course.CourseName, course.Fees)
	}
}

// Generate SQLBoiler Models
// Run the following command to generate models for your institute database:

// sqlboiler mysql --add-global-variants --add-panic-variants --no-tests --output models
// This will generate the models package inside your project.
