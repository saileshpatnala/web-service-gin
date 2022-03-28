package main

import (
	"fmt"
	"learn/web-service-gin/db"
	"learn/web-service-gin/db/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// getTransactions responds with the list of all transactions as JSON
func getTransactions(c *gin.Context) {
	// Connect to DB
	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Retrieve transactions from DB
	transactions, err := models.GetTransactions(db)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Return list of transactions
	c.IndentedJSON(http.StatusOK, transactions)
}

// getTransactionByID locates the transaction whose ID value matches the id
// parameter sent by the client, then retuns that tranasction as a response
func getTransactionByID(c *gin.Context) {
	// Get transaction ID from path params
	txn_id := c.Param("id")

	// Connect to DB
	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Retrieve single transaction with provided ID from DB
	transaction, err := models.GetTransactionByID(db, txn_id)
	if err != nil {
		log.Printf("Transaction with ID: %s not found", txn_id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
	}

	// Return retrieved single transaction
	log.Printf("Found transaction with ID: %s", txn_id)
	c.IndentedJSON(http.StatusOK, transaction)
}

// postTransaction adds a transaction from JSON received in the request body
func postTransaction(c *gin.Context) {
	var transactionCreate models.TransactionCreate

	// Call BindJSON to bind the received JSON to transactionCreate
	if err := c.BindJSON(&transactionCreate); err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Connect to DB
	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Add transaction to DB
	newTransaction, err := models.CreateTransaction(db, transactionCreate)
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Return created transaction
	c.IndentedJSON(http.StatusCreated, newTransaction)
}

// deleteTransaction deletes the transaction whose ID value matches the id
// parameter sent by the client
func deleteTransactionByID(c *gin.Context, id string) {
	// Get transaction ID from path params
	txn_id := c.Param("id")

	// Connect to DB
	db, err := db.GetDB()
	if err != nil {
		log.Println(err.Error())
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
	}

	// Delete transaction from DB
	deletedTransaction, err := models.DeleteTransactionByID(db, txn_id)
	if err != nil {
		log.Printf("Transaction with ID: %s not found", txn_id)
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Transaction not found"})
	}

	// Return ID of the deleted transaction
	c.IndentedJSON(http.StatusOK, gin.H{"message": fmt.Sprintf("Delete transaction with ID: %s", txn_id)})
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
