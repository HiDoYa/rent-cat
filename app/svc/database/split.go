package database

import (
	"fmt"

	"github.com/HiDoYa/rent-cat/app/models"
)

// InsertSplit adds a new split
func (db *Client) InsertSplit(split models.Split) (int, error) {
	stmt := "INSERT INTO split (my_percentage, her_percentage) VALUES ($1, $2)"
	result, err := db.client.Exec(stmt, split.MyPercentage, split.HerPercentage)
	if err != nil{
		return 0, fmt.Errorf("DB InsertSplit error: %w", err)
	}

	lastInsertID, err := result.LastInsertId()
	if err != nil{
		return -1, fmt.Errorf("DB InsertExpense error: %w", err)
	}

	return int(lastInsertID), err
}

// SelectSplit retrieves the currently active split
func (db *Client) SelectSplit() (models.Split, error) {
	stmt := "SELECT split_id, my_percentage, her_percentage, created_at FROM split " +
		"WHERE created_at = (SELECT MAX(created_at) FROM split)"

	rows, err := db.client.Query(stmt)

	if err != nil {
		return models.Split{}, fmt.Errorf("DB SelectSplit error: %w", err)
	}
	defer rows.Close()

	var split models.Split
	rows.Next()
	err = rows.Scan(&split.SplitID, &split.MyPercentage, &split.HerPercentage, &split.CreatedAt)
	if err != nil {
		return models.Split{}, fmt.Errorf("DB SelectSplit error: %w", err)
	}

	return split, nil
}