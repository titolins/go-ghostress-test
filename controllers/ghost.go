package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/titolins/go-ghostress-test/models"
)

// GhostController -> Controller for our ghost model
type GhostController struct {
	db *gorm.DB
}

// NewGhostController -> returns a new controller with the injected db handler
func NewGhostController(db *gorm.DB) *GhostController {
	return &GhostController{db: db}
}

// FindAll -> returns all ghosts in the db
func (ghostController *GhostController) FindAll() func(c echo.Context) error {
	return func(c echo.Context) error {
		var ghost []models.Ghost
		ghostController.db.Find(&ghost)

		fmt.Printf("ghosts = %+v\n", ghost)
		return c.JSON(200, ghost)
	}
}

// CreateGhost -> inserts a ghost into the db
func (ghostController *GhostController) CreateGhost() func(c echo.Context) error {
	return func(c echo.Context) error {
		ghost := new(models.Ghost)
		c.Bind(&ghost)
		return c.JSON(201, ghostController.db.Create(ghost))
	}
}

// DeleteGhost -> deletes a ghost from the db, given it's id
func (ghostController *GhostController) DeleteGhost() func(c echo.Context) error {
	return func(c echo.Context) error {
		ghost := new(models.Ghost)
		c.Bind(&ghost)
		return c.JSON(200, ghostController.db.Delete(ghost))
	}
}
