package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/labstack/echo"

	"github.com/titolins/go-ghostress-test/controllers"
)

func main() {
	e := echo.New()
	db, err := gorm.Open(
		"postgres",
		"host=db port=5432 user=postgres dbname=ghostress_test password=example sslmode=disable")
	ghostController := controllers.NewGhostController(db)

	if err != nil {
		panic(fmt.Sprintf("Error openning db connection: %s", err))
	}
	defer db.Close()

	e.GET("/ghost", ghostController.FindAll())
	e.POST("/ghost", ghostController.CreateGhost())
	e.DELETE("/ghost", ghostController.DeleteGhost())

	e.Logger.Fatal(e.Start(":1323"))
}
