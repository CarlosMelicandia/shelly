package operations

import (
	"database/sql"
	"fmt"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/tools"
)

func AddUser(user api.User) (any, error) {
    return tools.LoadDB(func(db *sql.DB) (any, error) {
        if user.UserId == "" {
            return nil, fmt.Errorf("UserID is empty, cannot proceed with query")
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
            INSERT INTO user (user_id, name, email) 
            VALUES (?, ?, ?)
        `
        _, err = db.Exec(insertQuery, user.UserId, user.Name, user.Email)
        if err != nil {
            return nil, fmt.Errorf("failed to insert user: %v", err)
        }

        return nil, nil
    })
}

