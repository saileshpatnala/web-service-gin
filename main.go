package main

import (
	"learn/web-service-gin/db"
	"learn/web-service-gin/db/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// album represents data about a record album
type Album struct {
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

// albums slice to seed record album data.
var albums = []Album{
	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

// getAlbums responds with the list of all albums as JSON
func getTransactions(c *gin.Context) {
	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	transactions, err := models.GetTransactions(db)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusOK, transactions)
}

// postAlbums adds an album from JSON received in the request body
func postTransaction(c *gin.Context) {
	var transactionCreate models.TransactionCreate

	// Call BindJSON to bind the received JSON to newAlbum
	if err := c.BindJSON(&transactionCreate); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Add the new album to the slice
	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	newTransaction, err := models.CreateTransaction(db, transactionCreate)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	c.IndentedJSON(http.StatusCreated, newTransaction)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then retuns that album as a response
func getTransactionByID(c *gin.Context) {
	txn_id := c.Param("id")

	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	transaction, err := models.GetTransactionByID(db, txn_id)
	if err != nil {
		log.Printf("Transaction with ID: %s not found", txn_id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
	} else {
		log.Printf("Found transaction with ID: %s", txn_id)
		c.IndentedJSON(http.StatusOK, transaction)
	}

}

// getDBConn checks the connection to the database
func getDBConn(c *gin.Context) {
	conn, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}
	if conn != nil {
		log.Println("DB Connection active ✅")
		c.IndentedJSON(http.StatusOK, gin.H{"message": "DB Connection active ✅"})
	}
}

func main() {
	router := gin.Default()
	router.GET("/transactions", getTransactions)
	router.GET("/transactions/:id", getTransactionByID)
	router.POST("/transactions", postTransaction)
	router.GET("/db_conn", getDBConn)

	router.Run("localhost:8080")
}
