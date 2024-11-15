package api

type User struct {
  ID string
  UserID string
  Name string
  Email string
  DiscordUsername *string
  Admin *bool
  HackerID *int
}
