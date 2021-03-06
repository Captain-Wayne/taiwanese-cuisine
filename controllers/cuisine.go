package controllers

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"strconv"

	"github.com/danielgg-coding/taiwanese-cuisine/queries"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"

	"cloud.google.com/go/firestore"
)

// GetCuisine get cuisine by id
func GetCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		cuisineIDString := c.Param("cuisineId")
		cuisineID, err := strconv.Atoi(cuisineIDString)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		cuisines, err := queries.GetCuisine(db, cuisineID)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.JSON(200, cuisines)
	}

	return gin.HandlerFunc(fn)

}

// GetAllCuisine get all cuisine
func GetAllCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		cuisines, err := queries.GetAllCuisine(db)

		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}

		c.JSON(200, cuisines)
	}

	return gin.HandlerFunc(fn)

}

// VoteCuisine write comparison entry
func VoteCuisine(db *sql.DB) gin.HandlerFunc {

	fn := func(c *gin.Context) {

		winner := c.Query("winner")
		losser := c.Query("losser")

		c.String(200, fmt.Sprintf("id %s beats id %s", winner, losser))
	}

	return gin.HandlerFunc(fn)
}

// GetCuisineFirestore gets cuisine data from firestore
func GetCuisineFirestore(client *firestore.Client) gin.HandlerFunc {
	fn := func(c *gin.Context) {

		dsnap, err := client.Collection("cuisine").Doc("food").Get(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		m := dsnap.Data()
		c.JSON(200, m)
	}

	return gin.HandlerFunc(fn)
}
