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