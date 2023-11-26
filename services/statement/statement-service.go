package statement

import (
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Prints the wallet balances of top 10 users
// Pagination ignored as part of this example
func ViewBalances(db *sql.DB) error {
	rows, err := db.Query("SELECT u.username, w.balance FROM users u INNER JOIN wallet w ON u.id = w.user_id ORDER BY w.balance DESC LIMIT 10")
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	fmt.Println("Showing the top 10 users with highest balance:")

	for rows.Next() {
		var username string
		var balance float64
		err = rows.Scan(&username, &balance)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("%s: %.2f\n", username, balance)
	}

	return nil
}

// View transactions of the user by scanning the username
// Fetch latest 20 transactions order by txn_date
func ViewTransactions(db *sql.DB) error {
	var username string
	fmt.Print("Enter username: ")
	fmt.Scan(&username)

	rows, err := db.Query("SELECT t.txn_type, t.txn_date, t.txn_amount, t.closing_balance FROM transactions t INNER JOIN users u ON t.user_id = u.id WHERE u.username = ? ORDER BY t.txn_date DESC LIMIT 20", username)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer rows.Close()

	fmt.Println("Showing the latest 20 transactions:")

	for rows.Next() {
		var txn_type string
		var txn_date string
		var txn_amount float64
		var closing_balance float64
		err = rows.Scan(&txn_type, &txn_date, &txn_amount, &closing_balance)
		if err != nil {
			fmt.Println(err)
			return err
		}
		fmt.Printf("%s | %s | %.2f | %.2f\n", txn_type, txn_date, txn_amount, closing_balance)
	}

	return nil
}
