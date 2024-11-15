package operations

import (
	"database/sql"
	"fmt"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/tools"
)

func AddUser(user api.User) error {
    return tools.LoadDB(func(db *sql.DB) error {
        if user.UserID == "" {
            return fmt.Errorf("UserID is empty, cannot proceed with query")
        }

        var exists bool
        checkQuery := `SELECT EXISTS (SELECT 1 FROM user WHERE user_id = ?)`
        fmt.Printf("Checking existence for UserID: %s\n", user.UserID)
        err := db.QueryRow(checkQuery, user.UserID).Scan(&exists)
        if err != nil {
            return fmt.Errorf("failed to check if user exists: %v", err)
        }

        if exists {
            fmt.Println("User already exists, skipping insert.")
            return nil
        }

        fmt.Println("Inserting new user...")
        insertQuery := `
            INSERT INTO user (user_id, name, email) 
            VALUES (?, ?, ?)
        `
        _, err = db.Exec(insertQuery, user.UserID, user.Name, user.Email)
        if err != nil {
            return fmt.Errorf("failed to insert user: %v", err)
        }

        fmt.Println("User successfully inserted.")
        return nil
    })
}

