package operations

import (
	"database/sql"
	"fmt"

	"github.com/weareinit/Opal/internal/tools"
)

func ConnectDiscordToUser(discordId string, userId string) (any, error) {
    return tools.LoadDB(func(db *sql.DB) (any, error) {
        selectQuery := `SELECT discord_id FROM user WHERE user_id = ?`
        var existingDiscordId string
        err := db.QueryRow(selectQuery, userId).Scan(&existingDiscordId)
        if err != nil {
            if err == sql.ErrNoRows {
                return nil, fmt.Errorf("user with ID %s not found", userId)
            }
            return nil, err }

        if existingDiscordId != discordId {
            updateQuery := `UPDATE user SET discord_id = ? WHERE user_id = ?`
            _, err := db.Exec(updateQuery, discordId, userId)
            if err != nil {
                return nil, err
            }
        }

        return "Discord ID successfully connected", nil
    })
}
