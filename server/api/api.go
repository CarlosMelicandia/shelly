// These structs are passed to the client and used on the server. Think of this file like a GraphQL schema.
// We have control over how the json will look like on the client and have autocompletion on the server.
// Add more attributes as you wish but make sure everything works and you might need to use the turso studio to
// make sure everything is correctly updated.

package api

// Anything that has a * next to the type means that the field can potentially be null

type User struct {
	Id          string  `json:"id"`
	UserId      string  `json:"userId"`
	FirstName   string  `json:"first_name"`
	LastName    *string `json:"last_name"`
	Email       string  `json:"email"`
	DiscordId   *string `json:"discordId"`
	IsAdmin     bool    `json:"is_admin"`
	IsVolunteer bool    `json:"is_volunteer"`
	IsMentor    bool    `json:"is_mentor"`
	IsSponsor   bool    `json:"is_sponsor"`
}

type Hacker struct {
	Id                int     `json:"id"`
	UserId            string  `json:"userId"`
	FirstName         string  `json:"first_name"`
	LastName          string  `json:"last_name"`
	Age               int     `json:"age"`
	School            string  `json:"school"`
	Major             string  `json:"major"`
	GradYear          int     `json:"grad_year"`
	LevelOfStudy      string  `json:"level_of_study"`
	Country           string  `json:"country"`
	Email             string  `json:"email"`
	PhoneNumber       string  `json:"phone_number"`
	ResumePath        string  `json:"resume_path"`
	Github            *string `json:"github"`
	Linkedin          *string `json:"linkedin"`
	IsInternational   bool    `json:"is_international"`
	Gender            string  `json:"gender"`
	Pronouns          string  `json:"pronouns"`
	Ethnicity         string  `json:"ethnicity"`
	Avatar            int     `json:"avatar"`
	AgreedMLHNews     bool    `json:"agreed_mlh_news"`
	ApplicationStatus string  `json:"application_status"`
	CreatedAt         int64   `json:"created_at"`
	UpdatedAt         int64   `json:"updated_at"`
}
