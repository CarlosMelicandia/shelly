package operations

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/helpers/user"
	"github.com/weareinit/Opal/internal/tools"
	"github.com/weareinit/Opal/internal/utils"
)

func CreateHacker(w http.ResponseWriter, r *http.Request, hacker api.Hacker) (any, error) {
	userId, err := user.GetUserId(w, r)
	if err != nil {
		return nil, fmt.Errorf("could not get the token from user")
	}

	// Ensure the hacker application has the correct userId
	hacker.UserId = userId

	insertQuery := `
		INSERT INTO hacker_application (
			user_id, first_name, last_name, age, school, major, grad_year,
			level_of_study, country, email, phone_number, resume_path, github,
			linkedin, is_international, gender, pronouns, ethnicity, avatar,
			agreed_mlh_news, application_status, created_at, updated_at
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	return tools.LoadDB(func(db *sql.DB) (any, error) {
		_, err := db.Exec(
			insertQuery,
			hacker.UserId,
			hacker.FirstName,
			hacker.LastName,
			hacker.Age,
			hacker.School,
			hacker.Major,
			hacker.GradYear,
			hacker.LevelOfStudy,
      hacker.Country,
			hacker.Email,
			hacker.PhoneNumber,
			hacker.ResumePath,
      utils.ToNullString(hacker.Github),
      utils.ToNullString(hacker.Linkedin),
			hacker.IsInternational,
			hacker.Gender,
			hacker.Pronouns,
			hacker.Ethnicity,
      0, // avatar
			hacker.AgreedMLHNews,
      0, // application status: registered
      time.Now().Unix(),
      time.Now().Unix(),
		)

		if err != nil {
			return nil, fmt.Errorf("failed to insert the hacker application: %v", err)
		}

		return nil, nil
	})
}
