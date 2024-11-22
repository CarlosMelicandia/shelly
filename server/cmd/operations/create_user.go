package operations

import (
	"database/sql"
	"fmt"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/tools"
)

func CreateUser(user api.User) (any, error) {
    return tools.LoadDB(func(db *sql.DB) (any, error) {
        if user.UserId == "" {
            return nil, fmt.Errorf("UserId is empty, cannot proceed with query")
        }

        var exists bool
        checkQuery := `SELECT EXISTS (SELECT 1 FROM user WHERE user_id = ?)`
        err := db.QueryRow(checkQuery, user.UserId).Scan(&exists)
        if err != nil {
            return nil, fmt.Errorf("failed to check if user exists: %v", err)
        }

        if exists {
            return nil, nil
        }

        insertQuery := `
            INSERT INTO user (user_id, first_name, last_name, email) 
            VALUES (?, ?, ?, ?)
        `
        _, err = db.Exec(insertQuery, user.UserId, user.FirstName, user.LastName, user.Email)
        if err != nil {
            return nil, fmt.Errorf("failed to insert user: %v", err)
        }

        return nil, nil
    })
}

