package hacker

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/weareinit/Opal/api"
	"github.com/weareinit/Opal/internal/helpers/user"
	"github.com/weareinit/Opal/internal/tools"
)

func GetHacker(w http.ResponseWriter, r *http.Request) (api.Hacker, error) {
	userId, err := user.GetUserId(w, r)
	if err != nil {
		return api.Hacker{}, fmt.Errorf("could not get the token from user")
	}

	return tools.LoadDB(func(db *sql.DB) (api.Hacker, error) {
		var hacker api.Hacker
		getUserQuery := `SELECT user_id, first_name, last_name, age, school, major, grad_year, level_of_study, country, email, phone_number, resume_path, github, linkedin, is_international, gender, pronouns, ethnicity, avatar, agreed_mlh_news, application_status, created_at, updated_at FROM hacker_application WHERE user_id = ?`

		row := db.QueryRow(getUserQuery, userId)

		err := row.Scan(
			&hacker.UserId,
			&hacker.FirstName,
			&hacker.LastName,
			&hacker.Age,
			&hacker.School,
			&hacker.Major,
			&hacker.GradYear,
			&hacker.LevelOfStudy,
			&hacker.Country,
			&hacker.Email,
			&hacker.PhoneNumber,
			&hacker.ResumePath,
			&hacker.Github,
			&hacker.Linkedin,
			&hacker.IsInternational,
			&hacker.Gender,
			&hacker.Pronouns,
			&hacker.Ethnicity,
			&hacker.Avatar,
			&hacker.AgreedMLHNews,
			&hacker.ApplicationStatus,
			&hacker.CreatedAt,
			&hacker.UpdatedAt,
		)

		if err != nil {
			if err == sql.ErrNoRows {
				return api.Hacker{}, fmt.Errorf("no user found with the given ID")
			}
			return api.Hacker{}, fmt.Errorf("failed to retrieve the user: %v", err)
		}

		return hacker, nil
	})
}
