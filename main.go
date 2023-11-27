package main

import (
	"fmt"
	"os"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	userService "example.com/wallet/services/user"
	statementService "example.com/wallet/services/statement"
	transactionService "example.com/wallet/services/transaction"
)

func openDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", "admin:password@tcp(localhost:3306)/wallet")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func main() {

	db, err := openDB()
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()

	var choice int

	for {
		fmt.Println("----------------------------------------------")
		fmt.Println("1. Add user")
		fmt.Println("2. View balances")
		fmt.Println("3. Simulate transaction")
		fmt.Println("4. Simulate refund")
		fmt.Println("5. View transactions of a user")
		fmt.Println("6. Exit")
		fmt.Println("----------------------------------------------")
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			userService.AddUser(db)
		case 2:
			statementService.ViewBalances(db)
		case 3:
			transactionService.CreateTransaction(db)
		case 4:
			transactionService.CreateRefund(db)
		case 5:
			statementService.ViewTransactions(db)
		case 6:
			fmt.Println("Exiting the program.")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 6.")
		}
	}
}
