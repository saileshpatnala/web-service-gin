package models

import (
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/mitchellh/mapstructure"
	"gorm.io/gorm"
)

type TransactionCreate struct {
	Description string  `json:"description" mapstructure:"description"`
	BaseType    string  `json:"base_type" mapstructure:"base_type"`
	Amount      float64 `json:"amount" mapstructure:"amount"`
}

type Transaction struct {
	ID                string `json:"id"`
	TransactionCreate `mapstructure:",squash"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

func GetTransactions(db *gorm.DB) ([]Transaction, error) {
	log.Println("Retrieving all transactions from DB...")
	var transactions []Transaction

	result := db.Find(&transactions)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	} else {
		log.Printf("Found %d transactions", len(transactions))
		return transactions, nil
	}
}

func GetTransactionByID(db *gorm.DB, txn_id string) (*Transaction, error) {
	log.Printf("Retreiving transaction with ID: %s...", txn_id)

	var transaction Transaction
	result := db.Where(map[string]interface{}{"id": txn_id}).First(&transaction)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	} else {
		log.Printf("Found transaction with ID: %s", txn_id)
		return &transaction, nil
	}
}

func CreateTransaction(db *gorm.DB, txn_create TransactionCreate) (*Transaction, error) {
	log.Printf("Creating transaction %+v", txn_create)

	var newTransactionMap map[string]interface{}
	mapstructure.Decode(txn_create, &newTransactionMap)
	newTransactionMap["ID"] = uuid.New().String()
	var newTransaction Transaction
	mapstructure.Decode(newTransactionMap, &newTransaction)

	result := db.Create(&newTransaction)
	if result.Error != nil {
		log.Println(result.Error.Error())
		return nil, result.Error
	} else {
		log.Printf("Created transaction with ID: %s", newTransaction.ID)
		return &newTransaction, nil
	}
}

// func DeleteTransactionID(db *gorm.DB, txn_id string) {

// }
