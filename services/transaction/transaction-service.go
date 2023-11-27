package user

import (
    "fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	userService "example.com/wallet/services/user"
)

// Function asks for from username, to username and amount & transfers the amount from from username to to username
func CreateTransaction(db *sql.DB) error {
	var fromUsername string
	var toUsername string
	var amount float64

	fmt.Print("Enter from username: ")
	fmt.Scan(&fromUsername)

	fmt.Print("Enter to username: ")
	fmt.Scan(&toUsername)

	fmt.Print("Enter amount: ")
	fmt.Scan(&amount)

	fromUserID, err := userService.GetUserID(db, fromUsername)
	if err != nil {
		fmt.Println(err)
		return err
	}

	toUserID, err := userService.GetUserID(db, toUsername)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println(fromUserID)
	fmt.Println(toUserID)

	err = transferAmount(db, fromUserID, toUserID, amount)
	if err != nil {
		fmt.Println(err)
		return err
	}

	fmt.Println("Transaction successful")

	return nil
}

func transferAmount(db *sql.DB, fromUserID int64, toUserID int64, amount float64) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		return err
	}

	var fromUserBalance float64
	err = tx.QueryRow("SELECT balance FROM wallet WHERE user_id = ?", fromUserID).Scan(&fromUserBalance)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if fromUserBalance < amount {
		fmt.Println("Insufficient balance")
		tx.Rollback()
		return nil
	}

	var toUserBalance float64
	err = tx.QueryRow("SELECT balance FROM wallet WHERE user_id = ?", toUserID).Scan(&toUserBalance)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = ? WHERE user_id = ?", fromUserBalance - amount, fromUserID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = ? WHERE user_id = ?", toUserBalance + amount, toUserID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, txn_type, txn_amount, closing_balance, other_party_id) VALUES (?, ?, ?, ?, ?)", fromUserID, "Debit", amount, fromUserBalance - amount, toUserID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, txn_type, txn_amount, closing_balance, other_party_id) VALUES (?, ?, ?, ?, ?)", toUserID, "Credit", amount, toUserBalance + amount, fromUserID)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}

// Function asks for an existing transaction id and creates reversal entries for those transactions
func CreateRefund(db *sql.DB) error {
	var txnID int64

	fmt.Print("Enter a debit transaction id: ")
	fmt.Scan(&txnID)

	refundTransaction(db, txnID)

	fmt.Println("Refund successful")

	return nil
}

// Function creates reversal entries for the given transaction id
func refundTransaction(db *sql.DB, transactionId int64) error {
	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	var txnType string
	var txnAmount float64
	var closingBalance float64
	var otherPartyId int64
	var userId int64
	var userBalance float64
	var otherPartyBalance float64
	err = tx.QueryRow("SELECT txn_type, txn_amount, closing_balance, other_party_id, user_id FROM transactions WHERE id = ?", transactionId).Scan(&txnType, &txnAmount, &closingBalance, &otherPartyId, &userId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	if txnType == "Credit" || txnType == "Credit Reversal" || txnType == "Debit Reversal" {
		fmt.Println("Cannot refund a credit transaction")
		tx.Rollback()
		return nil
	}

	err = tx.QueryRow("SELECT balance FROM wallet WHERE user_id = ?", userId).Scan(&userBalance)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	err = tx.QueryRow("SELECT balance FROM wallet WHERE user_id = ?", otherPartyId).Scan(&otherPartyBalance)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = ? WHERE user_id = ?", userBalance + txnAmount, userId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("UPDATE wallet SET balance = ? WHERE user_id = ?", otherPartyBalance - txnAmount, otherPartyId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, txn_type, txn_amount, closing_balance, other_party_id) VALUES (?, ?, ?, ?, ?)", userId, "Debit Reversal", txnAmount, userBalance + txnAmount, otherPartyId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	_, err = tx.Exec("INSERT INTO transactions (user_id, txn_type, txn_amount, closing_balance, other_party_id) VALUES (?, ?, ?, ?, ?)", otherPartyId, "Credit Reversal", txnAmount, otherPartyBalance - txnAmount, userId)
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Println(err)
		tx.Rollback()
		return err
	}

	return nil
}
