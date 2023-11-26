package user

import (
    "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Adds a user and initializes their wallet with the given balance
func AddUser(db *sql.DB) error {
	var username string
	var balance float64

	fmt.Print("Enter username: ")
	fmt.Scan(&username)

	fmt.Print("Enter initial balance: ")
	fmt.Scan(&balance)

	userCreateStmt, err := db.Prepare("INSERT INTO users (username) VALUES (?)")
	if err != nil {
		fmt.Println(err)
        return err
    }
    defer userCreateStmt.Close()

	result, err := userCreateStmt.Exec(username)
    if err != nil {
		fmt.Println(err)
        return err
    }

	lastInsertID, err := result.LastInsertId()

	walletCreateStmt, err := db.Prepare("INSERT INTO wallet (user_id, balance) VALUES (?, ?)")
	if err != nil {
		fmt.Println(err)
        return err
    }
    defer walletCreateStmt.Close()

	_, err = walletCreateStmt.Exec(lastInsertID, balance)
    if err != nil {
		fmt.Println(err)
        return err
    }

	fmt.Println("Account created with username & balance")

    return nil
}
